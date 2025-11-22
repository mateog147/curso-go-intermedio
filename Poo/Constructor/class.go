package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func main() {
	p1 := Person{}
	fmt.Printf("Name p1: %s, Age p1: %d\n", p1.Name, p1.Age)

	p2 := Person{Name: "John", Age: 25}
	fmt.Printf("Name p2: %s, Age p2: %d\n", p2.Name, p2.Age)

	p3 := new(Person) //practicaly same as &Person{}
	p3.Name = "Doe"
	p3.Age = 40
	fmt.Printf("Name p3: %s, Age p3: %d\n", p3.Name, p3.Age)

	p := NewPerson("Alice", 30)
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}
