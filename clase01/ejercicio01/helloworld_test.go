package ejercicio01

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	name := "GrupoASD"
	want := "Hello, GrupoASD. Welcome!"
	message := Hello(name)
	if message != want {
		t.Fatalf(`"%v" es diferente a "%v"`, message, want)
	}
}

func FuzzHello(f *testing.F) {
	testcases := []string{
		"GrupoASD", "Alice", "Bob",
	}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, input string) {
		want := fmt.Sprintf("Hello, %v. Welcome!", input)
		message := Hello(input)
		if message != want {
			t.Fatalf(`"%v" es diferente a "%v"`, message, want)
		}
	})
}
