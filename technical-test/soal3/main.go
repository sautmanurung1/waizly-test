package main

import (
	"fmt"
	"math/big"
)

func calculateSums() []*big.Int {
	sums := []*big.Int{}

	for i := 1; i <= 5; i++ {
		sum := big.NewInt(0)

		for j := 1; j <= 5; j++ {
			if i != j {
				sum.Add(sum, big.NewInt(int64(j)))
			}
		}

		sums = append(sums, sum)
	}
	return sums
}

func main() {
	result := calculateSums()
	fmt.Println(result)
}
