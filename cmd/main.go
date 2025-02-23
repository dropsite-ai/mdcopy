package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dropsite-ai/mdcopy"
)

// stringSlice collects multiple flag values.
type stringSlice []string

func (s *stringSlice) String() string {
	return fmt.Sprint(*s)
}

func (s *stringSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func main() {
	copyFlag := flag.Bool("copy", true, "Copy results to clipboard")
	dirFlag := flag.String("dir", ".", "Change starting directory")

	var extFlags stringSlice
	flag.Var(&extFlags, "ext", "File extension filter. Specify multiple times for multiple extensions.")
	var matchFlags stringSlice
	flag.Var(&matchFlags, "match", "Substring filter for inclusion. Specify multiple times (file is included if any match).")
	var unmatchFlags stringSlice
	flag.Var(&unmatchFlags, "unmatch", "Substring filter for exclusion. Specify multiple times.")

	flag.Parse()

	out, err := mdcopy.Run(
		*copyFlag,
		*dirFlag,
		extFlags,
		matchFlags,
		unmatchFlags,
		true,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	// If not copying to clipboard, print result to stdout.
	if !*copyFlag {
		fmt.Print(out)
	}
}
