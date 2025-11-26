// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present Harry Randazzo

package words

import (
	"crypto/rand"
	"math/big"
)

// Roll5Dice simulates rolling five dice and getting a singular numeric representation of the rolls.
func Roll5Dice() (int, error) {
	upper := big.NewInt(int64(5))
	one := big.NewInt(int64(1))
	result := 0

	for range 5 {
		randBigInt, err := rand.Int(rand.Reader, upper)
		if err != nil {
			return 0, err
		}
		die := int(randBigInt.Add(randBigInt, one).Int64())
		result = result*10 + die
	}

	return result, nil
}
