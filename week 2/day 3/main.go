package main

import (
	"fmt"
	"math"
	"strconv"
)

func primeNumber(number int) bool {
	for i := 2; i < int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true

}

func pow(x, n int) int {
	// your code here
	var hasil int = 1
	for i := 0; i < n; i++ {
		hasil *= x
	}
	return hasil
}

func ArrayMerge(arrayA, arrayB []string) []string {

	x := append(arrayA, arrayB...)
	return x
}

func munculSekali(angka string) []int {
	// your code here
	var result []int = []int{}
	var unix int
	for i := 0; i < len(angka); i++ {
		for y := 0; y < len(angka); y++ {
			if string(angka[i]) == string(angka[y]) {
				unix++
			}
		}
		if unix < 2 {
			angka, _ := strconv.Atoi(string(angka[i]))
			result = append(result, angka)
		}
		unix = 0
	}
	return result
}

func PairSum(arr []int, target int) []int {
	// your code here
	var result []int = []int{}

	for i := 0; i < len(arr); i++ {
		for y := i + 1; y < len(arr); y++ {
			if arr[i]+arr[y] == target {
				result = append(result, i)
				result = append(result, y)
				return result

			}
		}

	}
	return result
}

func main() {

	// fmt.Println(primeNumber(20)) // true
	// fmt.Println(pow(10, 5)) // 100000

	// Test cases
	// fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	// fmt.Println(munculSekali("12345"))
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12)) // [2, 3]

}
