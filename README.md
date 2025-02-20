# mdcopy

Copies files to the clipboard as markdown.

## Installation

### Homebrew (macOS or Compatible)

If you use Homebrew, install mdcopy with:
```bash
brew tap dropsite-ai/homebrew-tap
brew install mdcopy
```

### Download Binaries

Grab the latest pre-built binaries from the [GitHub Releases](https://github.com/dropsite-ai/mdcopy/releases). Extract them, then run the `mdcopy` executable directly.

### Build from Source

1. **Clone the repository**:
   ```bash
   git clone https://github.com/dropsite-ai/mdcopy.git
   cd mdcopy
   ```
2. **Build using Go**:
   ```bash
   go build -o mdcopy cmd/main.go
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