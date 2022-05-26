package ejercicio02

import "math"

//Return the floor and ceil functions of a float number
//Example: if 3.14 then return (3, 4)
func Round(num float64) (floor, ceil int) {
	floor = int(math.Floor(num))
	ceil = int(math.Ceil(num))
	return
}
