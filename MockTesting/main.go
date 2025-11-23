package main

import "time"

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	Position string
	Id       int
}

type FullTimeEmployee struct {
	Person
	Employee
	Salary float64
}

var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	//SELECT * FROM PERSON WHERE DNI = dni
	var person Person
	return person, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	//SELECT * FROM EMPLOYEE WHERE ID = id
	var employee Employee
	return employee, nil
}

func GetFulltimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var fulltimeEmployee FullTimeEmployee
	employee, error := GetEmployeeById(id)
	if error != nil {
		return fulltimeEmployee, error
	}

	fulltimeEmployee.Employee = employee

	person, error := GetPersonByDNI(dni)
	if error != nil {
		return fulltimeEmployee, error
	}

	fulltimeEmployee.Person = person

	return fulltimeEmployee, nil

}
