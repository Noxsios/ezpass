// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2025-Present Defense Unicorns

package main

import (
	"os"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		"maru2": func() {
			os.Exit(Main())
		},
	})
}

func TestE2E(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir:                "testdata",
		RequireUniqueNames: true,
		// UpdateScripts:      true,
	})
}
