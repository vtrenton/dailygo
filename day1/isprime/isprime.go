package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter a single number to test.")
		os.Exit(1)
	}

	in, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Numbers only please!")
		os.Exit(1)
	}

	result, err := isprime(in)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if result {
		fmt.Printf("%d is prime\n", in)
	} else {
		fmt.Printf("%d is not prime\n", in)
	}
}

func isprime(test int) (bool, error) {
	if test < 2 {
		return false, fmt.Errorf("Numbers less than 2 are never prime!")
	}

	factlimit := test / 2 // the lowest factor of a number is half

	for i := 2; i <= factlimit; i++ {
		if test%i == 0 {
			return false, nil
		}
	}
	return true, nil
}
