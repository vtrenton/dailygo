package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []Person

// ByAge needs to implement the functions defined in the spec for sort.Interface
// https://pkg.go.dev/sort#Interface

func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Less(i, j int) bool {
	return b[i].Age < b[j].Age
}

func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	person := []Person{
		{"Trent", 32},
		{"Tom", 30},
		{"Jeff", 46},
		{"Eric", 38},
		{"John", 31},
	}

	fmt.Println("Original:")
	fmt.Println(person)

	sort.Sort(ByAge(person))
	fmt.Println("Sorted:")
	fmt.Println(person)
}
