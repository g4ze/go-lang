package main

import "fmt"

func main() {
	fmt.Println("Structs is kinda similar to classes in golang but non inheritable; bviousky kind of a data structure")
	Animal := Animals{"Catto", 2}
	fmt.Printf("details of the animal is %+v\n", Animal)
	fmt.Printf("name of the animal is %v\n", Animal.Name)
}

type Animals struct {
	Name string
	Age  int
}
