package main

import "fmt"

func printBytesHex(printstring string){
	for i := 0; i < len(printstring); i++ {
		fmt.Printf("%X ", printstring[i])
	}
	fmt.Printf("\n")
}

func printBytesDec(printstring string){
	for i := 0; i < len(printstring); i++ {
		fmt.Printf("%d ", printstring[i])
	}
	fmt.Printf("\n")
}

func printChar1(s string){

	// Range method will return the index and indiviual character in rune format for a given string
	for _, rune1 := range s {
		fmt.Printf("%c ", rune1)
	}
	fmt.Printf("\n")
}


func printChar(printstring string){

	/*
	A rune is a builtin type in Go and its the alias of int32. 
	rune represents a Unicode code point in Go. It does not matter how many bytes the code point occupies, 
	it can be represented by a rune.
	*/
	runes := []rune(printstring)

/*
	// This code results in wrong output for incorrect unicode (Not UTF-8)
	for i := 0; i < len(printstring); i++ {
		fmt.Printf("%c ", printstring[i])
	}
	fmt.Printf("\n")
*/
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
	fmt.Printf("\n")
}

func main(){

	var name string = "Pavan"
	address := "Señor"
	fmt.Println(name)
	fmt.Println(address)

	printBytesHex(name)
	printBytesDec(name)
	printChar(name)
	printBytesHex(address)
	printBytesDec(address)
	printChar1(address)
}
