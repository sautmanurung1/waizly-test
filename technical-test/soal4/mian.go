package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calculateProportions(arr []int) []string {
	n := len(arr)
	positiveCount := 0
	negativeCount := 0
	zeroCount := 0

	for i := 0; i < n; i++ {
		if arr[i] > 0 {
			positiveCount++
		} else if arr[i] < 0 {
			negativeCount++
		} else {
			zeroCount++
		}
	}

	positiveProportion := float64(positiveCount) / float64(n)
	negativeProportion := float64(negativeCount) / float64(n)
	zeroProportion := float64(zeroCount) / float64(n)

	return []string{
		strconv.FormatFloat(positiveProportion, 'f', 6, 64),
		strconv.FormatFloat(negativeProportion, 'f', 6, 64),
		strconv.FormatFloat(zeroProportion, 'f', 6, 64),
	}
}

func main() {
	input := "-4 3 -9 0 4 1"
	inputArr := strings.Split(input, " ")
	arr := make([]int, len(inputArr))

	for i, numStr := range inputArr {
		num, _ := strconv.Atoi(numStr)
		arr[i] = num
	}

	proportions := calculateProportions(arr)

	for i := 0; i < 3; i++ {
		fmt.Println(proportions[i])
	}
}
