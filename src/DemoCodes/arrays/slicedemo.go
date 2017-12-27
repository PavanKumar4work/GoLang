package main

import "fmt"

func main(){
	var a = [...]int{1,2,3,4,5,6,7};

	b := a[1:3];
	fmt.Println(a);
	fmt.Println(b);

/*	b[0] = 100;
	fmt.Println(a);
	fmt.Println(b);
*/

	c := a[4:6];
/*	c[0] = 101;
	fmt.Println(a);
	fmt.Println(b);


	fmt.Printf("length of a: %d\tcapacity of a: %d\n", len(a), cap(a));
	fmt.Printf("length of b: %d\tcapacity of b: %d\n", len(b), cap(b));
	fmt.Printf("length of c: %d\tcapacity of c: %d\n", len(c), cap(c));
*/
	b = append(b, 200);
	fmt.Println(a);
	fmt.Println(b);
	fmt.Printf("length of a: %d\tcapacity of a: %d\n", len(a), cap(a));
	fmt.Printf("length of b: %d\tcapacity of b: %d\n", len(b), cap(b));
	fmt.Printf("length of c: %d\tcapacity of c: %d\n", len(c), cap(c));

}
