package ejercicio01

import "fmt"

//Return "Hello, 'name'. Welcome!"
//Example: "Hello, GrupoASD. Welcome!""
func Hello(name string) string {
	return fmt.Sprintf("Hello, %v. Welcome!", name)
}
