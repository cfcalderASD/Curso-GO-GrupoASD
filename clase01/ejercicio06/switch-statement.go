package ejercicio06

//Return different messages based on the input
//if ball is "amarillo" return "Ganaste"
//if ball is "azul" return "Tienes otro intento"
//if ball is "rojo" return "Perdiste"
//if ball is another string then return "Sin trampas"
func switchGame(ball string) string {
	msg := ""
	switch ball {
	case "amarillo":
		msg = "Ganaste"
	case "azul":
		msg = "Tienes otro intento"
	case "rojo":
		msg = "Perdiste"
	default:
		msg = "Sin trampas"
	}
	return msg
}
