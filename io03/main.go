package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the beginning")
	oss := os.Stdin
	reader := bufio.NewReader(oss)
	fmt.Printf("%T\n", reader)
	fmt.Println("Enter your name:")
	// ReadString will block until the delimiter is entered
	//comma ok
	name, _ := bufio.NewReader(oss).ReadString('\n')
	fmt.Println("Hello, " + name)
}
