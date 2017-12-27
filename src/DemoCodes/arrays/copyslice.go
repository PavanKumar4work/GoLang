package main

import "fmt"

func countries()[]string{
	countries := []string {"India", "Australia", "England", "Srilanka", "Pakistan"};
	neededCountries := countries[:3];

	countriescpy := make([]string, len(neededCountries));
	copy(countriescpy, neededCountries)
	return countriescpy
}


func main(){
	countriesNeeded := countries();
	fmt.Println(countriesNeeded);
}
