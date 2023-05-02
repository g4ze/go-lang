package main

import "fmt"

func main() {
	fmt.Printf("Variable type: %T\n", 42)
	// Common types in Go; classic declaration;
	var variable1 string = "Hello World"
	var variable2 int = 32
	var variable3 float64 = 3.14
	var variable4 bool = true

	fmt.Printf("Variable1 type: %T\n", variable1)
	fmt.Printf("Variable2 type: %T\n", variable2)
	fmt.Printf("Variable3 type: %T\n", variable3)
	fmt.Printf("Variable4 type: %T\n", variable4)

	// Short declaration; Go will infer the type
	variable5 := 63.2
	variable6 := false
	fmt.Printf("Variable5 type: %T\n", variable5)
	fmt.Printf("Variable6 type: %T\n", variable6)

	// Multiple declaration
	a, b := 10, "20"
	fmt.Printf("Variable7 &8 type: %T & %T\n", a, b)
	var variable9, variable10 int = 10, 20
	fmt.Printf("Variable9 &10 type: %T & %T\n", variable9, variable10)
	//no initialisation
	var variable11 string
	variable11 = "Hello World"
	fmt.Printf("Variable11 type: %T\n", variable11)

}
