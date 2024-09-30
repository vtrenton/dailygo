package main

import "fmt"

const PI float64 = 3.1415926535

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Cirle implementing Shape
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	// pi*r^2
	return PI * (c.radius * c.radius)
}

func (c Circle) Perimeter() float64 {
	//2*pi*r
	return 2.0 * PI * c.radius
}

// Rectangle implementing Shape
type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	//(length * width)^2
	half := r.length * r.width
	return half * half
}

func (r Rectangle) Perimeter() float64 {
	//(width * length) * 2
	return (r.width + r.length) * 2.0
}

func main() {
	c := Circle{radius: 6.9}
	r := Rectangle{length: 4.0, width: 3.0}

	// Circle
	fmt.Printf("The area of the Circle is: %f\nThe Perimeter of the Circle is: %f\n", c.Area(), c.Perimeter())

	// Rectangle
	fmt.Printf("The area of the Rectangle is: %f\nThe Perimeter of the Rectangle is: %f\n", r.Area(), r.Perimeter())

	// basic type assertion on empty interfaces
	var a interface{}
	a = 10
	v, b := a.(int) //get the value back and a bool representing if our assertion (int) was true
	if b {
		fmt.Printf("%d is an int!\n", v)
	}

	fmt.Println(getInterfaceType(1))
	fmt.Println(getInterfaceType("hello"))
	fmt.Println(getInterfaceType(4.8))
	fmt.Println(getInterfaceType(true))
}

func getInterfaceType(i interface{}) string {
	switch i.(type) {
	case int:
		return "The interface is a int!"
	case string:
		return "The interface is a string!"
	case float64:
		return "The interface is a float!"
	case bool:
		return "The interface is a bool!"
	default:
		return "not sure about that type!"
	}
}
