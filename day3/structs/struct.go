package main

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
	first_name string
	last_name  string
	age        int
	birthday   string
	weight     int
}

func (p *Person) Update(first string, last string) {
	reader := bufio.NewReader(os.Stdin)
	p.first_name = first
	p.last_name = last

	fmt.Println("How old are you?")
	if n, err := fmt.Scanf("%d", &p.age); err != nil || n != 1 {
		fmt.Println("bad input! ", err)
		os.Exit(1)
	}

	// strings gotta be handled a bit different
	fmt.Println("When is your birthday?")
	birthday, berr := reader.ReadString('\n')
	if berr != nil {
		fmt.Println("bad input! ", berr)
		os.Exit(1)
	}
	birthday = birthday[:len(birthday)-1]
	p.birthday = birthday

	fmt.Println("How many lbs of muscle are you made of?")
	if n, err := fmt.Scanf("%d", &p.weight); err != nil || n != 1 {
		fmt.Println("bad input! ", err)
		os.Exit(1)
	}
}

func main() {
	me := Person{
		first_name: "",
		last_name:  "",
		age:        0,
		birthday:   "",
		weight:     0, // all muscle
	}

	me.Update("Trent", "VanderWert")
	fmt.Printf("Hi my name is %s %s, I'm %d years old and was born on %s, I weigh %dlbs\n", me.first_name, me.last_name, me.age, me.birthday, me.weight)
}
