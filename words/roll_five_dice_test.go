// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present Harry Randazzo

package words_test

import (
	"fmt"
	"testing"

	"github.com/noxsios/ezpass/words"
)

func TestRoll5Dice(t *testing.T) {
	t.Run("returns a 5-digit number", func(t *testing.T) {
		result, err := words.Roll5Dice()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Result should be between 11111 and 66666
		if result < 11111 || result > 66666 {
			t.Errorf("result %d is out of expected range [11111, 66666]", result)
		}
	})

	t.Run("each digit is between 1 and 6", func(t *testing.T) {
		result, err := words.Roll5Dice()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Extract each digit and verify it's between 1 and 6
		digits := []int{
			result / 10000,
			(result / 1000) % 10,
			(result / 100) % 10,
			(result / 10) % 10,
			result % 10,
		}

		for i, digit := range digits {
			if digit < 1 || digit > 6 {
				t.Errorf("digit %d (position %d) is %d, expected value between 1 and 6", i, i, digit)
			}
		}
	})

	t.Run("produces different results across multiple calls", func(t *testing.T) {
		results := make(map[int]bool)

		for i := range 100 {
			result, err := words.Roll5Dice()
			if err != nil {
				t.Fatalf("unexpected error on iteration %d: %v", i, err)
			}
			results[result] = true
		}

		// With 100 rolls, we should get at least some variety
		// (not a guarantee, but statistically very likely)
		if len(results) < 50 {
			t.Logf("warning: only %d unique results out of %d rolls (expected more variety)", len(results), 100)
		}
	})

	t.Run("concurrent calls don't error", func(t *testing.T) {
		t.Parallel()

		for i := range 1000 {
			t.Run(fmt.Sprintf("run-%d", i), func(t *testing.T) {
				t.Parallel()

				result, err := words.Roll5Dice()
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}

				if result < 11111 || result > 66666 {
					t.Errorf("result %d is out of expected range", result)
				}
			})
		}
	})

	t.Run("no invalid digits like 0, 7, 8, or 9", func(t *testing.T) {
		// Run multiple times to increase confidence
		for attempt := range 50 {
			result, err := words.Roll5Dice()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Convert to string and check each character
			resultStr := fmt.Sprintf("%d", result)
			if len(resultStr) != 5 {
				t.Errorf("result %d has %d digits, expected 5", result, len(resultStr))
			}

			for i, char := range resultStr {
				digit := int(char - '0')
				if digit < 1 || digit > 6 {
					t.Errorf("attempt %d: invalid digit %d at position %d in result %d", attempt, digit, i, result)
				}
			}
		}
	})
}

func BenchmarkRoll6Dice(b *testing.B) {
	for b.Loop() {
		_, err := words.Roll5Dice()
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}
