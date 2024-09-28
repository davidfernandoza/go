package main

import "fmt"

func main() {
	/*
		TODO PUNTEROS
		Los punteros son variables que hacen referencia a la
		memoria de otra varianble, se crea agregandole & a la
		variable ya inicializada EJ:

		*	var <name_var> <type> = <value>
			var <name_pointer> = &<name_var>
		Si se imprime tal cual el puntero muestra la direccion
		de memoria que almacena, si se quiere el valor hay que
		colocarle al puntero un * al principio

		si se quiere modificar el valor de la variable original
		por medio del puntero se coloca el mismo * al principio
		del puntero y se le asigna el nuevo valor

		los valores de una funcion de un struct son punteros
		a la ram del struct instanciado, pero si nosotros queremos
		modificar estos valores en esta funcion no se puede hacer
		directamente porque se modificaria el pocisionamiento
		en ram del valor del struct, por eso se le debe de colocar
		el simbolo de * para acceder a los valores en ram y no a la
		pocision en ram.
	*/

	var number int = 100
	var pointer = &number
	fmt.Println("number", number)
	fmt.Println("pointer", pointer, *pointer)
	*pointer = 12
	fmt.Println("number", number)

	var myPc = pc{brand: "msi", price: 1000, ram: 2}
	fmt.Println(myPc)

	myPc.addRamToPc(34)
	fmt.Println(myPc)
}

// Acceso a los valores del puntero de pc con *
func (myPc *pc) addRamToPc(ram int) {
	myPc.ram = ram
}

type pc struct {
	brand string
	ram   int
	price int
}
