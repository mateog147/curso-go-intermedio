package main

import "time"

func DoThings(interval time.Duration, c chan int, param int) {
	time.Sleep(interval)
	c <- param
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	duration1 := 4 * time.Second
	duration2 := 2 * time.Second

	go DoThings(duration1, c1, 1)
	go DoThings(duration2, c2, 2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			println("Received from c1:", msg1)
		case msg2 := <-c2:
			println("Received from c2:", msg2)
		}
	}
}
