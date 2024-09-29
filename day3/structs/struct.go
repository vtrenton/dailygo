package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	First    string
	Last     string
	Age      int
	Birthday string
	Weight   int
}

func (p *Person) Update(first string, last string) {
	reader := bufio.NewReader(os.Stdin)
	p.First = first
	p.Last = last

	fmt.Println("How old are you?")
	if n, err := fmt.Scanf("%d", &p.Age); err != nil || n != 1 {
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
	p.Birthday = birthday

	fmt.Println("How many lbs of muscle are you made of?")
	if n, err := fmt.Scanf("%d", &p.Weight); err != nil || n != 1 {
		fmt.Println("bad input! ", err)
		os.Exit(1)
	}
}

func (p Person) jsonify() string {
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(j)
}

func main() {
	me := Person{
		First:    "",
		Last:     "",
		Age:      0,
		Birthday: "",
		Weight:   0, // all muscle
	}

	me.Update("Trent", "VanderWert")
	fmt.Printf("Hi my name is %s %s, I'm %d years old and was born on %s, I weigh %dlbs\n", me.First, me.Last, me.Age, me.Birthday, me.Weight)

	jsonstr := me.jsonify()
	fmt.Println(jsonstr)
}
