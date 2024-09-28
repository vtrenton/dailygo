package main

import (
	"fmt"
	"strings"
)

func main() {
	instr := "Imagine a society that subjects people to conditions that make them terribly unhappy, then gives them the drugs to take away their unhappiness. Science fiction? It is already happening to some extent in our own society... Instead of removing the conditions that make people depressed, modern society gives them antidepressant drugs. In effect, antidepressants are a means of modifying an individual's internal state in such a way as to enable him to tolerate social conditions that he would otherwise find intolerable."
	instrfmt := fmtstr(instr)
	inslice := strings.Fields(instrfmt)
	//fmt.Println(inslice)

	m := make(map[string]int)
	for _, v := range inslice {
		m[v] = m[v] + 1
	}
	fmt.Println(m)
}

// This function lowercases and removes all special chars from string
// in go strings are immutable
// pass by reference would have no benefit here
func fmtstr(s string) string {
	var result strings.Builder
	lowers := strings.ToLower(s)
	for i := 0; i < len(lowers); i++ {
		curr := lowers[i]
		if ('a' <= curr && curr <= 'z') || curr == ' ' {
			result.WriteByte(curr)
		}
	}
	return result.String()
}
