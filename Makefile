default: gen test

test:
	go test ./...

gen:
	$(MAKE) gen -C spdx

clean:
	$(MAKE) clean -C spdx

setup:
	go get -u golang.org/x/tools/cmd/goyacc

example:
	go run cmd/go-spdx-example/main.go
