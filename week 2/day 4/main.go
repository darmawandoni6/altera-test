package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	// your code here

	fmt.Println(strings.Compare(a, b))

	return b
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI")) // AKA

}
