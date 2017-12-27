// package : Specify the package name to which this go source file belongs to.
package testfunctions

import (
	"fmt"
)

func localFunction(){
	fmt.Println("This is a local function, cannot be called outside this source file");
}


// Function name should start with upper case letter, if we want to export this function to other packages.
func Display(){
	fmt.Println("\ndisplay function within testfunctions package\n");

	localFunction();
}
