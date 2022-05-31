package saludos

import "fmt"

// Saludo return un saludo a una persona.

func Saludo(name string) string {
	// Retorna un saluda esto el nombre
	message := fmt.Sprintf("Hola, %v. Welcome!", name)
	return message
}
