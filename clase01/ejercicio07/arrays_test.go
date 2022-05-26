package ejercicio07

import "testing"

type testcase struct {
	input  [5]int
	output [5]int
}

func TestSquare(t *testing.T) {
	tc := testcase{[5]int{1, 2, 3, 4, 5}, [5]int{1, 4, 9, 16, 25}}
	msg := square(tc.input)
	for i, _ := range tc.input {
		if tc.output[i] != msg[i] {
			t.Fatalf("Se esperaba %v, pero se obtuvo %v", tc.output[i], msg[i])
		}
	}

}
