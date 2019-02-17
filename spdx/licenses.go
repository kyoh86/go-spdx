package spdx

import (
	"encoding/json"
	"errors"
)

// License has SPDX license profile.
type License struct {
	ID          string `json:"-"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	OSIApproved bool   `json:"osiApproved"`
	Text        string `json:"licenseText"`
}

var (
	validLicenses map[string]License
	// ErrorNotFound represents that an id for license is not valid
	ErrorNotFound = errors.New("not found")
)

func init() {
	licenses, err := licensesJsonBytes()
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(licenses, &validLicenses); err != nil {
		panic(err)
	}
	for id, l := range validLicenses {
		l.ID = id
		validLicenses[id] = l
	}
}

// Get a license with an ID.
func Get(id string) (*License, error) {
	l, ok := validLicenses[id]
	if !ok {
		return nil, ErrorNotFound
	}
	return &l, nil
}

// List will get valid licenses.
func List() (licenses []License) {
	for _, l := range validLicenses {
		licenses = append(licenses, l)
	}
	return
}
