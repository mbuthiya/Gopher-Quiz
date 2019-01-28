package main

import(
	"fmt"
	"flag"
)

//ValidateFlag allows us to validate if the user has inputed a flag Item
func ValidateFlag(flagName string)  bool{
	if flagName == ""{
		return false
	}
	return true;
}


func main(){

	// Creating flags
	var csvString string
	flag.StringVar(&csvString,"csv","","Add the csv link you want to connect")
	flag.Parse()

	fmt.Println(ValidateFlag(csvString))

}