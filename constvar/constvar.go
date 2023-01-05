package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.14159
	var radius float64 = 3

	area := PI * m.Pow(radius, 2)
	fmt.Println("area", area)

	var i, j = 1, 2
	fmt.Println("i,j", i, j)

	const n int = 0
	fmt.Println("n", n)

	var (
		a float32 = 1.2
		b int     = 1
		c string  = "hello"
	)
	fmt.Println("a,b,c", a, b, c)
}
