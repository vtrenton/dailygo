package main

import "fmt"

func main() {
	av1 := average(1.55, 2.8643, 15.111334)
	av2 := average(4.67, 8.9032, 13.2, 3.889, 3.44, 8.35)
	numslice := []int{8, 3, 12, 31, 18, 15, 7, 11, 4}

	fmt.Printf("av1: %f and av2: %f\n", av1, av2)

	var evenfilter = func() []int {
		var odds []int
		for _, v := range numslice {
			if v%2 != 0 {
				odds = append(odds, v)
			}
		}
		return odds
	}
	fmt.Println(evenfilter())

	power := 8

	// function is returned and executed.
	// the int value of the returned function is what is stored
	// hence the end ()
	poweroftwo := twotothepowerof(power)()
	fmt.Printf("2 to the power of %d is %d\n", power, poweroftwo)
}

func average(a ...float64) float64 {
	var sum float64
	var average float64
	var counter int
	for i, v := range a {
		sum += v
		counter = i
	}
	average = sum / (float64(counter) + 1)
	return average
}

// Closure: return a function
// note return type is func() int
func twotothepowerof(n int) func() int {
	return func() int {
		// this will not work
		// because the ^ carrot is a bitwise or
		// not a power
		//return 2 ^ n

		//imagine importing math.Pow lol
		x := 1
		for i := 0; i < n; i++ {
			x *= 2
		}
		return x
	}
}
