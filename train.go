package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	var pass string
	fmt.Scan(&pass)
	m := []rune(pass)
	answ := true
	for j := 0; j < len(m); j++ {
		if unicode.Is(unicode.Latin, m[j]) == true {
			continue
		} else if unicode.IsDigit(m[j]) == true {
			continue
		} else {
			answ = false
		}
	}
	if utf8.RuneCountInString(pass) >= 5 && answ == true {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
