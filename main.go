package main

import "fmt"

func main() {
	var base int
	var altura int

	fmt.Scan(&base)
	fmt.Scan(&altura)

	area := (base * altura) / 2

	fmt.Print(area)
}
