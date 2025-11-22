package main

import "fmt"

type Address struct {
	Street string
	City   string
}

type BaseEmployee struct {
	Name    string
	Age     int
	Address // Composición: Employee "tiene un" Address
}

func (e BaseEmployee) GetCity() string {
	return e.Address.City
}

type FullTimeEmployee struct {
	BaseEmployee // Composición: FullTimeEmployee "tiene un" BaseEmployee
	Salary       float64
}

func main() {
	emp := FullTimeEmployee{
		BaseEmployee: BaseEmployee{
			Name: "Alice",
			Age:  30,
			Address: Address{
				Street: "123 Main St",
				City:   "Metropolis",
			},
		},
		Salary: 60000.0,
	}

	fmt.Printf("Name: %s\n", emp.Name)
	fmt.Printf("Age: %d\n", emp.Age)
	fmt.Printf("Street: %s\n", emp.Address.Street)
	fmt.Printf("City: %s\n", emp.Address.City)
	fmt.Printf("Salary: %.2f\n", emp.Salary)
	fmt.Println(emp.GetCity())
}
