package main

import (
	"fmt"
)

func getMinMax(numbers ...*int) (min int, max int) {
	// input
	mi := *numbers[0]
	mx := *numbers[0]

	for i := 1; i < len(numbers); i++ {
		if *numbers[i] < mi {
			mi = *numbers[i]
			continue
		}
		if *numbers[i] > mx {
			mx = *numbers[i]
			continue
		}
	}
	return mi, mx
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Print("input : ")
	fmt.Scan(&a1)
	fmt.Print("input : ")
	fmt.Scan(&a2)
	fmt.Print("input : ")
	fmt.Scan(&a3)
	fmt.Print("input : ")
	fmt.Scan(&a4)
	fmt.Print("input : ")
	fmt.Scan(&a5)
	fmt.Print("input : ")
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai min ", min)
	fmt.Println("Nilai max ", max)
}
