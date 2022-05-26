package ejercicio05

import "testing"

type testcase struct {
	lowerBound, upperBound, output int
}

func TestSumBetween(t *testing.T) {
	testcases := []testcase{
		{1, 1, 1},
		{3, 7, 25},
		{5, 13, 81},
	}
	for _, tc := range testcases {
		msg := SumBetween(tc.lowerBound, tc.upperBound)
		if tc.output != msg {
			t.Fatalf("Se esperaba %v, pero se obtuvo %v", tc.output, msg)
		}
	}
}
