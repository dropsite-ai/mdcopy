package mdcopy_test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/dropsite-ai/mdcopy"
)

func TestRunFixture(t *testing.T) {
	// Existing fixture test: tests extension and unmatch filtering.
	fixtureDir := filepath.Join("fixture")
	out, err := mdcopy.Run(
		false,                // copyFlag (no actual clipboard for this test)
		fixtureDir,           // dir
		[]string{"go", "sh"}, // exts: only .go and .sh files
		[]string{},           // match: none provided
		[]string{"hidden"},   // unmatch: skip paths containing "hidden"
		false,                // verbose flag
	)
	if err != nil {
		t.Fatalf("mdcopy.Run returned an error: %v", err)
	}

	if !strings.Contains(out, "File: main.go") {
		t.Error("Expected main.go to be included, not found in output")
	}

	if !strings.Contains(out, "File: subdir/script.sh") {
		t.Error("Expected script.sh from subdir to be included, not found in output")
	}

	if strings.Contains(out, "notes.txt") {
		t.Error("notes.txt should be ignored by .gitignore, but appeared in output")
	}

	if strings.Contains(out, "secret.go") {
		t.Error("secret.go was in a hidden directory, should not appear in output")
	}
}

func TestMatchFilters(t *testing.T) {
	// Test various scenarios of match filtering.
	// Using the same extensions as above to target main.go and subdir/script.sh.
	fixtureDir := filepath.Join("fixture")
	exts := []string{"go", "sh"}

	t.Run("no match filters", func(t *testing.T) {
		out, err := mdcopy.Run(false, fixtureDir, exts, []string{}, []string{}, false)
		if err != nil {
			t.Fatalf("Run error: %v", err)
		}
		// With no match filters, both files should appear.
		if !strings.Contains(out, "File: main.go") {
			t.Error("Expected main.go to be included when no match filters are set")
		}
		if !strings.Contains(out, "File: subdir/script.sh") {
			t.Error("Expected subdir/script.sh to be included when no match filters are set")
		}
	})

	t.Run("match 'main'", func(t *testing.T) {
		out, err := mdcopy.Run(false, fixtureDir, exts, []string{"main"}, []string{}, false)
		if err != nil {
			t.Fatalf("Run error: %v", err)
		}
		// Only main.go should match.
		if !strings.Contains(out, "File: main.go") {
			t.Error("Expected main.go to be included when matching 'main'")
		}
		if strings.Contains(out, "File: subdir/script.sh") {
			t.Error("Did not expect subdir/script.sh to be included when matching 'main'")
		}
	})

	t.Run("match 'sub'", func(t *testing.T) {
		out, err := mdcopy.Run(false, fixtureDir, exts, []string{"sub"}, []string{}, false)
		if err != nil {
			t.Fatalf("Run error: %v", err)
		}
		// Only script.sh in the subdir should match.
		if !strings.Contains(out, "File: subdir/script.sh") {
			t.Error("Expected subdir/script.sh to be included when matching 'sub'")
		}
		if strings.Contains(out, "File: main.go") {
			t.Error("Did not expect main.go to be included when matching 'sub'")
		}
	})

	t.Run("match 'main' OR 'sub'", func(t *testing.T) {
		out, err := mdcopy.Run(false, fixtureDir, exts, []string{"main", "sub"}, []string{}, false)
		if err != nil {
			t.Fatalf("Run error: %v", err)
		}
		// Both files should be included because the match filters are ORed.
		if !strings.Contains(out, "File: main.go") {
			t.Error("Expected main.go to be included when matching 'main' or 'sub'")
		}
		if !strings.Contains(out, "File: subdir/script.sh") {
			t.Error("Expected subdir/script.sh to be included when matching 'main' or 'sub'")
		}
	})

	t.Run("match filter that matches none", func(t *testing.T) {
		out, err := mdcopy.Run(false, fixtureDir, exts, []string{"nonexistent"}, []string{}, false)
		if err != nil {
			t.Fatalf("Run error: %v", err)
		}
		// Expect no files to be included.
		if strings.Contains(out, "File: main.go") || strings.Contains(out, "File: subdir/script.sh") {
			t.Error("Expected no files to be included when match filter doesn't match any file")
		}
	})
}
