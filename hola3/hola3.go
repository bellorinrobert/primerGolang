package main

import (
	"fmt"

	"ejemplo.com/saludos"
)

func main() {
	//Obtenemos un saludos y lo imprimimos
	mesange := saludos.Saludo("Doris")
	fmt.Println(mesange)
}