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

	"github.com/inconshreveable/log15"
	"golang.org/x/net/html"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type article struct {
	logger   log15.Logger
	section  *section
	URL      string
	ID       int
	LegacyID string
	path     string
	title    string
	content  string
	keywords []string
	loaded   bool
}

var updatedArticles map[int]bool

func init() {
	updatedArticles = make(map[int]bool)
}

func (a *article) load() error {
	if a.loaded {
		return nil
	}

	log := a.logger

	a.loaded = true

	data, err := ioutil.ReadFile(filepath.Join(a.path, "KEYWORDS.md"))
	if err == nil {
		lines := strings.Split(string(data), "\n")
		for i, line := range lines {
			line = strings.TrimSpace(line)
			if strings.ContainsAny(line, ",") {
				return fmt.Errorf("article '%s' contains illegal character in a defined keyword: %s", a.path, line)
			}
			lines[i] = line
		}
		a.keywords = lines
	} else {
		a.keywords = make([]string, 0, 0)
	}

	data, err = ioutil.ReadFile(filepath.Join(a.path, "README.md"))
	if err != nil {
		log.Warn("Article has no README.md file")
		data = []byte{}
		// return nil
	}

	str := string(data)
	str = strings.TrimSpace(str)

	if len(str) == 0 {
		log.Warn("Article has no description")
		str = ""
		// return nil
	}

	if str[0] != '#' {
		log.Debug("Article README.md file has no title")
		a.content = str
	} else {
		lines := strings.SplitN(str, "\n", 2)
		if len(lines) != 2 {
			log.Warn("Article has no description")
			lines = append(lines, "")
			// return nil
		}
		a.title = strings.TrimSpace(lines[0][1:])
		a.content = strings.TrimSpace(lines[1])
	}

	return a.convertMarkdownToHTML()

}

func (a *article) Title() (string, error) {
	err := a.load()
	if err != nil {
		return "", err
	}
	if a.title == "" {
		return a.LegacyID, nil
	}
	return a.title, nil
}

func (a *article) Content() (string, error) {
	err := a.load()
	if err != nil {
		return "", err
	}
	return a.content, nil
}

func (a *article) resolve() error {

	log := a.logger
	log.Debug("Resolving ID & URL")

	url := fmt.Sprintf("https://vorteil.kayako.com/api/v1/articles.json?legacy_ids=%s", url.QueryEscape(a.LegacyID))
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
		info := new(schemaArticle)
		err = convert(response.Data, info)
		if err != nil {
			panic(err)
		}

		a.ID = info.ID
		a.URL = info.HelpCenterURL
		if a.URL == "" {
			panic(errors.New("resource url is an empty string"))
		}
		log.Debug(fmt.Sprintf("Resolved ID: %v", a.ID))
		log.Debug(fmt.Sprintf("Resolved URL: %s", a.URL))
	case http.StatusNotFound:
		err = a.draft()
		if err != nil {
			return err
		}
	default:
		log.Debug(response.body)
		return fmt.Errorf("unexpected response from server: status code %v", response.Status)
	}

	return nil

}

func (a *article) draft() error {
	log := a.logger
	log.Info("Creating article skeleton")

	title, err := a.Title()
	if err != nil {
		return err
	}

	content, err := a.Content()
	if err != nil {
		return err
	}

	m := map[string]interface{}{
		"section_id": a.section.ID,
		"titles": []localisation{localisation{
			Locale:      "en-us",
			Translation: title,
		}},
		"legacy_id": a.LegacyID,
		"contents": []localisation{localisation{
			Locale:      "en-us",
			Translation: content,
		}},
		"keywords": strings.Join(append(append(a.keywords, a.section.keywords...), a.section.category.keywords...), ","),
		"status":   "DRAFT",
	}

	data, err := json.Marshal(&m)
	if err != nil {
		panic(err)
	}

	params := bytes.NewReader(data)

	url := "https://vorteil.kayako.com/api/v1/articles.json"
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
		info := new(schemaArticle)
		err = convert(response.Data, info)
		if err != nil {
			panic(err)
		}

		a.ID = info.ID
		a.URL = info.HelpCenterURL
		if a.URL == "" {
			panic(errors.New("resource url is an empty string"))
		}
		log.Debug(fmt.Sprintf("Resolved ID: %v", a.ID))
		log.Debug(fmt.Sprintf("Resolved URL: %s", a.URL))
	case http.StatusBadRequest:
		log.Debug(string(data))
		log.Debug(response.body)
		return fmt.Errorf("invalid data in: %s", a.path)
	case http.StatusNotFound:
		err = a.draft()
		if err != nil {
			return err
		}
	default:
		log.Debug(response.body)
		return fmt.Errorf("unexpected response from server: status code %v", response.Status)
	}

	return nil
}

func (a *article) push() error {
	log := a.logger
	log.Info("Pushing article")

	updatedArticles[a.ID] = true

	title, err := a.Title()
	if err != nil {
		return err
	}

	content, err := a.Content()
	if err != nil {
		return err
	}

	m := map[string]interface{}{
		"section_id": a.section.ID,
		"titles": []localisation{localisation{
			Locale:      "en-us",
			Translation: title,
		}},
		"contents": []localisation{localisation{
			Locale:      "en-us",
			Translation: content,
		}},
		"keywords": strings.Join(append(append(a.keywords, a.section.keywords...), a.section.category.keywords...), ","),
		"status":   "PUBLISHED",
		// TODO: files
		// TODO: tags
	}

	data, err := json.Marshal(&m)
	if err != nil {
		panic(err)
	}

	params := bytes.NewReader(data)

	url := fmt.Sprintf("https://vorteil.kayako.com/api/v1/articles/%v.json", a.ID)
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

	return nil
}

