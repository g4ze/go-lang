package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup // usually pointers
var signals = []string{"test"}
var mut sync.Mutex //usually pointer
func main() {

	// go greeter("Nilay")
	// greeter("Gupta")
	websitelist := []string{
		"https://lco.dev",
		"https://github.com",
		"https://go.dev",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}
func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

// go routines are different from a thread. checkout web for more interview stuff.
//check out the difference between concurrency and parallelism as well
func getStatusCode(endpoint string) {
	defer wg.Done()
	res, _ := http.Get(endpoint)
	mut.Lock()
	//mitex locks the memory adress so that only the current intance and no other on gpong goroutine can access it
	signals = append(signals, endpoint)
	mut.Unlock()

	fmt.Printf("%v status code %v\n", res.StatusCode, endpoint)
}
