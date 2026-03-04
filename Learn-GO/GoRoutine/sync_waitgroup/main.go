package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	// defer wg.Done()// signal that go routin is done
	fmt.Printf("Worker %d started\n", i)
	//task 
	fmt.Printf("Worker %d end\n", i) 
	wg.Done()
}

func main() {
	fmt.Println("Sync_WaitGroup")
	
	var wg sync.WaitGroup
	//Worker 1
	for i:=1; i<=3; i++ {
		wg.Add(1) // increment the wait group counter
		worker(i, &wg);
	}
	
	wg.Wait() //wait for all functions to complete 
	fmt.Println("Worker task completed!")
}