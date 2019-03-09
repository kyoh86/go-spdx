.SUFFIXES: .y .go
.y.go:
	goyacc -o $@ $<
	gofmt -w $@
	rm y.output

npm-install:
	npm install

licenses-asset: npm-install
	node fetch_list.js
	goblet -g -p spdx -o licenses_asset.go  --ignore-dotfiles ./json
	gofmt -w licenses_asset.go

gen: parser.go licenses-asset

clean:
	rm parser.go
