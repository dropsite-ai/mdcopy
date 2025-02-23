# mdcopy

Copies files to the clipboard as markdown.

## Introduction

**mdcopy** scans directories (honoring .gitignore and hidden directories) and puts file contents into Markdown-formatted code fences. It supports copying to clipboard, filtering files by extension, as well as inclusion (match) and exclusion (unmatch) filters.

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

### CLI Usage

```bash
  -copy
        Copy results to clipboard (default true)
  -dir string
        Change starting directory (default ".")
  -ext string
        File extension filter. Can be specified multiple times.
  -match string
        Substring filter for inclusion. Can be specified multiple times (file is included if any match).
  -unmatch string
        Substring filter for exclusion. Can be specified multiple times.
```

For instance, to scan the current directory for `.go` and `.sh` files that contain either "main" or "sub" in the path—but exclude paths containing "hidden"—you might run:

```bash
mdcopy -dir . -ext go -ext sh -match main -match sub -unmatch hidden
```

### Programmatic Usage

The `mdcopy` package can also be used directly within your Go code. First, import the package:

```go
import "github.com/dropsite-ai/mdcopy"
```

The primary function is:

```go
// Run scans the specified directory and returns a Markdown-formatted string.
// Parameters:
//   - copyFlag: if true, attempts to copy the result to the clipboard.
//   - dir: the root directory to scan.
//   - exts: a slice of file extensions (without the dot) to include.
//   - matches: a slice of substrings; a file is included if its path contains any of these.
//   - unmatches: a slice of substrings; if any is found in a file's path, that file is excluded.
//   - verbose: if true, prints debug output.
func Run(copyFlag bool, dir string, exts, matches, unmatches []string, verbose bool) (string, error)
```

Here’s a sample usage:

```go
package main

import (
	"fmt"
	"log"

	"github.com/dropsite-ai/mdcopy"
)

func main() {
	// Example: scan the current directory for .go and .txt files,
	// include only files that contain "example" in the path,
	// and exclude files that contain "test".
	output, err := mdcopy.Run(
		false,                 // copyFlag: don't copy to clipboard in this example
		".",                   // directory to scan
		[]string{"go", "txt"}, // file extensions to include
		[]string{"example"},   // match filters (OR: include if any match)
		[]string{"test"},      // unmatch filters (exclude if any found)
		true,                  // verbose output enabled
	)
	if err != nil {
		log.Fatalf("Error running mdcopy: %v", err)
	}
	fmt.Println(output)
}
```

## Test

To run the tests:

```bash
make test
```

The tests cover directory traversal, .gitignore handling, hidden directory skipping, and the new OR logic for match filtering.

## Release

To create a new release:

```bash
make release
```