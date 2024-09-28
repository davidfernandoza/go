package main

import "fmt"

/*
	TODO INTERFACE
	Una interface facilita el intercambio de metodos que
	puede tener diferentes structs
*/

// Interface
type figura2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (myCuadrado cuadrado) area() float64 {
	return myCuadrado.base * myCuadrado.base
}

func (myRectangulo rectangulo) area() float64 {
	return myRectangulo.altura * myRectangulo.base
}

/*
Como se puede notar el struct de cuadrado y de rectangulo
comparten el mismo metodo, entonces para no volver repetitivo
el uso de este lo que se hace es que se crea la interface que
tiene este metodo en ella, y se usa en una funcion que recibe
por parametro una instancia que debe de cumplir con la
interface osea tener este metodo de area(), y esta funcion
ya puede ejecutar area() independientemente de que instancia sea
si cuadrado o rectangulo.
*/
func calcularArea(instancia figura2D) {
	fmt.Println(instancia.area())
}

func main() {
	var myCuadrado = cuadrado{base: 23}
	var myRectangulo = rectangulo{base: 12, altura: 35}
	calcularArea(myRectangulo)
	calcularArea(myCuadrado)
}
