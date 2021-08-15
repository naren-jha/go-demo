package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	//contactInfo - we can also declare it like this as well
	// in this case member name becomes the type name, so it's like declaring
	// contactInfo contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex) // {Alex Anderson}

	// var alex person
	// fmt.Println(alex)       // { } - default value for string is ""
	// fmt.Printf("%+v", alex) // {firstName: lastName:}

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex) // {Alex Anderson}

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// fmt.Printf("%+v", jim) // {firstName:Jim lastName:Party contact:{email:jim@gmail.com zipCode:94000}}
	// jim.print()

	//jim.updateName("jimmy")
	//jim.print() // {firstName:Jim lastName:Party contact:{email:jim@gmail.com zipCode:94000}}
	// update didn't work..
	// why? pass by value.. pointers..

	pointerToJim := &jim
	pointerToJim.updateName("jimmy")
	jim.print() // {firstName:jimmy lastName:Party contact:{email:jim@gmail.com zipCode:94000}}

	// shortcut
	jim.updateName("aaja") // Go will do do auto conversion to pointer to person
	jim.print()            // {firstName:aaja lastName:Party contact:{email:jim@gmail.com zipCode:94000}}
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

/*
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
*/

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
