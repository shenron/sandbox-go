package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func Grenade(c chan int64) {
	delay, _ := rand.Int(rand.Reader, big.NewInt(10))
	seconds := delay.Int64()
	fmt.Printf("Grenade : Boom after %d seconds\n", seconds)
	time.Sleep(time.Duration(seconds) * time.Second)
	c <- seconds
}

func main() {
	c1 := make(chan int64, 1)
	c2 := make(chan int64, 1)

	go Grenade(c1)
	go Grenade(c2)

	i := 0
	for {
		select {
		case r1 := <-c1:
			fmt.Printf("Boomed in %d seconds\n", r1)
		case r2 := <-c2:
			fmt.Printf("Boom in %d seconds\n", r2)
		default:
			i += 1
			time.Sleep(1 * time.Second)
			if i >= 10 {
				fmt.Println("TIME OUT\n", i)
				return
			} else {
				fmt.Printf("%d ...\n", i)
			}
		}
	}
}
