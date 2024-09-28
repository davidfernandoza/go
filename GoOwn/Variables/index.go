package main

import "fmt"

/*
 TIPOS DE DATOS:

 Primitivos:
 - numeros
 - cadenas
 - boleanos

 Conjuntos o Set:
 - arreglos
 - estructuras

 Referencias:
 - punteros
 - segmentos
 - mapas
 - funciones
 - canales

 Interfaces:
 - interfaz
*/

/*
* Para declarar una variable de Scope general debe se hacerse con la palabra "var"
 */

var general = "Esto es una variable general"

func main() {
	/*
	* Forma larga: var - <name> - <type>
	* Go asigna de forma predeterninada un valor a la variable cuando no se le da un valor al
	* declararla
	 */
	var number int
	number = 25

	/*
	* Forma corta: <name> := <value>
	* Asigna el tipo a la variable automaticamente con el valor iniciandola con ":=" y solo se utiliza
	* en la creacion de la variable
	 */
	fullName := "Juan"
	fullName = "David"

	/*
	* Se pueden asignar variable de forma multiple
	 */
	fullName, number = "Dario", 45
	fmt.Println(number, fullName)

	/*
	* esto se puede usar para el intercambio de valores entre variables sin tener que utilizar
	* una variable temporal
	 */
	fullNameOne := "Maria"
	fullNameTwo := "Carlos"
	fmt.Println(fullNameOne, fullNameTwo)

	fullNameOne, fullNameTwo = fullNameTwo, fullNameOne
	fmt.Println(fullNameOne, fullNameTwo)

	/*
	* Go se auto-protege y no deja compilar si no se esta usando una variable o un paquete que se
	* declaro o importo.
	 */

	/*
	* Se pueden declarar variables de forma masiva con la funcion var()
	 */
	var (
		saludo = "Hola"
		feo    = "HP"
	)
	fmt.Println(saludo, feo)

	/*
	* Al declarar una variable con la palabra "var" no se necesita agregar el tipo y/o caracter ":="
	 */
	var cerveza = "Poker"
	fmt.Println(cerveza)
	fmt.Println(general + " en main")
	generalFunc()
}

func generalFunc() {
	fmt.Println(general + " en generalFunc")

}
