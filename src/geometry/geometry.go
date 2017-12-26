//geometry.go
package main

import (
	"fmt"
	"geometry/rectangle"
	"log"
	"arithematics"
)

// Package variables
var rectLen, rectWidth float64 = 6, 7

// Init function to check whether length and width are positive integers
func init(){
	println("Main package initialized");
	if(rectLen < 0){
		log.Fatal("Length is less than zero");
	}

	if(rectWidth < 0){
		log.Fatal("Width is less than zero");
	}

}

func main(){

	if arithematics.EvenOdd(int(rectLen)){
		fmt.Println("rectLen is even number");
	}else{
		fmt.Println("rectLen is odd number");
	}

	fmt.Printf("Area of rectangle: %.2f\n", rectangle.Area(rectLen, rectWidth));
	fmt.Printf("Diagonal of rectangle: %.2f\n", rectangle.Diagonal(rectLen, rectWidth));

}


