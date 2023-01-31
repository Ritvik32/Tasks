package main

import (
	"fmt"
)

var pub_key int
var priv_key int
var n int

//var input_string = "hi"

//var p = 3
//var q = 7
//var n = p * q
//var e = 2
func encode(s1 string) []int {
	arr := []int{}
	for i := 0; i < len(s1); i++ {
		arr = append(arr, enc(int(s1[i])))
	}
	return arr
}
func enc(s1 int) int {
	var a = 1

	for e := pub_key; e > 0; e-- {
		a *= s1
		a %= n
	}
	//a = a % n
	return a
}
func decode(n []int) string {
	var s string
	//t := encode(input_string)
	for i := 0; i < len(n); i++ {

		s += fmt.Sprintf("%c", dec(n[i]))
		//fmt.Print(s)
	}
	return s
}
func dec(s1 int) int {
	var a = 1
	for d := priv_key; d > 0; d-- {
		a = a * s1
		a = a % n
	}
	fmt.Print(a)
	return a
}
func gcd(a1, a2 int) int {
	var gcd1 int
	for i := 1; i <= a1 && i <= a2; i++ {
		if a1%i == 0 && a2%i == 0 {
			gcd1 = i
		}
	}
	return gcd1
} // s1 here is encrypted form of text
func main() {
	p := 13
	q := 17
	n = p * q
	e := 2
	phi := (p - 1) * (q - 1)
	for e < phi {
		if gcd(e, phi) == 1 {
			break
		} else {
			e++
		}
	}
	pub_key = e

	d := 2
	for {
		if (d*e)%phi == 1 {
			break
		}
		d++
	}
	priv_key = d

	input_string := "hi RJ"
	fmt.Println("Initial message is", input_string)
	fmt.Println("The encoded message is", encode(input_string))

	fmt.Println("decoded message is ", decode(encode(input_string)))
}
