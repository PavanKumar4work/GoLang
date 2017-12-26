package arithematics


import _ "fmt" // fmt module not used in this file

// Returns true for even, false for odd numbers

func EvenOdd(num int) bool{
	//if (num % 2 ){ 	// Not allowed, if expects only bool, this statement results in int type data
	
	if (num % 2 == 0){
		return true;
	}else{
		return false;
	}
}

