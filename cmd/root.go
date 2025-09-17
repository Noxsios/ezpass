// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present Harry Randazzo

// Package cmd is the entrypoint for the ezpass CLI
package cmd

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
	"slices"
	"strings"

	"github.com/noxsios/ezpass/words"
)

func Main() int {
	ezFlags := flag.NewFlagSet("ezpass", flag.ExitOnError)

	var n int
	ezFlags.IntVar(&n, "n", 4, "Number of words in resulting password.")

	var delimiter string
	ezFlags.StringVar(&delimiter, "d", ".", "Delimiter between words.")

	var help bool
	ezFlags.BoolVar(&help, "h", false, "Print this message and exit.")

	var ver bool
	ezFlags.BoolVar(&ver, "v", false, "Print the version number of ezpass and exit.")

	if err := ezFlags.Parse(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "error parsing flags:", err)
		os.Exit(2)
	}

	delimiter = strings.Trim(delimiter, `"`)
	delimiter = strings.Trim(delimiter, "'")

	if help || slices.Contains(os.Args[1:], "--help") {
		ezFlags.PrintDefaults()
		return 0
	}

	if ver {
		bi, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Fprintln(os.Stderr, "version information not available")
			return 1
		}
		fmt.Println(bi.Main.Version)
		return 0
	}

	if n <= 0 || n >= len(words.ALL) {
		fmt.Fprintln(os.Stderr, "error: number of words must be 0 < n <", len(words.ALL), ", got:", n)
		return 1
	}

	if err := words.PrintEzpass(os.Stdout, n, delimiter); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err.Error())
		return 1
	}
	fmt.Print("\n")
	return 0
}
