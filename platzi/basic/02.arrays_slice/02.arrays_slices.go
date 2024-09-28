package main

import "fmt"

func main() {
	/*
		TODO ARRAY
		Los array son parecidos a los slices solo que los arrays son
		inmutables y como se creen así se deben de quedar, ya que el tamaño
		y el tipo no se pueden cambiar EJ:

		var <name> = [<size>]<type>{<value>,<etc>...}
	*/
	/*
		TODO LEN, CAP
		Son dos funciones para el manejo de slice y arrays
		* len(<array|slice>) cuenta la cantidad de items en un array|slice
		* cap(<array|slice>) cuenta la capacidad del array|slice
	*/

	var array = [4]int{1, 2, 3, 4}
	fmt.Println("array ", array[1])
	array[3] = 34
	fmt.Println(array[3])
	fmt.Println("array items:", len(array), "capacity:", cap(array))

	/*
		TODO SLICE
		Los slice funcionan como un array pero estos si son mutables,
		el tipo no se cambia pero si es mas flexible en el tamaño,
		se declara como un array pero sin darle tamaño EJ:

		var <name> = []<type>{<value>,<etc>...}
	*/
	var slice = []int{1, 2, 3, 4, 5, 6}
	fmt.Println("slice ", slice[1])
	slice[3] = 38
	fmt.Println(array[3])
	fmt.Println("slice: items:", len(slice), "capacity:", cap(slice))

	/*
		TODO METODOS DE IMPRESION
		Para imprimir un rango deseado en los slice|array se usa los dos
		puntos (:) EJ:

		* slice[1:3] imprime desde la posicion 1 hasta la posicion 2, el 3
			no se agrega al rango
		* slice[:4] imprime desde la pisicion 0 hasta la posicion 3
		* slice[2:] imprime desde la posicion 2 en adelante
	*/
	fmt.Println("Impresion: ")
	fmt.Println(array[3:])
	fmt.Println(array[1:4])
	fmt.Println(array[:4])

	/*
		TODO APPEND
		Es un metodo exclusivo del slice que nos ayuda agregar
		mas items al slice.
		EJ:

		* slice = append(slice, 56)

		Si se necesita agregar los elementos de una lista a otra
		se debe de colocar tres puntos(...) en el valor de la lista
		del append EJ:

		* slice = append(slice, newSlice...)

		de esta misma manera podemos hacer un unshif, el cual coloca
		el valor al principio del slice EJ:

		* slice = append([]int{34}, slice...)

	*/
	slice = append(slice, 56)
	fmt.Println("append 56:", slice)

	var newSlice = []int{52, 53, 58}
	fmt.Println("new slice:", newSlice)
	slice = append(slice, newSlice...)
	fmt.Println("merge slices:", slice)

	// unshif
	slice = append([]int{34}, slice...)
	fmt.Println("unshif 34:", slice)
}
