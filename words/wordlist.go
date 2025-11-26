package words

import (
	"fmt"

	"github.com/spf13/pflag"
)

type Wordlist string

var _ pflag.Value = (*Wordlist)(nil)

const (
	EffWordlist          Wordlist = "eff"
	UsrShareDictWordlist Wordlist = "usr_share_dict"
)

func AvailableWordlists() []string {
	return []string{
		string(EffWordlist),
		string(UsrShareDictWordlist),
	}
}

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

func (w *Wordlist) Type() string {
	return "wordlist"
}
