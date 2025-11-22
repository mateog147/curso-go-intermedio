package main

import "fmt"

type Employee struct {
	Name string
	Age  int
}

func (e Employee) GetName() string {
	return e.Name
}

func (e Employee) GetAge() int {
	return e.Age
}

type Person interface {
	GetName() string
	GetAge() int
}

func GetNameAndAge(p Person) (string, int) {
	return p.GetName(), p.GetAge()
}

func main() {

	emp := Employee{
		Name: "Clark Kent",
		Age:  28,
	}

	fmt.Printf("Employee Name: %s\n", emp.Name)
	fmt.Printf("Employee Age: %d\n", emp.Age)

	name, age := GetNameAndAge(emp)

	fmt.Printf("Name from interface: %s, Age from interface: %d\n", name, age)

}
