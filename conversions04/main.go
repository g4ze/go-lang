package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to conversions in go")
	fmt.Println("Enter your rate: ")
	reader := bufio.NewReader(os.Stdin)
	userrate, _ := reader.ReadString('\n')
	//to add a number to userrate we will need to perfrom a conversion from string to float64;
	//we will use the strconv package
	userrate = strings.TrimSpace(userrate) //trimming extra space of \n in the String
	newuserate, error := strconv.ParseFloat(userrate, 64)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("Your rate+1 is: ", newuserate+1)
	}
}
