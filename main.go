package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"

	"github.com/noxsios/ezpass/words"
)

func main() {
	length := 4
	separator := "."

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
			fmt.Print(separator)
		}
	}
}
