.PHONY: default gen clean test lint setup example

default:
	echo use gen, clean test, setup or example

gen:
	$(MAKE) gen -C spdx

clean:
	$(MAKE) clean -C spdx

test:
	go test ./...

lint:
	gometalinter ./...

setup:
	go get -u golang.org/x/tools/cmd/goyacc

example:
	go run -tags=example cmd/go-spdx-example/main.go
