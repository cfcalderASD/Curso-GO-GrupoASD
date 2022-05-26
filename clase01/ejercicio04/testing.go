package ejercicio04

//Reverse function returns the reverse string.
//Example: Reverse("hola") returns "aloh"
func Reverse(msg string) string {
	runes := []rune(msg)
	n := len(runes)
	for i, j := 0, n-1; i < n/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
