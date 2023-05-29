package main

import "fmt"

func main() {
	/*
				string, int, float, bool, structs are not directly updatable by calling another function
		these are value types
			you need to pass pointer to update the value
	*/
	i := 7
	updateValue(&i)
	fmt.Println(i)
}

func updateValue(value *int) int {
	*value += 7
	return *value
}
