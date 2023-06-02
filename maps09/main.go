package main

import "fmt"

func main() {
	map1 := make(map[string]int)
	map1["one"] = 1
	map1["two"] = 2
	map1["three"] = 3

	fmt.Println(map1["three"])

	for keyss, valuesss := range map1 {
		fmt.Println(keyss, ":", valuesss)
	}
}
