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
	//fmt.Println(colors["velvet"])
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
	fmt.Printf("%+v", jim)
}
