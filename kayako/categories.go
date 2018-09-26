package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/inconshreveable/log15"
)

type category struct {
	logger       log15.Logger
	ID           int
	LegacyID     string
	URL          string
	path         string
	title        string
	description  string
	loaded       bool
	displayOrder int
	Sections     map[string]*section
	latestOrder  int
	keywords     []string
}

var updatedCategories map[int]bool

func init() {
	updatedCategories = make(map[int]bool)
}

func (c *category) load() error {
	if c.loaded {
		return nil
	}

	log := c.logger

	c.loaded = true

	data, err := ioutil.ReadFile(filepath.Join(c.path, "KEYWORDS.md"))
	if err == nil {
		lines := strings.Split(string(data), "\n")
		for i, line := range lines {
			line = strings.TrimSpace(line)
			if strings.ContainsAny(line, ",") {
				return fmt.Errorf("article '%s' contains illegal character in a defined keyword: %s", c.path, line)
			}
			lines[i] = line
		}
		c.keywords = lines
	} else {
		if !os.IsNotExist(err) {
			return err
		}
		c.keywords = make([]string, 0, 0)
	}

	log.Info(fmt.Sprintf("Keywords: %v", c.keywords))

	data, err = ioutil.ReadFile(filepath.Join(c.path, "README.md"))
	if err != nil {
		log.Warn("Category has no README.md file")
		return nil
	}

	s := string(data)
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		log.Warn("Category has no description")
		return nil
	}

	if s[0] != '#' {
		log.Debug("Category README.md file has no title")
		c.description = s
		return nil
	}

	lines := strings.SplitN(s, "\n", 2)
	if len(lines) != 2 {
		log.Warn("Category has no description")
		return nil
	}

	c.title = strings.TrimSpace(lines[0][1:])
	c.description = strings.TrimSpace(lines[1])

	return nil
}

func (c *category) Title() (string, error) {
	err := c.load()
	if err != nil {
		return "", err
	}
	if c.title == "" {
		return c.LegacyID, nil
	}
	return c.title, nil
}

func (c *category) Description() (string, error) {
	err := c.load()
	if err != nil {
		return "", err
	}
	return c.description, nil
}

func (c *category) resolve() error {

	log := c.logger
	log.Debug("Resolving ID & URL")

	url := fmt.Sprintf("https://vorteil.kayako.com/api/v1/categories.json?legacy_ids=%s", url.QueryEscape(c.LegacyID))
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
		info := new(schemaCategory)
		err = convert(response.Data, info)
		if err != nil {
			panic(err)
		}

		c.ID = info.ID
		c.URL = info.ResourceURL
		if c.URL == "" {
			panic(errors.New("resource url is an empty string"))
		}
		log.Debug(fmt.Sprintf("Resolved ID: %v", c.ID))
		log.Debug(fmt.Sprintf("Resolved URL: %s", c.URL))
	case http.StatusNotFound:
		err = c.draft()
		if err != nil {
			return err
		}
	default:
		log.Debug(response.body)
		return fmt.Errorf("unexpected response from server: status code %v", response.Status)
	}

	err = resolveSections(c)
	if err != nil {
		return err
	}

	return nil

}

func (c *category) draft() error {
	log := c.logger
	log.Info("Creating category skeleton")

	if c.displayOrder == 0 {
		model.latestOrder++
		c.displayOrder = model.latestOrder
	}
	c.displayOrder--

	title, err := c.Title()
	if err != nil {
		return err
	}

	description, err := c.Description()
	if err != nil {
		return err
	}

	m := map[string]interface{}{
		"titles": []localisation{localisation{
			Locale:      "en-us",
			Translation: title,
		}},
		"legacy_id": c.LegacyID,
		"descriptions": []localisation{localisation{
			Locale:      "en-us",
			Translation: description,
		}},
		"display_order": c.displayOrder,
	}

	data, err := json.Marshal(&m)
	if err != nil {
		panic(err)
	}

	params := bytes.NewReader(data)

	url := "https://vorteil.kayako.com/api/v1/categories.json"
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
		info := new(schemaCategory)
		err = convert(response.Data, info)
		if err != nil {
			panic(err)
		}

		c.ID = info.ID
		c.URL = info.ResourceURL
		if c.URL == "" {
			panic(errors.New("resource url is an empty string"))
		}
		log.Debug(fmt.Sprintf("Resolved ID: %v", c.ID))
		log.Debug(fmt.Sprintf("Resolved URL: %s", c.URL))
	default:
		log.Debug(response.body)
		return fmt.Errorf("unexpected response from server: status code %v", response.Status)
	}

	return nil
}

func (c *category) push() error {

	log := c.logger
	log.Debug("Pushing category")

	updatedCategories[c.ID] = true

	title, err := c.Title()
	if err != nil {
		return err
	}

	description, err := c.Description()
	if err != nil {
		return err
	}

	m := map[string]interface{}{
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

	url := fmt.Sprintf("https://vorteil.kayako.com/api/v1/categories/%v.json", c.ID)
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

	err = pushSections(c)
	if err != nil {
		return err
	}

	return nil

}

func scanCategories(path string) error {

	model.Categories = make(map[string]*category)

	fis, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, fi := range fis {
		if !fi.IsDir() {
			if fi.Name() != "README.md" {
				log.Warn(fmt.Sprintf("Skipping non-category object: %s", fi.Name()))
			}
			continue
		}

		err := scanCategory(filepath.Join(path, fi.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

func scanCategory(path string) error {

	c := new(category)
	c.LegacyID = filepath.Base(path)
	c.logger = log.New("category", c.LegacyID)
	c.path = path
	c.Sections = make(map[string]*section)

	log := c.logger
	log.Debug("Scanning category")

	key := filepath.Base(path)

	// check if filename includes article display order
	elems := strings.SplitN(key, "_", 2)
	if len(elems) == 2 {
		x, err := strconv.ParseUint(elems[0], 10, 64)
		if err == nil {
			c.displayOrder = int(x) + 1
			key = elems[1]
			if c.displayOrder > model.latestOrder {
				model.latestOrder = c.displayOrder
			}
		}
	}

	model.Categories[key] = c

	err := scanSections(c)
	if err != nil {
		return nil
	}

	return nil
}

func resolveCategories() error {

	for _, v := range model.Categories {
		err := v.resolve()
		if err != nil {
			return err
		}
	}

	return nil
}

func pushCategories() error {

	for _, v := range model.Categories {
		err := v.push()
		if err != nil {
			return err
		}
	}

	return nil

}
