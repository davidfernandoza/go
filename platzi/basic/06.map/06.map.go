package main

import "fmt"

func main() {
	/*
		TODO MAP
		Es una forma de tener una estructura de datos de tipo llave valor
		si se necesita almacenar algun dato con una llave de busqueda
		se puede utilizar un map, el unico problema del map como es una
		estructura concurrente no guarda el orden si se hace un foreach,
		si se necesita algun orden se recomienda utilizar un slice.
		EJ:

		* var <nombre_variable> = map[<tipe_key>]<tipe_value>{<key>:<value>}

		los maps  cuando se imprimen tampoco muestran sus comas, sino
		un espacio.

		estos son mas eficientes que los slice ya que implementan concurrencia
	*/

	var nameAndAge = map[string]int{
		"carlos": 21,
		"mario":  31,
	}
	fmt.Println(nameAndAge)
	fmt.Println(nameAndAge["carlos"])

	/*
		Tambien se pueden crear map con make, si queremos que este vacio
	*/

	var mapWithMake = make(map[string]string)
	mapWithMake["persona_1"] = "Jose"
	fmt.Println(mapWithMake)
	fmt.Println(mapWithMake["persona_1"])

	// iterar maps
	for i, value := range nameAndAge {
		fmt.Println(i, value)
	}

	/*
		Si se accede a una llave que no fue previamente
		guardada en el map, go te responde con un zero value
		de el tipo del valor.

		para saber si el map cuenta con ese valor se debe de
		inyectar a una variable con doble asignacion
		la segunda asignacion dira si ese valor existe en el
		map EJ:

		*	var <name>,<validation> = <name_map>[<key_map>]
	*/
	fmt.Println(mapWithMake["persona_2"]) // "" vacio
	fmt.Println(nameAndAge["marta"])      // 0 cero

	var nombre, ok = mapWithMake["persona_2"]
	fmt.Println(nombre, ok) // false

	nombre, ok = mapWithMake["persona_1"]
	fmt.Println(nombre, ok) //true
}
