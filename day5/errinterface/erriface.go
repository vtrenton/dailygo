package main

import (
	"fmt"
)

type myErr interface {
	Error() string
}

type customErr struct {
	msg  string
	code int
}

func (e *customErr) Error() string {
	return fmt.Sprintf("Error: %s %d", e.msg, e.code)
}

func main() {
	s, err := errfunc()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf(s)
}

func errfunc() (string, error) {
	err := &customErr{msg: "File not found!", code: 404}
	if err != nil {
		return "", err
	}
	return "No Error", nil
}
