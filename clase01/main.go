package main

//This is the main package
//func main() is always executed when you run your module
//Open a new terminal, step at the clase01 directory,
//type "go run ." and see what happen

import (
	"fmt"
)

//If a method from another package start with camel case then
//you can use it here. For example when you finish ejercicio01
//try this below the welcome line:
// fmt.Println(ejercicio01.Hello("World"))
//then, type in the console "go run ."
func main() {
	fmt.Println("Welcome! Ready to learn go?")
}
