package words

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
)

func PrintEzpass(wr io.Writer, n int, delimiter string) error {
	max := big.NewInt(int64(len(ALL)))

	for i := range n {
		randBigInt, err := rand.Int(rand.Reader, max)
		if err != nil {
			return err
		}

		index := int(randBigInt.Int64())

		fmt.Fprint(wr, ALL[index])
		if i < n-1 {
			fmt.Fprint(wr, delimiter)
		}
	}
	return nil
}
