package main

import "fmt"

func main() {
	input := []int{100, 5, 12, 55, 55, 15, 22, 12, 5, 8, 63, 22, 1, 12}
	fmt.Println(revarray(input))
	fmt.Println(rmdup(input))
}

func revarray(input []int) []int {
	outslice := []int{}
	for i := len(input) - 1; i >= 0; i-- {
		outslice = append(outslice, input[i])
	}
	return outslice
}

func rmdup(input []int) []int {
	for i, v := range input {
		for m := i + 1; m < len(input); m++ {
			if v == input[m] {
				input = append(input[:m], input[m+1:]...)
			}
		}
	}
	return input
}
