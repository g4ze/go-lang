package main

import (
	"fmt"
	"sync"
)

//we will learn go race condition

func main() {
	fmt.Println("Race condition")
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	var score = []int{0}
	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Print("One R\n")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		fmt.Print("Two R\n")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Print("Three R\n")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}

//we use mutex to avoid the race condition of the memory
