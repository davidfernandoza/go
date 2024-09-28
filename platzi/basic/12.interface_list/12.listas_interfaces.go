package main

import "fmt"

func main() {
	/*
			TODO LISTAS DE INTERFACES
			En lenguajes como php se puede crear arrays o slice
			con valores de diferentes tipos como por ejemplo
			[12, "hola", true] pero en go no se permite hacer esto
			por que a la hora de crear un slice se le agrega el tipo del
			contenido, para este problema se utiliza las listas
			de interfaces la creacion se hace como un slice normal
			pero con cierta variacion EJ:

		*	var <name_slice> = []interface{}{<value1>, <value2>}

		solo se agrega unas llaves mas que quiere decir que cada valor
		puede ser diferente
	*/

	var mySlice = []interface{}{"hola", 12, true, 3.24}
	for _, value := range mySlice {
		fmt.Println(value)
	}
}
