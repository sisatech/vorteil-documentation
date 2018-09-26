package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type response struct {
	Status     int         `json:"status"`
	Data       interface{} `json:"data"`
	Resource   string      `json:"resource"`
	MissingIDs []string    `json:"missing_ids"`
	TotalCount int         `json:"total_count"`
	SessionID  string      `json:"session_id"`
	Offset     int         `json:"offset"`
	Limit      int         `json:"limit"`
	NextURL    string      `json:"next_url"`
	LastURL    string      `json:"last_url"`
	body       string
}

type localisation struct {
	Locale      string `json:"locale"`
	Translation string `json:"translation"`
}

func resolveLocalisation(id int) (string, error) {
	url := fmt.Sprintf("https://vorteil.kayako.com/api/v1/locale/fields/%d.json", id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := do(req)
	if err != nil {
		return "", err
	}

	switch resp.Status {
	case http.StatusOK:
	default:
		panic(err)
	}

	locale := new(localisation)
	err = convert(resp.Data, locale)
	if err != nil {
		return "", err
	}

	return locale.Translation, nil
}

type resource struct {
	ID           int    `json:"id"`
	ResourceType string `json:"resource_type"`
}

type schemaCategory struct {
	resource
	Titles       []resource `json:"titles"`
	Slugs        []resource `json:"slugs"`
	Descriptions []resource `json:"descriptions"`
	Brand        resource   `json:"brand"`
	DisplayOrder int        `json:"display_order"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	LegacyID     string     `json:"legacy_id"`
	ResourceURL  string     `json:"resource_url"`
}

func (o *schemaCategory) resolveTitle() (string, error) {

	if len(o.Titles) != 1 {
		panic("too many titles")
	}

	title := o.Titles[0]
	if title.ResourceType != "locale_field" {
		panic("unexpected data")
	}

	return resolveLocalisation(title.ID)

}

func (o *schemaCategory) resolveDescription() (string, error) {

	if len(o.Descriptions) != 1 {
		panic("too many descriptions")
	}

	description := o.Descriptions[0]
	if description.ResourceType != "locale_field" {
		panic("unexpected data")
	}

	return resolveLocalisation(description.ID)

}

type schemaSection struct {
	resource
	Titles       []resource `json:"titles"`
	Slugs        []resource `json:"slugs"`
	Descriptions []resource `json:"descriptions"`
	Brand        resource   `json:"brand"`
	DisplayOrder int        `json:"display_order"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	LegacyID     string     `json:"legacy_id"`
	ResourceURL  string     `json:"resource_url"`
}

func (o *schemaSection) resolveTitle() (string, error) {

	if len(o.Titles) != 1 {
		panic("too many titles")
	}

	title := o.Titles[0]
	if title.ResourceType != "locale_field" {
		panic("unexpected data")
	}

	return resolveLocalisation(title.ID)

}

func (o *schemaSection) resolveDescription() (string, error) {

	if len(o.Descriptions) != 1 {
		panic("too many descriptions")
	}

	description := o.Descriptions[0]
	if description.ResourceType != "locale_field" {
		panic("unexpected data")
	}

	return resolveLocalisation(description.ID)

}

type schemaArticle struct {
	resource
	Titles      []resource `json:"titles"`
	Slugs       []resource `json:"slugs"`
	Contents    []resource `json:"contents"`
	Keywords    string     `json:"keywords"`
	Section     resource   `json:"section"`
	Creator     resource   `json:"creator"`
	Author      resource   `json:"author"`
	Attachments []resource `json:"attachments"`
	// DownloadAll
	Status        string     `json:"status"`
	UpvoteCount   int        `json:"upvote_count"`
	DownvoteCount int        `json:"downvote_count"`
	Views         int        `json:"views"`
	Rank          int        `json:"rank"`
	Tags          []resource `json:"tags"`
	IsFeatured    bool       `json:"is_featured"`
	AllowComments bool       `json:"allow_comments"`
	TotalComments int        `json:"total_comments"`
	CreatedAt     time.Time  `json:"created_at"`
	PublishedAt   time.Time  `json:"published_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	LegacyID      string     `json:"legacy_id"`
	UUID          string     `json:"uuid"`
	HelpCenterURL string     `json:"helpcenter_url"`
	ResourceURL   string     `json:"resource_url"`
}

func (o *schemaArticle) resolveTitle() (string, error) {

	if len(o.Titles) != 1 {
		panic("too many titles")
	}

	title := o.Titles[0]
	if title.ResourceType != "locale_field" {
		panic("unexpected data")
	}

	return resolveLocalisation(title.ID)

}

func (o *schemaArticle) resolveContent() (string, error) {

	if len(o.Contents) != 1 {
		panic("too many contents")
	}

	contents := o.Contents[0]
	if contents.ResourceType != "locale_field" {
		panic("unexpected data")
	}

	return resolveLocalisation(contents.ID)

}

func convert(a interface{}, b interface{}) error {

	if v, ok := a.([]interface{}); ok {
		if len(v) == 0 {
			return errors.New("no data in response payload")
		}
		a = v[0]
	}

	data, err := json.Marshal(a)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, b)
}
