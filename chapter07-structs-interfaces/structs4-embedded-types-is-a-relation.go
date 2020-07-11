package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person // Using this way, fields and methods from Person will be directly acessible
	// so, in this case, Android is-a Person
	Model string
}

func main() {
	a := new(Android)
	a.Name = "Android"
	a.Person.Talk() // Two ways to call Talk()
	a.Talk()
}
