package main

import (
	"fmt"

	"github.com/noxsios/ezpass/words"
)

func main() {
	counts := map[int]int{}

	for _, word := range words.ALL {
		counts[len(word)] = counts[len(word)] + 1
	}

	fmt.Println(counts)
}
