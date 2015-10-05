package main

import "fmt"

func main() {
	var x int32 = 42
	y := &x

	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %p\n", &x)
	fmt.Printf("Value of pointer y: %d\n", *y)
	fmt.Printf("Address of y: %p\n", y)
}
