package main

import "fmt"

// Declaration: Method 1
type Employee struct {
	firstname string
	lastname string
	psno string
	age int
}
/*
// Declaration: Method 2
// Same type elements can be written in same line
type Employee struct {
	firstname, lastname, psno string
	age int
}
*/

// Display employee information using pointer variable
func DisplayEmpInfoPtr(empPtr *Employee ){

//	fmt.Println(*empPtr)	// Reference to entire structure.

	fmt.Println((*empPtr).firstname) // Reference to individual elements of structure.
	fmt.Println((*empPtr).lastname) // Reference to individual elements of structure.
	fmt.Println((*empPtr).psno) // Reference to individual elements of structure.
	fmt.Println((*empPtr).age) // Reference to individual elements of structure.
}

func DisplayEmpInfo(info Employee){
	fmt.Printf("\n%-15s: %s\n", "FirstName",info.firstname)
	fmt.Printf("%-15s: %s\n","LastName ", info.lastname)
	fmt.Printf("%-15s: %s\n","PS No", info.psno)
	fmt.Printf("%-15s: %d\n\n","Age", info.age)
}

func main(){
	fmt.Println("Structure demo")

	// Creating structure using field names
	// Assignment can be done in any order
	emp1 := Employee {
		psno: "20137453",
		firstname: "Pavan",
		age: 29,
		lastname: "Kumar",
	}

	// Creating structure without using field names
	// Works fine since assigned value types matches with declaration flow.
	// i.e string, string, string, int
	emp2 := Employee {"Pooja", "R", "20129022", 26}


	// Fails since assigned value types does not match with declaration flow.
	// i.e string, string, string, int != string, string, int, string 
	//emp2 := Employee {"Pooja", "20129022", 26, "R"}


	// Structure variable with no values initialized, zero values by default.
	emp3 := Employee {}

	_ = emp3
	// Nil value variable.
	var emp4 Employee

	emp4.firstname = "Sri"
	emp4.age = 29


/*	var empPtr *Employee;
	empPtr = &emp1;
*/
	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(emp3)
	fmt.Println(emp4)

	DisplayEmpInfo(emp1)
	DisplayEmpInfo(emp2)
	DisplayEmpInfo(emp3)
	DisplayEmpInfo(emp4)

	DisplayEmpInfoPtr(&emp1)
}
