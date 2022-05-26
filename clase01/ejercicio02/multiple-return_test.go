package ejercicio02

import "testing"

type mystruct struct {
	num   float64
	floor int
	ceil  int
}

func TestRound(t *testing.T) {
	testcases := []mystruct{
		{3.14, 3, 4},
		{5.78, 5, 6},
		{12.00, 12, 12},
	}
	for _, tc := range testcases {
		floor, ceil := Round(tc.num)
		if floor != tc.floor || ceil != tc.ceil {
			t.Fatalf("Se espera %v, %v, y se obtuvo: %v, %v", tc.floor, tc.ceil, floor, ceil)
		}
	}
}
