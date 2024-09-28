package main

import (
	"fmt"
	"strings"
)

func main() {
	instr := ""
	lwrstr := strings.ToLower(instr)

	fmt.Println("vim-go")
}

func fmtstr(s string) string { // in go strings are immutible - pass by reference would have no benifit here
	for i := 0; i < len(s); i++ {

