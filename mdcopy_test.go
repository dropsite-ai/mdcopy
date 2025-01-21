package mdcopy_test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/dropsite-ai/mdcopy"
)

// TestRunFixture checks if Run picks up files recursively from testdata,
// respects .gitignore, and skips hidden directories.
func TestRunFixture(t *testing.T) {
	// Path to our fixture directory
	fixtureDir := filepath.Join("fixture")

	// We want `.go` and `.sh` files, skip anything with 'hidden' in path
	out, err := mdcopy.Run(
		false,      // copyFlag (no actual clipboard for this test)
		fixtureDir, // dir
		"go,sh",    // ext
		"",         // match
		"hidden",   // unmatch
		false,      // verbose
	)
	if err != nil {
		t.Fatalf("mdcopy.Run returned an error: %v", err)
	}

	// Verify main.go is present
	if !strings.Contains(out, "File: main.go") {
		t.Error("Expected main.go to be included, not found in output")
	}

	// Verify script.sh is found inside subdir
	if !strings.Contains(out, "File: subdir/script.sh") {
		t.Error("Expected script.sh from subdir to be included, not found in output")
	}

	// notes.txt should be ignored by .gitignore
	if strings.Contains(out, "notes.txt") {
		t.Error("notes.txt should be ignored by .gitignore, but appeared in output")
	}

	// hidden/secret.go should be excluded by the 'hidden' unmatch filter
	if strings.Contains(out, "secret.go") {
		t.Error("secret.go was in a hidden directory, should not appear in output")
	}
}
