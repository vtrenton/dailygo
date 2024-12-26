package main

import (
	"fmt"
	"strings"
)

func main() {
	instr := "Imagine a society that subjects people to conditions that make them terribly unhappy, then gives them the drugs to take away their unhappiness. Science fiction? It is already happening to some extent in our own society... Instead of removing the conditions that make people depressed, modern society gives them antidepressant drugs. In effect, antidepressants are a means of modifying an individual's internal state in such a way as to enable him to tolerate social conditions that he would otherwise find intolerable."
	inslice := fmtstr(instr)
	//inslice := strings.Fields(instrfmt)
	//for _, v := range inslice {
	//	fmt.Println(v)
	//}

	m := make(map[string]int)
	for _, v := range inslice {
		m[v] = m[v] + 1
	}
	for key, val := range m {
		fmt.Printf("%s: %d\n", key, val)
	}
}

// This function lowercases and removes all special chars from string
// then creates a slice that splits on space and newlines
func fmtstr(s string) []string {
	var value strings.Builder
	var result []string

	lowers := strings.ToLower(s)
	for i := 0; i < len(lowers); i++ {
		curr := lowers[i]
		if 'a' <= curr && curr <= 'z' {
			value.WriteByte(curr)
		}
		// hmmm what if there is no \n as EOL
		if curr == ' ' || curr == '\n' {
			result = append(result, value.String())
			value.Reset() // reset the value
		}
	}
	return result
}
