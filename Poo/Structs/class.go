package main

//en POO el pilar son las clases, en go podemos usar type y structs para definir tipos de datos complejos

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) SetID(id int) {
	e.id = id
}

func (e *Employee) GetID() int {
	return e.id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetName() string {
	return e.name
}

func (e *Employee) String() string {
	return fmt.Sprintf("Employee[ID=%d, Name=%s]", e.id, e.name)
}

func main() {
	e := Employee{}
	fmt.Printf("%v", e)
	e.id = 1
	e.name = "Name"
	fmt.Printf("%v", e)
	fmt.Println(e)
	e.SetID(2)
	e.SetName("Otro Name")
	fmt.Printf("\nID: %d, Name: %s \n", e.GetID(), e.GetName())
}
