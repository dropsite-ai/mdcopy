.PHONY: build install test

build:
	goreleaser release --snapshot --clean

install: build
	sudo cp -f ./dist/mdcopy_darwin_arm64_v8.0/mdcopy /usr/local/bin/mdcopy

test:
	go test ./... -v -cover
