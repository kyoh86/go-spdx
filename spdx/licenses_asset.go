package spdx

import (
	_ "embed"
)

//go:embed json/licenses.json
var licenses []byte
