package main

import "fmt"

// Function to display key value pair
func displayMap( msg string, info map[string]int ){
	fmt.Printf("%s Map variables content\n", msg)
	for key, value := range info{
		fmt.Printf("Key: %-10s\tValue: %-10d\n", key, value)
	}
	fmt.Printf("\n")
}


func main(){
	// Syntax: var variablename map [key] value
	var PersonSalary map[string]int
	PersonSalary = make(map[string]int)
	PersonSalary["Pavan"] = 100
	PersonSalary["Hemanth"] = 1000
	PersonSalary["Pooja"] = 1000000

	// Declaration with initialization
	var age = map[string]int {
		"Pavan" : 29,
		"Hemanth":  28,
		"Pooja": 25,
	}

	displayMap("PersonSalary", PersonSalary);
	displayMap("age", age);

	// Delete a map using the key value.
	delete(PersonSalary, "Pooja")
	displayMap("PersonSalary", PersonSalary);

	// Maps are reference types
	newPersonSalary := PersonSalary
	displayMap("newPersonSalary", newPersonSalary);

	// Value modified by new map variable will affect the old map variable
	newPersonSalary["Pavan"] = 1000
	displayMap("newPersonSalary", newPersonSalary);
	displayMap("PersonSalary", PersonSalary);

}
