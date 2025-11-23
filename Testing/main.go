package main

func Suma(a int, b int) int {
	return a + b
}

func GetMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}
