package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/lunny/html2md"
)

var backupDir string

func scrape(path string) error {

	backupDir = path

	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}

	fis, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	if len(fis) > 0 {
		return fmt.Errorf("backup directory not empty: %s", path)
	}

	err = scrapeCategories()
	if err != nil {
		return err
	}

	return nil
}

func scrapeCategories() error {

	url := "https://vorteil.kayako.com/api/v1/categories.json"

	// looping because result may not contain all categories
	for {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return err
		}

		resp, err := do(req)
		if err != nil {
			return err
		}

		switch resp.Status {
		case http.StatusOK:
			slice, ok := resp.Data.([]interface{})
			if !ok {
				log.Debug(resp.body)
				panic("unexpected response structure")
			}

			for _, elem := range slice {
				category := new(schemaCategory)
				err = convert(elem, category)
				if err != nil {
					log.Debug(resp.body)
					panic(err)
				}

				legacyID := category.LegacyID
				var title, description string

				title, err = category.resolveTitle()
				if err != nil {
					return err
				}
				description, err = category.resolveDescription()
				if err != nil {
					return err
				}

				log.Debug(fmt.Sprintf("Backing up category: %s", title))

				if legacyID == "" {
					legacyID = strings.Replace(title, " ", "_", -1)
				}
				legacyID = strings.Replace(title, "/", "_", -1)

				dir := filepath.Join(backupDir, legacyID)
				err = os.MkdirAll(dir, 0777)
				if err != nil {
					panic(err)
				}
				filename := filepath.Join(dir, "README.md")
				readme := fmt.Sprintf("# %s\n\n%s", title, description)
				err = ioutil.WriteFile(filename, []byte(readme), 0777)
				if err != nil {
					panic(err)
				}

				updatedCategories[category.ID] = false

				err = scrapeSections(dir, category.ID)
				if err != nil {
					return err
				}

			}

		default:
			panic(fmt.Sprintf("unexpected response status: %v", resp.Status))
		}

		url = resp.NextURL
		if url == "" {
			break
		}

	}

	return nil
}

func scrapeSections(base string, id int) error {

	url := fmt.Sprintf("https://vorteil.kayako.com/api/v1/sections.json?category_ids=%v", id)

	// looping because result may not contain all categories
	for {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return err
		}

		resp, err := do(req)
		if err != nil {
			return err
		}

		switch resp.Status {
		case http.StatusOK:
			slice, ok := resp.Data.([]interface{})
			if !ok {
				log.Debug(resp.body)
				panic("unexpected response structure")
			}

			for _, elem := range slice {
				section := new(schemaSection)
				err = convert(elem, section)
				if err != nil {
					log.Debug(resp.body)
					panic(err)
				}

				legacyID := section.LegacyID
				var title, description string

				title, err = section.resolveTitle()
				if err != nil {
					return err
				}
				description, err = section.resolveDescription()
				if err != nil {
					return err
				}

				log.Debug(fmt.Sprintf("Backing up section: %s", title))

				if legacyID == "" {
					legacyID = strings.Replace(title, " ", "_", -1)
				}
				legacyID = strings.Replace(title, "/", "_", -1)

				dir := filepath.Join(base, legacyID)
				err = os.MkdirAll(dir, 0777)
				if err != nil {
					panic(err)
				}
				filename := filepath.Join(dir, "README.md")
				readme := fmt.Sprintf("# %s\n\n%s", title, description)
				err = ioutil.WriteFile(filename, []byte(readme), 0777)
				if err != nil {
					panic(err)
				}

				updatedSections[section.ID] = false

				err = scrapeArticles(dir, section.ID)
				if err != nil {
					return err
				}

			}

		default:
			panic(fmt.Sprintf("unexpected response status: %v", resp.Status))
		}

		url = resp.NextURL
		if url == "" {
			break
		}

	}

	return nil
}

func scrapeArticles(base string, id int) error {

	url := fmt.Sprintf("https://vorteil.kayako.com/api/v1/articles.json?section_id=%v", id)

	// looping because result may not contain all categories
	for {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return err
		}

		resp, err := do(req)
		if err != nil {
			return err
		}

		switch resp.Status {
		case http.StatusOK:
			slice, ok := resp.Data.([]interface{})
			if !ok {
				log.Debug(resp.body)
				panic("unexpected response structure")
			}

			for _, elem := range slice {
				article := new(schemaArticle)
				err = convert(elem, article)
				if err != nil {
					log.Debug(resp.body)
					panic(err)
				}

				legacyID := article.LegacyID
				var title, content string

				title, err = article.resolveTitle()
				if err != nil {
					return err
				}
				content, err = article.resolveContent()
				if err != nil {
					return err
				}

				log.Debug(fmt.Sprintf("Backing up article: %s", title))

				content = html2md.Convert(content)

				if legacyID == "" {
					legacyID = strings.Replace(title, " ", "_", -1)
				}
				legacyID = strings.Replace(title, "/", "_", -1)

				dir := filepath.Join(base, legacyID)
				err = os.MkdirAll(dir, 0777)
				if err != nil {
					panic(err)
				}
				filename := filepath.Join(dir, "README.md")
				readme := fmt.Sprintf("# %s\n\n%s", title, content)
				err = ioutil.WriteFile(filename, []byte(readme), 0777)
				if err != nil {
					panic(err)
				}

				updatedArticles[article.ID] = false

			}

		default:
			panic(fmt.Sprintf("unexpected response status: %v", resp.Status))
		}

		url = resp.NextURL
		if url == "" {
			break
		}

	}

	return nil
}
