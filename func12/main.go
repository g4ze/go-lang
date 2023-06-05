package main

import "fmt"

func main() {
	i, _ := input()
	display(i)
}
func display(values []string) {
	for _, value := range values {
		fmt.Print(value, "\n")
	}
}
func input() ([]string, string) {
	values := make([]string, 0)
	values = append(values, "Nilay", "shubh", "catto")
	return values, "returned values"
}
