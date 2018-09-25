package main

import (
	"fmt"
)

type documentation struct {
	Categories  map[string]*category
	latestOrder int
}

var model documentation

func scan(path, backup string) error {

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

	// resolve all IDs & URLs
	err = resolveCategories()
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

	return nil
}
