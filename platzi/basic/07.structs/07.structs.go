package main

import "fmt"

/*
	TODO STRUCTS
	En go no hay una forma de crear clases como se hacen en
	otros lenguajes, sino que existe los "structs" que como
	su nombre lo dice, son estructuras de datos.
	para declarar se hace a fuera de las funciones globales.
	EJ:

	*	type <name> struct {code...}

	las variables del struct se declara nombre y tipo
	EJ:

	*	type <name> struct{
			<name_var> <type_var>
		}
	No hay una forma concreta de hacer un struct con valores por
	defecto, se puede hacer una funcion que cree el struct
	y le asigne un valor por defecto.

	El struct hay que verlo como un objeto de JS

	Para agregar una funcion en un struct se debe de inicializar fuera
	de el y pasarle por parentecis de que struct es EJ:

	*	func (<var_struct><struct_name>) <function_name> (<param> <type_param>){
			code ...
		}
	y se llama como instancia del struct:
	<struct_name>.<funtion_name>(<params>)

	Los struct son unicos en cada proyecto

*/

type Car struct {
	brant string
	year  int
}

func (MyCar Car) startCar(key string) {
	fmt.Println(MyCar.brant, key)
}

func main() {
	fmt.Println(NewCar())

	// Instanciacion de struct
	var mayCar = Car{brant: "Honda", year: 2013}
	fmt.Println(mayCar)

	// Llamaado de la funcion start car
	mayCar.startCar("Run")
}

func NewCar() Car {
	// Otra forma de instanciacion de struct
	var car Car
	car.brant = "Audi"
	car.year = 293
	return car
}
