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
	//alex := person{"Alex", "Anderson"}
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var alex1 person
	fmt.Println(alex1)
	fmt.Printf("%+v", alex1)

	alex1.firstName = "Alex1"
	alex1.lastName = "Anderson1"
	fmt.Printf("%+v", alex1)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
