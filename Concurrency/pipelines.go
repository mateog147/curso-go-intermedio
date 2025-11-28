package main

// canal de escritura
func Generator(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

// canal de lectura y otro de escritura
func Double(in <-chan int, out chan<- int) {
	for val := range in {
		out <- val * 2
	}
	close(out)
}

func Printer(c chan int) {
	for val := range c {
		println(val)
	}
}

func main() {
	genChan := make(chan int)
	doubleChan := make(chan int)

	go Generator(genChan)
	go Double(genChan, doubleChan)
	Printer(doubleChan)
}
