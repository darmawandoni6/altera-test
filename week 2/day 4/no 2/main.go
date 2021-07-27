package main

import (
	"fmt"
	"strings"
)

func caesar(offset int, input string) string {
	const abjad = "abcdefghijklmnopqrstuvwxyz"
	var str string

	pos := offset

	if pos > len(abjad)-1 {
		pos = offset % len(abjad)
	}

	for i := 0; i < len(input); i++ {
		idx := strings.Index(abjad, string(input[i])) + pos
		if idx > len(abjad)-1 {
			idx -= len(abjad)
		}
		str += string(abjad[idx])
	}
	return str
}

func main() {

	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl

}
