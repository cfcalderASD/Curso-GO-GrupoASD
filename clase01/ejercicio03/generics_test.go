package ejercicio03

import "testing"

type IntegerTestCase struct {
	input  int
	output int
}

type FloatTestCase struct {
	input  float64
	output float64
}

func TestOperation(t *testing.T) {
	intTC := []IntegerTestCase{
		{2, 1},
		{5, 2},
		{13, 6},
	}
	floatTC := []FloatTestCase{
		{5, 2.5},
		{6.42, 3.21},
		{2.5, 1.25},
	}
	for _, tc := range intTC {
		outGetted := Operation(tc.input)
		if tc.output != outGetted {
			t.Fatalf("Se esperaba %v, pero se obtuvo %v", tc.output, outGetted)
		}
	}
	for _, tc := range floatTC {
		outGetted := Operation(tc.input)
		if tc.output != outGetted {
			t.Fatalf("Se esperaba %v, pero se obtuvo %v", tc.output, outGetted)
		}
	}
}
