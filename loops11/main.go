package main

import "fmt"

func main() {
	fmt.Println("beginning of loops")
	slice := []string{"Nilay", "Shubh", "CAtto", "Java"}
	//type 1
	for i := 0; i < 5; i++ {
		fmt.Println("i is ", i)
	}
	//type2
	for i := range slice {
		fmt.Print(slice[i], "\n")
	}
	//variation of type2
	for index, name := range slice {
		fmt.Println(index, name)
	}
	//variaion 2 type 2
	for _, name := range slice {
		fmt.Println(name)
	}

	value := 1
	//usage of labels and jump statements, break also works bur not used here
	for ; value < 5; value++ {
		if value == 3 {
			continue
		} else if value == 1 {
			fmt.Println("lol nooo")
		} else if value == 2 {
			goto label1
		}
	}
label2:
	fmt.Println("(helu label 2)")
	goto label3
label1:
	fmt.Print("Thrown to label 1\n")
	goto label2
label3:
	fmt.Println("END")

}
