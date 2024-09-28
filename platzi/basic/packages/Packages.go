package Packages

import "fmt"

/*
	TODO PUBLIC
	Si se quiere que un struct sea publico se debe de
	llamar con mayuscula la primera letra igual sus
	atributos,

	por buneas practicas hay que dejar un comentario
	con el nombre del struct y diciendo que es publico

	este struct se llama en el archivo 08.modificadores_acceso
	esto funciona igual con las funciones
*/

// Car is public
type Car struct {
	Brand string
	Year  int
}

// Saludo is public
func Saludo(name string) {
	fmt.Printf("Hola %s\n", name)
}
