package rectangle

import (
	"math"
	"fmt"
)

func init(){
	fmt.Println("rectangle package initialized\n");
}


func Area(len, width float64)(float64){
	area := len * width;
	return area;
}

func Diagonal(len, width float64) float64{
	diagonal := math.Sqrt((len * len) + (width * width));
	return diagonal;

}
