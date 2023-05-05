package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// rpogram in golang exploring time package; prtty funny lol

func main() {
	fmt.Println("We manipulate time")
	currentTime := time.Now()
	fmt.Println("Current time is: ", currentTime)

	// trying to print more clean format:
	fmt.Println("\nToday is: " + currentTime.Format("01-02-2006 Monday 15:04:05"))

	// to save a date/time:
	date := time.Date(2022, time.July, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println("\nStored Date is:", date)
	fmt.Println("\nsorted strored date is:", date.Format("01-02-2006"))

	//telling the type if date variable
	fmt.Println("\nType of Date is ", reflect.TypeOf(date))

	//inputting a time form user
	fmt.Println("\nEnter a year")

	read := bufio.NewReader(os.Stdin)
	inputYear, _ := read.ReadString('\n')

	// converting the input String to int value
	conv, _ := strconv.Atoi(strings.TrimSpace(inputYear))

	// initializing a new date variable if type time.Time
	newDate := time.Date(conv, time.July, 12, 23, 23, 0, 0, time.UTC)

	fmt.Println("\nNew Date is: ", newDate)
}
