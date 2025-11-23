package main

import (
	"fmt"

	"github.com/donvito/hellomod"
	hellomod2 "github.com/donvito/hellomod/v2"
	utils "github.com/mateog147/helloplatzi"
)

func main() {
	hellomod.SayHello()
	hellomod2.SayHello("Belli ")
	saludo := utils.HolaMundo()
	fmt.Println(saludo)
}
