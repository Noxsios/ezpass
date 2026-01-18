// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present Harry Randazzo

package words

import (
	"fmt"

	"github.com/spf13/pflag"
)

// Wordlist represents a wordlist type that can be used for password generation.
type Wordlist string

var _ pflag.Value = (*Wordlist)(nil)

const (
	// EffWordlist represents the EFF wordlist.
	EffWordlist Wordlist = "eff"
	// UsrShareDictWordlist represents the system dictionary wordlist.
	UsrShareDictWordlist Wordlist = "usr_share_dict"
)

// AvailableWordlists returns a slice of all available wordlist names.
func AvailableWordlists() []string {
	return []string{
		string(EffWordlist),
		string(UsrShareDictWordlist),
	}
}

// Set implements the pflag.Value interface to allow setting the wordlist from command line flags.
func (w *Wordlist) Set(s string) error {
	switch s {
	case string(EffWordlist):
		*w = EffWordlist

		return nil
	case string(UsrShareDictWordlist):
		*w = UsrShareDictWordlist

		return nil
	default:
		return fmt.Errorf("%q is not a valid wordlist", s)
	}
}

func (w *Wordlist) String() string {
	return string(*w)
}

// Type implements the pflag.Value interface to return the type name for help text.
func (w *Wordlist) Type() string {
	return "wordlist"
}
