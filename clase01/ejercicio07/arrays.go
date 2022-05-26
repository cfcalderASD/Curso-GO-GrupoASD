package ejercicio07

//Return an array with squared numbers of the input
//Example Square({1, 2, 3, 4, 5}) should return
//{1, 4, 9, 16, 25}
func square(numbers [5]int) [5]int {
	var squares [5]int
	for i, _ := range numbers {
		squares[i] = numbers[i] * numbers[i]
	}
	return squares
}