func scanArticles(s *section) error {

	log := s.logger

	fis, err := ioutil.ReadDir(s.path)
	if err != nil {
		return err
	}

	for _, fi := range fis {
		if !fi.IsDir() {
			if fi.Name() != "README.md" {
				log.Warn(fmt.Sprintf("Skipping non-article object: %s", fi.Name()))
			}
			continue
		}

		err := scanArticle(s, filepath.Join(s.path, fi.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

func scanArticle(s *section, path string) error {

	a := new(article)
	a.section = s
	a.LegacyID = filepath.Join(s.LegacyID, filepath.Base(path))
	a.logger = s.logger.New("article", a.LegacyID)
	a.path = path

	log.Debug("Scanning article")

	key := filepath.Base(path)
	s.Articles[key] = a

	return nil
}

func resolveArticles(s *section) error {

	for _, v := range s.Articles {
		err := v.resolve()
		if err != nil {
			return err
		}
	}

	return nil
}

func pushArticles(s *section) error {

	for _, v := range s.Articles {
		err := v.push()
		if err != nil {
			return err
		}
	}

	return nil

}

func (a *article) convertMarkdownToHTML() error {
	log := a.logger
	log.Debug("Converting Markdown to HTML in article")

	if strings.ContainsAny(a.content, "<>&") {
		return fmt.Errorf("article contains an illegal character (TODO: solve this annoying problem)")
	}

	a.content = string(blackfriday.Run([]byte(a.content)))
	if strings.TrimSpace(a.content) == "" {
		a.content = "PLACEHOLDER"
	}
	return nil
}

func precheckLinks() error {
	for _, category := range model.Categories {
		category.logger.Debug("Pre-checking links in category")
		for _, section := range category.Sections {
			section.logger.Debug("Pre-checking links in section")
			for _, article := range section.Articles {
				err := article.precheckLinks()
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (a *article) resolveLink(old string) (string, error) {

	if strings.HasPrefix(old, "http://") || strings.HasPrefix(old, "https://") {
		return old, nil
	}
	path := filepath.Clean(old)
	elems := strings.Split(path, "/")
	var relative int
	category := a.section.category
	section := a.section
	article := a
	for _, elem := range elems {
		switch elem {
		case "..":
			relative--
			continue
		case ".":
			panic("didn't we clean this?")
		default:
			relative++
		}

		var ok bool
		switch relative {
		case 0:
			article, ok = section.Articles[elem]
			if !ok {
				return "", fmt.Errorf("broken relative link due to missing article: %s", old)
			}
		case -1:
			section, ok = category.Sections[elem]
			if !ok {
				return "", fmt.Errorf("broken relative link due to missing section: %s", old)
			}
		case -2:
			category, ok = model.Categories[elem]
			if !ok {
				return "", fmt.Errorf("broken relative link due to missing category: %s", old)
			}
		default:
			return "", fmt.Errorf("bad relative link: %s", old)
		}
	}

	switch relative {
	case 0:
		return article.URL, nil
	case -1:
		return section.URL, nil
	case -2:
		return category.URL, nil
	default:
		panic(fmt.Errorf("broken relative link doesn't point to another object: %s", old))
	}

}

func (a *article) precheckLinks() error {
	log := a.logger
	log.Debug("Pre-checking links in article")

	var fn func(*html.Node) error
	fn = func(n *html.Node) error {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					oldLink := attr.Val
					_, err := a.resolveLink(oldLink)
					if err != nil {
						return err
					}
					// n.Attr[i].Val = newLink
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			err := fn(c)
			if err != nil {
				return err
			}
		}
		return nil
	}

	content, err := a.Content()
	if err != nil {
		return fmt.Errorf("article %s: %v", a.LegacyID, err)
	}

	doc, err := html.Parse(strings.NewReader(content))
	if err != nil {
		return fmt.Errorf("article %s: %v", a.LegacyID, err)
	}

	err = fn(doc)
	if err != nil {
		return fmt.Errorf("article %s: %v", a.LegacyID, err)
	}
	return nil
}

func resolveLinks() error {
	for _, category := range model.Categories {
		category.logger.Debug("Resolving links in category")
		for _, section := range category.Sections {
			section.logger.Debug("Resolving links in section")
			for _, article := range section.Articles {
				err := article.resolveLinks()
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (a *article) resolveLinks() error {
	log := a.logger
	log.Debug("Resolving links in article")

	var fn func(*html.Node) error
	fn = func(n *html.Node) error {
		if n.Type == html.ElementNode && n.Data == "a" {
			for i, attr := range n.Attr {
				if attr.Key == "href" {
					oldLink := attr.Val
					newLink, err := a.resolveLink(oldLink)
					if err != nil {
						return err
					}
					n.Attr[i].Val = newLink
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			err := fn(c)
			if err != nil {
				return err
			}
		}
		return nil
	}

	content, err := a.Content()
	if err != nil {
		return fmt.Errorf("article %s: %v", a.LegacyID, err)
	}

	doc, err := html.Parse(strings.NewReader(content))
	if err != nil {
		return fmt.Errorf("article %s: %v", a.LegacyID, err)
	}

	err = fn(doc)
	if err != nil {
		return fmt.Errorf("article %s: %v", a.LegacyID, err)
	}

	buf := new(bytes.Buffer)
	err = html.Render(buf, doc)
	if err != nil {
		return fmt.Errorf("article %s: %v", a.LegacyID, err)
	}

	a.content = buf.String()
	return nil
}
