# mdcopy

Copies files to the clipboard as markdown.

## Install

Download from [releases](https://github.com/dropsite-ai/mdcopy/releases) or manually install:

```bash
git clone git@github.com:dropsite-ai/mdcopy.git
cd mdcopy
make install
```

## Usage

```bash
$ mdcopy -h
Usage of mdcopy:
  -copy
    	Copy results to clipboard (default true)
  -dir string
    	Change starting directory (default ".")
  -ext string
    	Comma-separated file extensions (e.g. go,txt)
  -match string
    	Comma-separated substrings that paths must match
  -unmatch string
    	Comma-separated substrings that paths must not match
  -verbose
    	Enable verbose logs for non-matching paths
```

## Test

```bash
make test
```

## Release

```bash
git tag -a v0.1.0 -m "Release description"
git push origin v0.1.0
goreleaser release
```