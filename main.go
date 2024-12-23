// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present Harry Randazzo

// Package main is the entrypoint for the ezpass CLI
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
	"slices"

	"github.com/noxsios/ezpass/words"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 4, "Number of words in resulting password.")

	var delimiter string
	flag.StringVar(&delimiter, "d", ".", "Delimiter between words.")

	var help bool
	flag.BoolVar(&help, "h", false, "Print this message and exit.")

	var ver bool
	flag.BoolVar(&ver, "v", false, "Print the version number of ezpass and exit.")

	flag.Parse()

	if help || slices.Contains(os.Args[1:], "--help") {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if ver {
		bi, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Println("version information not available")
			os.Exit(1)
		}
		fmt.Println(bi.Main.Version)
		os.Exit(0)
	}

	if n <= 0 || n >= len(words.ALL) {
		fmt.Fprintln(os.Stderr, "error: number of words must be 0 < n <", len(words.ALL), ", got:", n)
		os.Exit(1)
	}

	if err := words.PrintEzpass(os.Stdout, n, delimiter); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err.Error())
		os.Exit(1)
	}
	fmt.Print("\n")
}
