package main

import (
	"fmt"
	"strings"
)

const pt = "abcdefghijklmnopqrstuvwxyz"
const cp = "zyxwvutsrqponmlkjihgfedcba"

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {

	encode := ""
	for _, v := range s.name {
		idx := strings.Index(pt, string(v))
		encode += string(cp[idx])
	}
	s.nameEncode = encode

	return encode
}

func (s *student) Decode() string {
	decode := ""
	for _, v := range s.nameEncode {
		idx := strings.Index(cp, string(v))
		decode += string(pt[idx])
	}
	s.name = decode

	return decode
}

func main() {
	var menu int
	var a = student{}
	var c Chiper = &a
	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)
	if menu == 1 {
		fmt.Print("\nInput Student's Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of Student’s Name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's Encode Name : ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of Student’s Name " + a.nameEncode + " is : " + c.Decode())
	} else {
		fmt.Println("Wrong input name menu")
	}
}
