package mdcopy

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/atotto/clipboard"
	ignore "github.com/sabhiram/go-gitignore"
)

// Run scans 'dir' (defaults to "."), respecting .gitignore and hidden dirs,
// then gathers matching files in Markdown fences. If copyFlag is true, results
// go to the clipboard. Returns the full Markdown string (and any walk error).
// Note: The ext, match, and unmatch filters are now provided as slices.
func Run(copyFlag bool, dir string, exts, matches, unmatches []string, cmd bool) (string, error) {
	if dir == "" {
		dir = "."
	}
	ig := loadGitignore(filepath.Join(dir, ".gitignore"))

	var out strings.Builder
	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // skip problematic paths
		}
		relPath, _ := filepath.Rel(dir, path)
		relPath = filepath.ToSlash(relPath)

		if shouldSkip(relPath, info, ig, cmd) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if !info.IsDir() && passFilters(relPath, exts, matches, unmatches, cmd) {
			if cmd {
				fmt.Println("Matched file:", relPath)
			}

			out.WriteString(fmt.Sprintf("\nFile: %s\n```%s\n", relPath, langID(filepath.Ext(path))))
			appendFile(path, &out, cmd)
			out.WriteString("```\n")
		}
		return nil
	}); err != nil {
		return "", err
	}
	if copyFlag {
		if err := clipboard.WriteAll(out.String()); err != nil && cmd {
			return out.String(), err
		}
	}
	return out.String(), nil
}

func loadGitignore(path string) *ignore.GitIgnore {
	if _, err := os.Stat(path); err == nil {
		if ig, err := ignore.CompileIgnoreFile(path); err == nil {
			return ig
		}
	}
	return nil
}

func shouldSkip(relPath string, info os.FileInfo, ig *ignore.GitIgnore, cmd bool) bool {
	if ig != nil && ig.MatchesPath(relPath) {
		if cmd {
			fmt.Println("Skipping (gitignore):", relPath)
		}
		return true
	}
	// Skip hidden directories (except ".")
	if info.IsDir() && relPath != "." && strings.HasPrefix(filepath.Base(relPath), ".") {
		if cmd {
			fmt.Println("Skipping hidden directory:", relPath)
		}
		return true
	}
	return false
}

func passFilters(path string, exts, matches, unmatches []string, cmd bool) bool {
	// If exts given, file extension must be in that list.
	if len(exts) > 0 {
		fileExt := strings.TrimPrefix(filepath.Ext(path), ".")
		if !contains(exts, fileExt) {
			if cmd {
				fmt.Println("No match (ext):", path)
			}
			return false
		}
	}
	// If match filters are provided, accept if at least one is found.
	if len(matches) > 0 {
		matchFound := false
		for _, m := range matches {
			if strings.Contains(path, m) {
				matchFound = true
				break
			}
		}
		if !matchFound {
			if cmd {
				fmt.Println("No match (none of the match patterns found):", path)
			}
			return false
		}
	}
	// Exclude if any unmatch filter is present.
	for _, u := range unmatches {
		if strings.Contains(path, u) {
			if cmd {
				fmt.Printf("No match (forbidden '%s'): %s\n", u, path)
			}
			return false
		}
	}
	return true
}

func appendFile(path string, out *strings.Builder, cmd bool) {
	f, err := os.Open(path)
	if err != nil {
		if cmd {
			fmt.Println("Open error:", err)
		}
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		out.WriteString(sc.Text() + "\n")
	}
	if err := sc.Err(); err != nil && cmd {
		fmt.Println("Read error:", err)
	}
}

func langID(ext string) string {
	ext = strings.ToLower(strings.TrimPrefix(ext, "."))
	switch ext {
	case "js":
		return "javascript"
	case "ts":
		return "typescript"
	case "py":
		return "python"
	case "rb":
		return "ruby"
	case "sh":
		return "bash"
	case "md":
		return "markdown"
	default:
		if ext == "" {
			return "text"
		}
		return ext
	}
}

func contains(ss []string, val string) bool {
	for _, s := range ss {
		if s == val {
			return true
		}
	}
	return false
}
