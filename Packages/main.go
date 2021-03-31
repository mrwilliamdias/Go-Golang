package main

import (
	"fmt"
	"modulo/Auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing on main archive")
	Auxiliar.Escrever()

	erro := checkmail.ValidateFormat("116")
	fmt.Println(erro)
}
