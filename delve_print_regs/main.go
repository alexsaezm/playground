package main

import "fmt"

//go:noinline
func test() (result float64) {
	a := 43843.343
	b := 9549.439
	result = a + b
	return
}

func main() {
	fmt.Println("c: ", test())
}
