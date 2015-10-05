package main

import "fmt"

func main() {
	type Person struct {
		firstname, name string
	}

	// assign after
	var shenron *Person = new(Person)
	shenron.name = "CASTELLY"
	shenron.firstname = "Thomas"

	// create with values
	var toto *Person = &Person{"Toto", "TOTO"}

	fmt.Printf("shenron firstname: %s\n", shenron.firstname)
	fmt.Printf("shenron name: %s\n", shenron.name)
	fmt.Printf("\n")
	fmt.Printf("toto firstname: %s\n", toto.firstname)
	fmt.Printf("toto name: %s\n", toto.name)
	fmt.Printf("\n")
}
