package main

import "fmt"

func main(){
	var a[3] int;		// Array a with 3 integer elements (Standard declaration)
	b := [3]int{1,2,3};	// Array b with 3 integer elements with values initialized (Short hand declaration)

	c := [...]int {1,2,3,4,5}	// Array c calculates the sizeof array with the values assigned

	d := [3]int {10}	// Less values can be initialized remaining elements is assigned with zero

	// Initailizing array a.
	a[0] = 1;
	a[1] = 2;
	a[2] = 3;

	fmt.Println("a: ", a);	//Print array a
	fmt.Println("b: ", b);	//Print array b
	fmt.Println("c: ", c);	//Print array c
	fmt.Println("d: ", d);	//Print array d

	a = b;	// Arrays of same can be assigned to each other
	fmt.Println("a = b: ", a);	//Print array a

//	c = b;	// Error since c and b are different data type, in go code array size is also part of data type
	// cannot use b (type [3]int) as type [5]int in assignment


	fmt.Println("Iterating array elements using len method");
	// Iterating array elements: Method 1
	for i := 0; i < len(a); i++{
		fmt.Printf("%d element of a: %d\n", i, a[i])
	}

	fmt.Println("Iterating array elements using range method");
	// Iterating array elements: Method 2 (using Range Method)
	for i, v := range a {
		fmt.Printf("%d element of a: %d\n", i, v)
	}

}
