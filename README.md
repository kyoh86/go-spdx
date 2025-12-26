# go-spdx

The package parses SPDX license expression strings describing license terms

[![PkgGoDev](https://pkg.go.dev/badge/kyoh86/go-spdx)](https://pkg.go.dev/kyoh86/go-spdx)
[![Go Report Card](https://goreportcard.com/badge/github.com/kyoh86/go-spdx)](https://goreportcard.com/report/github.com/kyoh86/go-spdx)
[![Release](https://github.com/kyoh86/go-spdx/workflows/Release/badge.svg)](https://github.com/kyoh86/go-spdx/releases)

## Install

```
go get github.com/kyoh86/go-spdx
```

## Usage

See [example](https://github.com/kyoh86/go-spdx/blob/main/cmd/go-spdx-example/main.go) or [test](https://github.com/kyoh86/go-spdx/blob/main/spdx/parser_test.go)

## Updating licenses

In order to update the licenses_asset.go file with the newest license list from SPDX, execute the following commands:

```
cd spdx
npm install spdx-license-list@latest
make licenses-asset
```

(You should also add a test to spdx/parser_test.go to make sure everything worked as expected.)

# LICENSE

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](http://www.opensource.org/licenses/MIT)

This is distributed under the [MIT License](http://www.opensource.org/licenses/MIT).
