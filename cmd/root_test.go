// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present Harry Randazzo

package cmd_test

import (
	"os"
	"testing"

	"github.com/noxsios/ezpass/cmd"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		"ezpass": func() {
			os.Exit(cmd.Main())
		},
	})
}

func TestE2E(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir:                 ".",
		RequireUniqueNames:  true,
		RequireExplicitExec: true,
		// UpdateScripts:       true,
	})
}
