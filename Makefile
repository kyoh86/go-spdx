.PHONY: gen lint test sample

VERSION := `git vertag get`
COMMIT  := `git rev-parse HEAD`

gen:
	$(MAKE) gen -C spdx

lint: gen
	golangci-lint run

test: lint
	go test v --race ./...

sample:
	go run -tags=sample ./cmd/go-spdx-sample/main.go
