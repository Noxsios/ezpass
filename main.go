package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"

	"github.com/noxsios/ezpass/words"
)

func main() {
	var length int
	flag.IntVar(&length, "l", 4, "Number of words to use")

	var delimiter string
	flag.StringVar(&delimiter, "d", ".", "Delimiter between words")

	var help bool
	flag.BoolVar(&help, "h", false, "Print this message and exit.")

	flag.Parse()

	if len(os.Args[1:]) > 0 || help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	max := big.NewInt(int64(len(words.ALL)))

	for i := range length {
		randBigInt, err := rand.Int(rand.Reader, max)
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, "error: %s", err.Error())
			if err != nil {
				// truly panic now
				panic(err)
			}
		}

		index := int(randBigInt.Int64())

		fmt.Print(words.ALL[index])
		if i < length-1 {
			fmt.Print(delimiter)
		}
	}
}
