package main

import "fmt"

func main() {
	/*
		TODO FMT
		Es un paquete que nos ayua a imprimir en consola
	*/
	var hello string = "Hola"
	var word string = "Mundo"
	var number int = 300

	/*
		TODO PRINTLN
		Imprime con un salto de linea al final
	*/

	fmt.Println(hello, word)

	/*
		TODO PRINTF
		Permite escapar variables dentro del mensaje de salida
		%s: string
		%d: integer
		%v: cualquier dato
		%T: tipo de dato
	*/

	fmt.Printf("el saludo %s %v es bonito para mas de %d personas \n", hello, word, number)

	/*
		TODO TIPO DE DATO
		Si se necesita saber el tipo de dato de una variable que se esta
		imprimiendo se usa el caracter %T
	*/
	fmt.Printf("Tipo de dato es: %T \n", number)

	/*
		TODO SPRINTF
		Permite crear un string pero no lo imprime en consola, solo
		retorna ya el formato listo como un string para luego usar
	*/

	var saludo string = fmt.Sprintf("este es el segundo saludo %s %v y es bonito para mas de %d personas \n", hello,
		word, number)
	fmt.Printf(saludo)

	/*
		TODO SCANLN
		Permite recibir valores que el usuario ingresa por consola
	*/
	fmt.Println("Escribe un numero:")
	fmt.Scanln(&number)
	fmt.Printf("\nEl numero es %d\n", number)
}
