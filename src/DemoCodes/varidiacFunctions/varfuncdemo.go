package main

import "fmt"

// varidiac function to recieve a string arguments
func change (s ... string){
	s[0] = "Go"

	s = append(s, "India")
	fmt.Println("Inside change function ", s);
}


// find whether a num exists in the passed variadiac arguments
func find(num int, nums ... int){
	found := false
	if nums == nil {
		fmt.Println("No optional arguments missing");
		return
	}
	for i, v := range nums{
		if num == v{
			fmt.Println(num, " found at index ", i, " in ", nums);
			found = true
		}
	}
	if !found {
		fmt.Println(num, " not found")
	}
//	fmt.Printf("\n")
}


func main(){
	fmt.Println("Main function");
	find(10, 1,2,3,4)
	find(10, 10,2)
	find(100, 1,2,3,4,5,6,7)
	find(10)

	// create a new slice variable
	slice := []string{"Pavan", "Kumar"}

	//change(slice) --> results in error
	//cannot use slice (type []string) as type string in argument to change

	// Passing slice to vardiac function
	change(slice...)
	fmt.Println(slice)
}

