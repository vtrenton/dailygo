package main

import "fmt"

func main() {
	a, b := 69, 420
	fmt.Printf("Pre Swap: A: %d and B: %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After Swap: A: %d and B: %d\n", a, b)
}

// x contains the address of a
// y contains the address of b
func swap(x *int, y *int) {
	m := *x //place the deref x (a) into int m as an interm placeholder
	*x = *y //take the defref value of y and assign it to the deref of x
	*y = m  //deref y and assign it to the intermediate m value
}
