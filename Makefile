.PHONY: default gen clean test setup example

default:
	echo use gen, clean test, setup or example

gen:
	$(MAKE) gen -C spdx

clean:
	$(MAKE) clean -C spdx

test:
	go test ./...

setup:
	go get -u golang.org/x/tools/cmd/goyacc

example:
	go run -tags=example cmd/go-spdx-example/main.go
