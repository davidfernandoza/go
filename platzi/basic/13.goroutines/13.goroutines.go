package main

import (
	"fmt"
	"sync"
	"time"
)

/*
TODO GORUTINES (concurrencia)
Es la forma de go para hacer asyncrono una ejecucion  de una
tarea en go, esto funciona muy parecido a javascript y
su event loop, la ventaja de las gorutines al paralelismo
es que no esperan ejecuciones sino que va ejecutando al tiempo
todas las tareas.

como go se ejecuta siempre syncronicamente, para ejecutar
una gorutine se debe de hacer unos arreglos antes a la
funcion que la ejecutara.

para que un proceso sea ejecutado en una gorutine se usa
esta sintaxis EJ:

go <proceso>
*/

func main() {

	/*
		Si esto se deja así no pasara nada en consola,
		ya que el hilo conductor de main se ejecuta primero
		que la gorutine, y esta gorutine se ejecuta en otro
		hilo, el cual el hilo principal no espera.
	*/

	go printVar("hola")

	/*
		para evitar esto se puede colocar un setTimeout (time.Sleep)
		el cual no es recomendable, ya que se perderia
		el sentido de la concurrencia.
	*/

	time.Sleep(time.Second * 1)

	/*
		Una mejor forma es con promesas (WaitGroup)
		el cual es un recopilador de gorutines que las ejecuta
		y las libera poco a poco de forma concurrente.

		TODO USO
		1. 	se crea con el paquete sync un WaitGroup
		2. 	se le dice cuantas gorutines va a ejecutar el waitgroup
			con la funcion Add()
		3.	le pasamos por parametro el puntero del waitgroup a la
			funcion	que va estar ejecutando el gorutine, esta funcion
			debe de recibir el valor del puntero con *sync.WaitGroup
		4. 	se ejecuta la funcion Done del puntero del waitgroup
			ejecutandose en un difer para avisarle que ya esta listo
		5. 	se usa la funcion Wait del waitgroup original para
			que el hilo principal espere
	*/

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go printVarWithWaitGroup("amigo", &waitGroup)

	/*
		Esto tambien se ve mucho hacerse con funciones
		anonimas

		las funciones anonimas se crean y ejecutan al
		mismo tiempo

		su sintaxis es EJ:

		* func (<var><var_type>){code..}(<var>)
	*/
	go func(text string, waitGroup *sync.WaitGroup) {
		fmt.Println(text)
		defer waitGroup.Done()
	}("anonimo", &waitGroup)

	waitGroup.Wait()
}

func printVar(variable string) {
	fmt.Println(variable)
}

func printVarWithWaitGroup(variable string, waitGoup *sync.WaitGroup) {
	fmt.Println(variable)
	defer waitGoup.Done()
}
