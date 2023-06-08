package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("welcome to channels")
	//channel chan with the passing type of int
	myCh := make(chan int)
	//arrow always point towarsds left
	// //initialising channel with 5
	// myCh <- 5
	// fmt.Println(<-myCh) this would would throw error cause mych cant be made if noone is lostening to it
	// kind of a recursive error
	wg := &sync.WaitGroup{}
	wg.Add(2)
	defer wg.Wait()
	go func(ch chan int) {
		fmt.Println(<-myCh)
		fmt.Println(<-myCh)
		//a way to check if a channe is open or closed
		val, isopen := <-myCh

		fmt.Println(isopen, val)
		wg.Done()
	}(myCh)
	go func(ch chan int) {
		myCh <- 5
		myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh)
}
