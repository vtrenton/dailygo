package main

import (
	"fmt"
	"sync"
)

func main() {
	// set up waitgroup
	var wg sync.WaitGroup
	wg.Add(2)

	input := []int{3, 8, 16, 4, 8, 15, 11, 1}

	go sumslice(input, &wg)
	go poweroftwoslice(input, &wg)

	wg.Wait()
}

func sumslice(s []int, wg *sync.WaitGroup) {
	defer wg.Done()
	var sum int
	for _, v := range s {
		sum += v
	}
	fmt.Println(sum)
}

func poweroftwoslice(s []int, wg *sync.WaitGroup) {
	defer wg.Done()
	var resultslice []int
	for _, v := range s {
		counter := 1
		for i := 0; i < v; i++ {
			counter *= 2
		}
		resultslice = append(resultslice, counter)
	}
	fmt.Println(resultslice)
}
