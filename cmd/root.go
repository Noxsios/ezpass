// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present Harry Randazzo

// Package cmd is the entrypoint for the ezpass CLI
package cmd

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	"github.com/noxsios/ezpass/words"
	"github.com/spf13/pflag"
)

// Main is the entry point for the ezpass CLI application.
// It processes command line arguments and returns an exit code.
func Main() int {
	ezFlags := pflag.NewFlagSet("ezpass", pflag.ExitOnError)
	ezFlags.SortFlags = false
	ezFlags.Usage = func() {}

	var n int
	ezFlags.IntVarP(&n, "number-of-words", "n", 4, "Number of words in resulting password.")

	var delimiter string
	ezFlags.StringVarP(&delimiter, "delimiter", "d", ".", "Delimiter between words.")

	var ver bool
	ezFlags.BoolVarP(&ver, "version", "v", false, "Print the version number of ezpass and exit.")

	var help bool
	ezFlags.BoolVarP(&help, "help", "h", false, "Print this message and exit.")

	if err := ezFlags.Parse(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "error parsing flags:", err)
		os.Exit(2)
	}

	delimiter = strings.Trim(delimiter, `"`)
	delimiter = strings.Trim(delimiter, "'")

	if help {
		fmt.Fprintln(os.Stderr, ezFlags.FlagUsages())
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
