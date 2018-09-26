package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/inconshreveable/log15"
)

type section struct {
	logger      log15.Logger
	category    *category
	ID          int
	LegacyID    string
	URL         string
	path        string
	title       string
	description string
	loaded      bool
	Articles    map[string]*article
	keywords    []string
}

var updatedSections map[int]bool

func init() {
	updatedSections = make(map[int]bool)
}

func (s *section) load() error {
	if s.loaded {
		return nil
	}

	log := s.logger

	s.loaded = true

	data, err := ioutil.ReadFile(filepath.Join(s.path, "KEYWORDS.md"))
	if err == nil {
		lines := strings.Split(string(data), "\n")
		for i, line := range lines {
			line = strings.TrimSpace(line)
			if strings.ContainsAny(line, ",") {
				return fmt.Errorf("section '%s' contains illegal character in a defined keyword: %s", s.path, line)
			}
			lines[i] = line
		}
		s.keywords = lines
	} else {
		s.keywords = make([]string, 0, 0)
	}

	log.Info(fmt.Sprintf("Keywords: %v", append(s.category.keywords, s.keywords...)))

	data, err = ioutil.ReadFile(filepath.Join(s.path, "README.md"))
	if err != nil {
		log.Warn("Category has no README.md file")
		return nil
	}

	str := string(data)
	str = strings.TrimSpace(str)

	if len(str) == 0 {
		log.Warn("Category has no description")
		return nil
	}

	if str[0] != '#' {
		log.Debug("Category README.md file has no title")
		s.description = str
		return nil
	}

	lines := strings.SplitN(str, "\n", 2)
	if len(lines) != 2 {
		log.Warn("Category has no description")
		return nil
	}

	s.title = strings.TrimSpace(lines[0][1:])
	s.description = strings.TrimSpace(lines[1])

	return nil
}

func (s *section) Title() (string, error) {
	err := s.load()
	if err != nil {
		return "", err
	}
	if s.title == "" {
		return s.LegacyID, nil
	}
	return s.title, nil
}

func (s *section) Description() (string, error) {
	err := s.load()
	if err != nil {
		return "", err
	}
	return s.description, nil
}

func (s *section) resolve() error {

	log := s.logger
	log.Debug("Resolving ID & URL")

	url := fmt.Sprintf("https://vorteil.kayako.com/api/v1/sections.json?legacy_ids=%s", url.QueryEscape(s.LegacyID))
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	response, err := do(req)
	if err != nil {
		return err
	}

	switch response.Status {
	case http.StatusOK:
		info := new(schemaSection)
		err = convert(response.Data, info)
		if err != nil {
			panic(err)
		}

		s.ID = info.ID
		s.URL = info.ResourceURL
		if s.URL == "" {
			panic(errors.New("resource url is an empty string"))
		}
		log.Debug(fmt.Sprintf("Resolved ID: %v", s.ID))
		log.Debug(fmt.Sprintf("Resolved URL: %s", s.URL))
	case http.StatusNotFound:
		err = s.draft()
		if err != nil {
			return err
		}
	default:
		log.Debug(response.body)
		return fmt.Errorf("unexpected response from server: status code %v", response.Status)
	}

	err = resolveArticles(s)
	if err != nil {
		return err
	}

	return nil

}

