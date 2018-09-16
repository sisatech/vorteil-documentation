package main

import (
	"fmt"
	"net/http"
)

func clean() error {

	for id, updated := range updatedArticles {
		if !updated {
			log.Warn(fmt.Sprintf("Stale article: %v", id))
			url := fmt.Sprintf("https://vorteil.kayako.com/api/v1/articles/%v.json", id)
			req, err := http.NewRequest(http.MethodDelete, url, nil)
			if err != nil {
				panic(err)
			}
			resp, err := do(req)
			if err != nil {
				return err
			}
			switch resp.Status {
			case http.StatusOK:
			default:
				panic(resp)
			}
		}
	}

	for id, updated := range updatedSections {
		if !updated {
			log.Warn(fmt.Sprintf("Stale section: %v", id))
			url := fmt.Sprintf("https://vorteil.kayako.com/api/v1/sections/%v.json", id)
			req, err := http.NewRequest(http.MethodDelete, url, nil)
			if err != nil {
				panic(err)
			}
			resp, err := do(req)
			if err != nil {
				return err
			}
			switch resp.Status {
			case http.StatusOK:
			default:
				panic(resp)
			}
		}
	}

	for id, updated := range updatedCategories {
		if !updated {
			log.Warn(fmt.Sprintf("Stale category: %v", id))
			url := fmt.Sprintf("https://vorteil.kayako.com/api/v1/categories/%v.json", id)
			req, err := http.NewRequest(http.MethodDelete, url, nil)
			if err != nil {
				panic(err)
			}
			resp, err := do(req)
			if err != nil {
				return err
			}
			switch resp.Status {
			case http.StatusOK:
			default:
				panic(resp)
			}
		}
	}

	return nil
}
