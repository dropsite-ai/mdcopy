package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dropsite-ai/mdcopy"
)

func main() {
	copyFlag := flag.Bool("copy", true, "Copy results to clipboard")
	dirFlag := flag.String("dir", ".", "Change starting directory")
	extFlag := flag.String("ext", "", "Comma-separated file extensions (e.g. go,txt)")
	matchFlag := flag.String("match", "", "Comma-separated substrings that paths must match")
	unmatchFlag := flag.String("unmatch", "", "Comma-separated substrings that paths must not match")

	flag.Parse()

	out, err := mdcopy.Run(
		*copyFlag,
		*dirFlag,
		*extFlag,
		*matchFlag,
		*unmatchFlag,
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
