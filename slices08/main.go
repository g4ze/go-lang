package main

import "fmt"

//this is dorectly 8 skipping 7 cz i have included array part in here as well
// make sure to check out an article about deifference bw array and slices on memory level
func main() {
	fmt.Println("Slices are just normal arrays with a layer of abstraction in the go lang")
	slicetype1 := make([]int, 4)
	slicetype1[0] = 679
	slicetype1[1] = 567
	slicetype1[2] = 987
	slicetype1[3] = 345
	fmt.Printf("Return type of amend func is: %T\n", append(slicetype1, 55))

	var list [2]int // an array
	list[0] = 1
	list[1] = 2
	fmt.Println(list)
	//if we would have given it some size value the it would have been an array
	//since the soze wasnt specofoed while declaration and the subsets were empty, it was read as a slice.
	var list2 = []string{}
	list2 = append(list2, "34")
	fmt.Println(list2)
}
