// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present Harry Randazzo

// Package main is the entrypoint for the ezpass CLI
package main

import (
	"os"

	"github.com/noxsios/ezpass/cmd"
)

func main() {
	os.Exit(cmd.Main())
}
