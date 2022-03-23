package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(retrieveDecimalValue(reverse(strings.Split("10000000", ""))))
}

func retrieveDecimalValue(binaries []string) float64 {
	result := 0.0
	for position, value := range binaries {
		if value == "1" {
			result += math.Pow(2, (float64(position)))
		}
	}
	return result
}

func reverse(stringSlice []string) []string {
	last := len(stringSlice) - 1
	for i := 0; i < len(stringSlice)/2; i++ {
		stringSlice[i], stringSlice[last-i] = stringSlice[last-i], stringSlice[i]
	}
	return stringSlice
}
