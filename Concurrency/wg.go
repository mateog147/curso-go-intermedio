package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*asi corre todo en paralelo
	for i := 0; i < 10; i++ {
		go doSomething(i, nil)
	}
	// simple wait to allow goroutines to finish
	time.Sleep(4 * time.Second)*/

	var wg sync.WaitGroup
	//asi corre todo en paralelo pero espera a que terminen
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			doSomething(i, &wg)
		}()
	}
	wg.Wait()
	fmt.Println("All done!")
}

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Doing something...", i)
	time.Sleep(3 * time.Second)
	fmt.Println("Done!", i)
}
