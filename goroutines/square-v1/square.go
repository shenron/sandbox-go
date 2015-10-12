package main

import "fmt"

/**
 * Generate chanel
 */
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

/**
 * Calculate square
 */
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	c := generate(2, 4)
	out := square(c)

	for i := 0; i < 2; i++ {
		fmt.Println(<-out)
	}
	// 4 16
}
