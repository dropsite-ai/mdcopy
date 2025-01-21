# mdcopy

Copies files to the clipboard as markdown.

## Install

Download from [Releases](https://github.com/dropsite-ai/mdcopy/releases):

```bash
tar -xzf mdcopy_Darwin_arm64.tar.gz
chmod +x mdcopy
sudo mv mdcopy /usr/local/bin/
```

Or manually build and install:

```bash
git clone git@github.com:dropsite-ai/mdcopy.git
cd mdcopy
make install
```

## Usage

```bash
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
```

## Test

```bash
make test
```

## Release

```bash
make release
```