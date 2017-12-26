package main

import (
	"fmt"
	"testfunctions"
)


func main() {
	
//	var age int; // Variable declaration with static data type
	
//	var name = "Pavan"; // Type inference, type will be inferred as string
//	var age = 29; // Type inference, type will be inferred as string

//	address, city := "varthur", "Bengaluru";


	var (
		age = 29;
		name = "Pavan"
		address = "varthur"
		city = "Bengaluru"
		pincode string
	);

	pincode = "560087"

//	pincode = 10;	// Error since pincode is of type string and we are assiging integer

    	fmt.Printf("Name: %s\nAge: %d\nAddress: %s\nCity: %s\npincode: %s\n", name, age, address, city, pincode);

	testfunctions.Display()


//	testfunctions.localFunction() // Non exported function, cannot be accessed

}

