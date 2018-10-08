package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type documentation struct {
	Categories map[string]*category
}

var model documentation

func scan(path, backup string, firstTime bool) error {

	log.Info(fmt.Sprintf("Scanning: %s", path))

	err := scanCategories(path)
	if err != nil {
		return err
	}

	err = precheckLinks()
	if err != nil {
		return err
	}

	log.Info("Validation succeeded.")
	if !sync {
		log.Info("Skipping sync.")
		return nil
	}

	if backup == "" {
		log.Info(fmt.Sprintf("Skipping backup."))
	} else {
		// scrape existing documentation as a backup
		log.Info(fmt.Sprintf("Performing backup to: %s", backup))
		err = scrape(backup)
		if err != nil {
			return err
		}
	}

	// load manifest
	mf, err := loadManifest(path, firstTime)
	if err != nil {
		return err
	}

	// resolve all IDs & URLs
	err = resolveCategories(mf)
	if err != nil {
		return err
	}

	err = resolveLinks()
	if err != nil {
		return err
	}

	// push
	err = pushCategories()
	if err != nil {
		return err
	}

	// clean up outdated documentation
	err = clean()
	if err != nil {
		return err
	}

	// generate URL manifest
	err = manifest(path)
	if err != nil {
		return err
	}

	return nil
}

// ManifestTuple ..
type ManifestTuple struct {
	ID  int
	URL string
}

// Manifest ..
type Manifest struct {
	Categories map[string]ManifestTuple
	Sections   map[string]ManifestTuple
	Articles   map[string]ManifestTuple
}

func manifest(path string) error {
	manifest := new(Manifest)
	manifest.Categories = make(map[string]ManifestTuple)
	manifest.Sections = make(map[string]ManifestTuple)
	manifest.Articles = make(map[string]ManifestTuple)
	for _, category := range model.Categories {
		manifest.Categories[category.LegacyID] = ManifestTuple{ID: category.ID, URL: category.URL}
		for _, section := range category.Sections {
			manifest.Sections[section.LegacyID] = ManifestTuple{ID: section.ID, URL: section.URL}
			for _, article := range section.Articles {
				manifest.Articles[article.LegacyID] = ManifestTuple{ID: article.ID, URL: article.URL}
			}
		}
	}
	data, err := json.MarshalIndent(manifest, "", "    ")
	if err != nil {
		panic(err)
	}
	return ioutil.WriteFile(filepath.Join(path, "manifest.json"), data, 0666)
}

func loadManifest(path string, firstTime bool) (*Manifest, error) {
	manifest := new(Manifest)
	data, err := ioutil.ReadFile(filepath.Join(path, "manifest.json"))
	if err != nil {
		if firstTime && os.IsNotExist(err) {
			manifest.Categories = make(map[string]ManifestTuple)
			manifest.Sections = make(map[string]ManifestTuple)
			manifest.Articles = make(map[string]ManifestTuple)
			return manifest, nil
		}
		return nil, err
	}
	err = json.Unmarshal(data, manifest)
	if err != nil {
		return nil, err
	}
	return manifest, nil
}
