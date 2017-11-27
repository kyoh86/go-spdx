.PHONY: default gen clean test vendor install

default:
	echo use gen, test, vendor or install

gen:
	$(MAKE) gen -C spdx

clean:
	$(MAKE) clean -C spdx

test:
	go test ./...

setup:
	go get -u golang.org/x/tools/cmd/goyacc

vendor:
	dep ensure

example:
	go run cmd/go-spdx-example/main.go

install:
	go install ./...
