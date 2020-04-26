package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	romain := person{
		firstName: "Romain",
		lastName:  "Testard",
		contactInfo: contactInfo{
			email:   "romain.rtestard@gmail.com",
			zipCode: 93210,
		}}

	/* romainPointer := &romain */
	romain.updateName("Rusty")
	romain.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson *person) print() {
	fmt.Printf("%+v", pointerToPerson)
}
