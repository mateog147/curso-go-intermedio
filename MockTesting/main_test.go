package main

import "testing"

func TestGetFulltimeEmployeeById(t *testing.T) {
	// Guardar las funciones originales
	originalGetPersonByDNI := GetPersonByDNI
	originalGetEmployeeById := GetEmployeeById

	// Restaurar las funciones originales al final del test
	defer func() {
		GetPersonByDNI = originalGetPersonByDNI
		GetEmployeeById = originalGetEmployeeById
	}()

	// Mock de GetPersonByDNI
	GetPersonByDNI = func(dni string) (Person, error) {
		return Person{
			DNI:  dni,
			Name: "John Doe",
			Age:  30,
		}, nil
	}

	// Mock de GetEmployeeById
	GetEmployeeById = func(id int) (Employee, error) {
		return Employee{
			Position: "Developer",
			Id:       id,
		}, nil
	}

	fulltimeEmployee, err := GetFulltimeEmployeeById(1, "12345678A")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if fulltimeEmployee.Person.DNI != "12345678A" {
		t.Errorf("Expected DNI '12345678A', got %s", fulltimeEmployee.Person.DNI)
	}

	if fulltimeEmployee.Person.Name != "John Doe" {
		t.Errorf("Expected Name 'John Doe', got %s", fulltimeEmployee.Person.Name)
	}

	if fulltimeEmployee.Employee.Position != "Developer" {
		t.Errorf("Expected Position 'Developer', got %s", fulltimeEmployee.Employee.Position)
	}
}
