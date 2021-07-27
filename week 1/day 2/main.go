package main

import (
	"fmt"
	"math"
)

// Nilai 80 - 100: A
// Nilai 65 - 79: B
// Nilai 50 - 64: C
// Nilai 35 - 49: D
// Nilai 0 - 34: E

func hasil(nilai int) string {
	if nilai <= 100 && nilai >= 80 {
		return "A"
	} else if nilai <= 79 && nilai >= 65 {
		return "B"
	} else if nilai <= 64 && nilai >= 50 {
		return "C"
	} else if nilai <= 49 && nilai >= 35 {
		return "D"
	} else if nilai <= 34 && nilai >= 0 {
		return "E"
	}

	return "Nilai invalid"
}

func faktorBilangan(input int) {

	for i := 1; i <= input; i++ {
		if input%i == 0 {
			fmt.Println(i)
		}
	}
}

func bilanganPrima(num float64) string {
	x := 0

	var value float64 = math.Floor(num / 2)
	for i := 2; i < int(value); i++ {
		x++
		if int(num)%i == 0 {
			return "Bukan bilangan prima"
		}
	}
	return "Bilangan Prima"
}

func palindrome(kata string) string {

	var kebalikan string

	for i := len(kata) - 1; i >= 0; i-- {
		kebalikan = kebalikan + string(kata[i])
	}
	if kebalikan == kata {
		return "Palindrome"
	}
	return "Bukan Palindrome"
}

func pangkat(base, pangkat int) int {
	// your code here
	var hasil int = 1
	for i := 0; i < pangkat; i++ {
		hasil *= base
	}
	return hasil
}

func playWithAsterik(n int) {
	for baris := 0; baris <= n; baris++ {
		// memBuat sejumlah spasi
		for i := 1; i <= n-baris; i++ {
			fmt.Print(" ")
		}
		// menampilkan bintang
		for i := 1; i < 2*baris; i++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}

func cetakTablePerkalian(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			fmt.Print(i * j)
			if j*i > 9 {
				fmt.Print("   ")
			} else {
				fmt.Print("    ")
			}
		}
		fmt.Println()
	}
}

func main() {
	// no 1
	// var T, r float32
	// fmt.Print("input Tinggi: ")
	// fmt.Scan(&T)
	// fmt.Print("input radius: ")
	// fmt.Scan(&r)
	// var phi float32 = 3.14

	// var luas float32 = 2 * phi * r * (r * T)
	// fmt.Println("Luas Permukaan Tabung adalah", luas)

	// no 2
	// var input int
	// fmt.Print("input nilai :")
	// fmt.Scan(&input)
	// fmt.Println("Nilai anda : ", hasil(input))

	//no 3
	// var input int
	// fmt.Print("input bilangan :")
	// fmt.Scan(&input)

	// faktorBilangan(input)

	// no 4
	// var input float64
	// fmt.Print("input bilangan :")
	// fmt.Scan(&input)
	// fmt.Print(bilanganPrima(input))

	//no 5
	// var kata string
	// fmt.Print("masukkan kata: ")
	// fmt.Scan(&kata)

	// fmt.Print(palindrome(kata))

	// no 6
	// var x, n int

	// fmt.Print("input angka : ")
	// fmt.Scan(&x)
	// fmt.Print("input Pangkat : ")
	// fmt.Scan(&n)

	// fmt.Print(pangkat(x, n))

	//no 7
	playWithAsterik(5)

	//no 8
	// cetakTablePerkalian(9)

}
