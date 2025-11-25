// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present Harry Randazzo

// Package words contains the generated wordlist and passphrase printer
package words

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
)

//go:generate go run ../gen/main.go

// PrintEzpass prints to wr a random passphrase of n words separated by a given delimiter.
func PrintEzpass(wr io.Writer, n int, delimiter string) error {
	upper := big.NewInt(int64(len(ALL)))

	for i := range n {
		randBigInt, err := rand.Int(rand.Reader, upper)
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

// PrintEzpassEFF prints to wr a random passphrase of n words separated by a given delimiter using the EFF wordlist.
func PrintEzpassEFF(wr io.Writer, n int, delimiter string) error {
	for i := range n {
		roll, err := Roll5Dice()
		if err != nil {
			return err
		}

		fmt.Fprint(wr, EFF[roll])

		if i < n-1 {
			fmt.Fprint(wr, delimiter)
		}
	}

	return nil
}
