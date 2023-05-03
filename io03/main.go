package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the beginning")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")
	// ReadString will block until the delimiter is entered
	//comma ok
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello, ", name)
}
