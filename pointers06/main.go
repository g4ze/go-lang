package main

import "fmt"

func main() {
	fmt.Print("Usage of pointers")
	myvar := 36069
	//pointer is like a variabke only but this thime there is no possibility of it being a "copy"
	//or duplicate reference
	var oe *int = &myvar //assigning a pointer to the vriable
	fmt.Println("Pointer location ", oe)
	fmt.Println("Pointer Vlaue ", *oe) // usage of astrick for pointer value
	myvar = *oe + 5
	*oe += 5
	fmt.Println("Variabke value", myvar)

}