func (s *section) draft() error {
	log := s.logger
	log.Info("Creating section skeleton")

	title, err := s.Title()
	if err != nil {
		return err
	}

	description, err := s.Description()
	if err != nil {
		return err
	}

	m := map[string]interface{}{
		"category_id": s.category.ID,
		"titles": []localisation{localisation{
			Locale:      "en-us",
			Translation: title,
		}},
		"legacy_id": s.LegacyID,
		"descriptions": []localisation{localisation{
			Locale:      "en-us",
			Translation: description,
		}},
		"article_order_by": "POPULAR",
	}

	data, err := json.Marshal(&m)
	if err != nil {
		panic(err)
	}

	// put logic in a for loop to automatically retry in the case of internal server errors
	var failures int
	delay := time.Second

	for {
		params := bytes.NewReader(data)

		url := "https://vorteil.kayako.com/api/v1/sections.json"
		req, err := http.NewRequest(http.MethodPost, url, params)
		if err != nil {
			return err
		}

		response, err := do(req)
		if err != nil {
			return err
		}

		switch response.Status {
		case http.StatusCreated:
			info := new(schemaSection)
			err = convert(response.Data, info)
			if err != nil {
				panic(err)
			}

			s.ID = info.ID
			s.URL = info.ResourceURL
			if s.URL == "" {
				panic(errors.New("resource url is an empty string"))
			}
			log.Debug(fmt.Sprintf("Resolved ID: %v", s.ID))
			log.Debug(fmt.Sprintf("Resolved URL: %s", s.URL))
			// log.Debug(fmt.Sprintf("BODY: %s", string(data)))
			// log.Debug(fmt.Sprintf("RESPONSE: %s", response.body))
		case http.StatusInternalServerError:
			failures++
			if failures > 10 {
				// log.Debug(fmt.Sprintf("BODY: %s", string(data)))
				// log.Debug(fmt.Sprintf("RESPONSE: %s", response.body))
				return errors.New("failed too many times in a row")
			}
			log.Debug(fmt.Sprintf("Retrying in %v", delay))
			<-time.After(delay)
			delay += time.Second
			continue
		default:
			log.Debug(response.body)
			return fmt.Errorf("unexpected response from server: status code %v", response.Status)
		}

		return nil
	}
}

func (s *section) push() error {
	log := s.logger
	log.Info("Pushing section")

	updatedSections[s.ID] = true

	title, err := s.Title()
	if err != nil {
		return err
	}

	description, err := s.Description()
	if err != nil {
		return err
	}

	m := map[string]interface{}{
		"category_id": s.category.ID,
		"titles": []localisation{localisation{
			Locale:      "en-us",
			Translation: title,
		}},
		"descriptions": []localisation{localisation{
			Locale:      "en-us",
			Translation: description,
		}},
	}

	data, err := json.Marshal(&m)
	if err != nil {
		panic(err)
	}

	params := bytes.NewReader(data)

	url := fmt.Sprintf("https://vorteil.kayako.com/api/v1/sections/%v.json", s.ID)
	req, err := http.NewRequest(http.MethodPut, url, params)
	if err != nil {
		return err
	}

	response, err := do(req)
	if err != nil {
		return err
	}

	switch response.Status {
	case http.StatusOK:
	default:
		log.Debug(response.body)
		return fmt.Errorf("unexpected response from server: status code %v", response.Status)
	}

	err = pushArticles(s)
	if err != nil {
		return err
	}

	return nil
}

func scanSections(c *category) error {

	log := c.logger

	fis, err := ioutil.ReadDir(c.path)
	if err != nil {
		return err
	}

	for _, fi := range fis {
		if !fi.IsDir() {
			if fi.Name() != "README.md" {
				log.Warn(fmt.Sprintf("Skipping non-section object: %s", fi.Name()))
			}
			continue
		}

		err := scanSection(c, filepath.Join(c.path, fi.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

func scanSection(c *category, path string) error {

	s := new(section)
	s.category = c
	s.LegacyID = filepath.Join(c.LegacyID, filepath.Base(path))
	s.logger = c.logger.New("section", s.LegacyID)
	s.path = path
	s.Articles = make(map[string]*article)

	log.Debug("Scanning section")

	key := filepath.Base(path)
	c.Sections[key] = s

	err := scanArticles(s)
	if err != nil {
		return nil
	}

	return nil
}

func resolveSections(c *category) error {

	for _, v := range c.Sections {
		err := v.resolve()
		if err != nil {
			return err
		}
	}

	return nil
}

func pushSections(c *category) error {

	for _, v := range c.Sections {
		err := v.push()
		if err != nil {
			return err
		}
	}

	return nil

}
