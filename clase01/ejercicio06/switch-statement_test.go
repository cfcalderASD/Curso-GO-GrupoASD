package ejercicio06

import "testing"

type testcase struct {
	input, output string
}

func TestSwitchGame(t *testing.T) {
	testcases := []testcase{
		{"amarillo", "Ganaste"},
		{"azul", "Tienes otro intento"},
		{"rojo", "Perdiste"},
		{"negro", "Sin trampas"},
	}
	for _, tc := range testcases {
		msg := switchGame(tc.input)
		if tc.output != msg {
			t.Fatalf("Se esperaba %v, pero se obtuvo %v", tc.output, msg)
		}
	}
}
