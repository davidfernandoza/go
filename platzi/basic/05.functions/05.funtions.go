package main

import "fmt"

func main() {
	/*
		TODO FUNC
		Es la forma de declarar una funcion en go, si se requiere
		pasarle un parametro se le tiene que pasar la variable
		y el tipo de dato de la variable EJ:

		*	func <nombreFuncion> (<variable> <tipo>){
				fmt.Println(<variable>)
			}
		si la funcion recibe varios parametros con el mismo tipo de
		variable es buena practica simplificar el tipado

		*	func <nombreFuncion> (<variable1>, <variable2 <tipo1>, <variable3> <tipo2>){
				fmt.Println(<variable1>, <variable2>, <variable3>)
			}
		si se quiere retornar un valor se le debe de decir a la
		funcion que tipo de dato va a retornar

		*	func <nombreFuncion> (<variable> <tipo>) <tipo_retorno>{
				retunr <variable>
			}

		se pueden retornar varios valores a la vez de una misma
		funcion

		*	func <nombreFuncion> (<variable> <tipo>) (a, b <tipo_retorno>){
				retunr <variable_para_a>, <variable_para_b>
			}
		al llamar a la funcion se debe de llenar las variables igual como
		se retorna con comas, pero como go obliga a usar todas las variables
		y por ejemplo no se quiere usar una del return, pues se usa ralla al
		piso para que go la descarte

	*/

	//Dato
	saludar("Hola amigos")

	//Multi dato
	llamadoDeNumero(1, 2, "hola")

	//Return
	var numero int = suma(1, 3)
	fmt.Println(numero)

	//Multi return
	var suma, multiplicacion int = sumaMultiplicacion(2, numero)
	fmt.Println(suma, multiplicacion)

	//Descarte de variable "_"
	suma, _ = sumaMultiplicacion(2, numero)
	fmt.Println(suma)

}

func saludar(saludo string) {
	fmt.Println(saludo)
}

func llamadoDeNumero(numero1, numero2 int, cadena string) {
	fmt.Println(cadena, numero1, numero2)
}

func suma(numero1, numero2 int) int {
	return numero1 + numero2
}

func sumaMultiplicacion(numero1, numero2 int) (suma, multiplicacion int) {
	return numero1 + numero2, numero2 * numero1
}
