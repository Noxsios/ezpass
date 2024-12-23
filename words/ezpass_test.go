package words_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/noxsios/ezpass/words"
)

func TestPrintEzpass(t *testing.T) {
	// will not error if nil is passed for the writer (or on any writer errors)
	err := words.PrintEzpass(nil, 0, "")
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	// this amount of runs takes ~12s on M2 Macbook Air
	for i := range 1000000 {
		t.Run(fmt.Sprintf("run-%d", i), func(t *testing.T) {
			t.Parallel()
			var buf strings.Builder
			err := words.PrintEzpass(&buf, 4, ".")
			if err != nil {
				t.Errorf("unexpected error: %s", err.Error())
			}
		})
	}
}
