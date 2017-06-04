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

var (
	validLicenses map[string]License
	ErrorNotFound = errors.New("not found")
)

// Get a license with an ID.
func Get(id string) (*License, error) {
	l, ok := validLicenses[id]
	if !ok {
		return nil, ErrorNotFound
	}
	return &l, nil
}
