package main

import "fmt"

// existen canales sin buffer y con buffer
// canales sin buffer: la comunicacion entre gorutinas es sincronica
// canales con buffer: la comunicacion entre gorutinas es asincronica
func main() {
	c := make(chan int, 3) // canal con buffer de tamaño 3

	c <- 42
	c <- 27
	c <- 13

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
