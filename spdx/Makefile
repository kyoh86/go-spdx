.SUFFIXES: .y .go
.y.go:
	goyacc -o $@ $<
	gofmt -w $@
	rm y.output

npm-install:
	npm install

licenses-asset: npm-install
	node fetch_list.js

gen: parser.go licenses-asset

clean:
	rm parser.go
