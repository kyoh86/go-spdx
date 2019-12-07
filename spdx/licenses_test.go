package spdx

import (
	"testing"
)

func TestSPDX(t *testing.T) {
	t.Run("invalid license id is not found", func(t *testing.T) {
		_, err := Get("invalid license name")
		if err == nil {
			t.Error("expected error with getting invalid license id, but not error")
		}
	})
	t.Run("mit", func(t *testing.T) {
		mit, err := Get("MIT")
		if err != nil {
			t.Error("expected to be able to get MIT license, but got an error")
		}
		if mit.ID != "MIT" {
			t.Errorf("expected to be able to get MIT license, but got another license: (%s)", mit.ID)
		}
	})

	t.Run("list", func(t *testing.T) {
		if len(List()) == 0 {
			t.Error("there's no license")
		}
	})
}
