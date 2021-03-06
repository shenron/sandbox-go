package main

import (
	"fmt"
	"time"
	"math/big"
	"crypto/rand"
)

/**
 * Fake calculation
 * return random int64
 */
func Calculate() int64 {
	c := big.NewInt(10)
	fakeResult, _ := rand.Int(rand.Reader, c)
	return fakeResult.Int64()
}

/**
 * Before exec the parameter function 
 * > wait fiew seconds
 * > update the chanel
 */
func HelpCalculate(c chan int64, fct func() int64) {	
	fmt.Println("- No problem, but I restart my WAS wait some hours ...")
	seconds := 5 
	time.Sleep(time.Duration(seconds) * time.Second)
	c <- fct()
}

func main() {
	fmt.Println("- Hey I need help, calculate this please !")

	c := make(chan int64)
	go HelpCalculate(c, Calculate)
	result := <- c

	fmt.Println(" ... Result ", result)
}
