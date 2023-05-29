package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	//map implementation
	//var colors map[string]string
	//colors := make(map[string]string)
	//colors := map[string]string{
	//	"red":   "#ff000",
	//	"green": "#ffdd9",
	//}
	//colors["white"] = "#ffff"
	//fmt.Println(colors)
	//fmt.Println(colors["velvet"])VVV
	//alex := person{firstName: "Alex", lastname: "Anderson"}
	//var alex person
	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"
	//fmt.Println(alex)
	//fmt.Printf("%+v", alex)
	jim := person{
		firstName: "jim",
		lastName:  "kong",
		contact: contactInfo{
			email:   "jim@kong.com",
			zipCode: 3400,
		},
	}
	//jimPointer := &jim // get the memory address of jim
	//jimPointer.updateName("kim")

	jim.updateName("kim") // in go, you can pass as a person or as a memory address
	jim.print()
}

// receiver function to print the info of a person
func (p person) print() {
	fmt.Printf("%+v", p)
}

// *person, is a type description, it means we're working with a pointer to a person
func (p *person) updateName(firstName string) {
	// *pointerToPerson is an operator, it means we want to manipulate the value the pointer is referencing
	(*p).firstName = firstName
}
