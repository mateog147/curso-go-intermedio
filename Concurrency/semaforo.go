package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int, 2) //esto define una capacidad de 2

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		c <- 1
		go doWork(i, c, &wg)
	}

	wg.Wait()
	fmt.Println("All goroutines have completed.")
}

func doWork(id int, c chan int, wg *sync.WaitGroup) {
	fmt.Println("Goroutine", id, "is waiting to enter critical section")
	defer wg.Done()
	time.Sleep(time.Millisecond * 2000) // Simulate some initial work

	println("Goroutine", id, "is done work")
	// Critical section ends
	<-c
}
