package main

import "fmt"

type Person struct {
	firstname, name, initials string
}

func (p Person) GetFullName() (string, string) {
	return p.firstname, p.name
}

func (p Person) GetInitials() string {
	return p.initials
}

func main() {

	var shenron *Person = &Person{"Thomas", "CASTELLY", "TCY"}

	// multiple var in return
	firstname, name := shenron.GetFullName()

	fmt.Printf("Full name: %s %s\n", firstname, name)
	fmt.Printf("Initials: %s\n", shenron.GetInitials())
}
