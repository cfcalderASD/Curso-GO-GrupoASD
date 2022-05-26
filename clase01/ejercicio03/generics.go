package ejercicio03

//Return the number divided by two: number/2
//If the number is an integer, return the integer division
//Example: if the number is 5, then return 2.
//Example: if the number is 5.0 (float64), then return 2.5
func Operation[Number int | float64](number Number) Number {
	return number / 2
}
