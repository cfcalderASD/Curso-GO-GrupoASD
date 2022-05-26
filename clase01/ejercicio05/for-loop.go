package ejercicio05

//Return the sum of the numbers between lowerbound and upperbound [lb, ub]
//Example: SumBetween(3, 5) should return 12
func SumBetween(lowerBound, upperBound int) int {
	sum := 0
	for ; lowerBound <= upperBound; lowerBound++ {
		sum += lowerBound
	}
	return sum
}
