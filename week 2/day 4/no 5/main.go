package main

import (
	"fmt"
	"math"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) Avarage() float64 {
	var avg float64

	for i := 0; i < len(s.score); i++ {
		avg += float64(s.score[i])
	}

	avg = avg / float64(len(s.score))
	return math.Round(avg*100) / 100

}

func (s Student) Min() (min int, name string) {

	min = s.score[0]
	name = s.name[0]
	for i := 1; i < len(s.score); i++ {
		if min > s.score[i] {
			min = s.score[i]
			name = s.name[i]
		}
	}

	return min, name

}

func (s Student) Max() (max int, name string) {
	max = s.score[0]
	name = s.name[0]
	for i := 1; i < len(s.score); i++ {
		if max < s.score[i] {
			max = s.score[i]
			name = s.name[i]
		}
	}

	return max, name
}

func main() {
	var a = Student{}

	for i := 0; i < 3; i++ {
		var name string
		fmt.Print("Input \t" + string(i) + " Student’s Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input \t" + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAvarage Score Students is", a.Avarage())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")
}
