package main

import "fmt"

func main() {

	//map implementation
	var colors map[string]string
	//colors := make(map[string]string)
	//colors := map[string]string{
	//	"red":   "#ff000",
	//	"green": "#ffdd9",
	//}
	colors["white"] = "#ffff"
	fmt.Println(colors)
	//fmt.Println(colors["velvet"])
}