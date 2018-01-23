package main

import "fmt"

func main() {
	one()
	two()
	three()
	four()
}
func one() {
	println("===Question One===")
	for a := 0; a < 10; a++ {
		fmt.Println(a)
	}
}
func two() {
	println("===Question Two===")
	b := 0
Here:
	println(b)
	b++
	if b < 10 {
		goto Here
	}
}
func three() {
	println("===Question Three===")
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for _, value := range a {
		println(value)
	}
}
func four() {
	println("===Question Four===")
	for a := 0; a < 100; a++ {
		switch {
		case a%3 == 0 && a%5 == 0:
			println(a)
			println("BizzFuzz")
			continue
		case a%5 == 0:
			println(a)
			println("Fuzz")
			continue
		case a%3 == 0:
			println(a)
			println("Bizz")
		}
	}
}
