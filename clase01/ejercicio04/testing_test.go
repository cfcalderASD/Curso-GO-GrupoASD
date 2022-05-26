package ejercicio04

import "testing"

type testcase struct {
	input, output string
}

//Try a test of the function Reverese() from testing.go
//Tip: use t.FatalF() inside an "if" statement if the test fail
//Example: testing Reverse("hola") should return "aloh"
func TestReverse(t *testing.T) {
	testcases := []testcase{
		{"Hello World!", "!dlroW olleH"},
		{"GrupoASD", "DSAopurG"},
		{"ana", "ana"},
	}
	for _, tc := range testcases {
		outGetted := Reverse(tc.input)
		if tc.output != outGetted {
			t.Fatalf("Se esperaba %v, pero se obtuvo %v", tc.output, outGetted)
		}
	}
}
