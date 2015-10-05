package main

import (
	"fmt"
	"sandbox-go/oop/Child"
	"sandbox-go/oop/LifeCycle"
	"sandbox-go/oop/Parent"
)

func main() {
	fmt.Printf("Test OOP \n")
	GenericSleep := LifeCycle.GenericSleep

	p := Parent.New("Thomas", "CASTELLY")
	fmt.Printf("%s %s\n", p.GetFirstName(), p.GetName())
	fmt.Printf("slepp: %s\n", GenericSleep(p))
	fmt.Printf("\n")
	c := Child.New("MiniMoi", p)
	fmt.Printf("%s %s\n", c.GetFirstName(), c.GetName())
	fmt.Printf("sleep: %s\n", GenericSleep(c))

}
