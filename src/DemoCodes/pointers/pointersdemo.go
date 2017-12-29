package main

import "fmt"

func change(val *int){
	*val = 100

	fmt.Println(*val)
}

func main(){
	b := 255;

	var ptr *int = &b;
	//var ptr *int; // Declaration without initialization, ptr contains nil by default which on dereference results in run time error

	if ptr == nil {
		fmt.Println("Pointer not assigned with valid address");
	}else{
		fmt.Printf("b: %-20d\t&b: %-20p\n", b, &b)
		fmt.Printf("ptr: %-20p\t*ptr: %-20d\n", ptr, *ptr)

		*ptr++	// pointer value incrementaion (value pointed by pointer)
		fmt.Printf("b: %-20d\t&b: %-20p\n", b, &b)
		fmt.Printf("ptr: %-20p\t*ptr: %-20d\n", ptr, *ptr)

//		ptr++	// Invalid operation, pointer arithematics not allowed in golang.


		//change(&b)
		change(ptr)
		fmt.Printf("b: %-20d\t&b: %-20p\n", b, &b)
		fmt.Printf("ptr: %-20p\t*ptr: %-20d\n", ptr, *ptr)
	}


}
