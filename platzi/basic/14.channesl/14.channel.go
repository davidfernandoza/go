package main

import "fmt"

func main() {
	/*
		TODO CHANNELS
		Los channels no son tan optimos como lo son los
		waitgroups, pero si el sistema no requiere esto,
		estos channels para ejecutar gorutines son una
		mejor opcion ya que son mas faciles de mantener.

		1 	el canal se crea con la palabra reservada make
			y el tipo de dato "chan" y el tipo de dato que va
			almacenar, es buena practica pasarle como segundo parametro
			la cantidad de gorutines que debe de manejar
			EJ:

			*	var <channel_var> = make(chan <type_var>, <num_gorutines>)

		2	se crea una funcion la cual es la que manejara
			el canal, esta funcion inyecta el valor al
			canal, la inyeccion se hace con el simbolo "<-"

			*	go func(<var><type_var>, <channel_var> chan<- <type_channel>){
					<channel_var> <- <var>
				}

		3	para usarse solo le pedimos la salida a la variable
			del canal yan no con el simbolo despues sino antes
			*	<- <var_channel>

	*/

	var channel = make(chan string, 1)
	go func(text string, chennel chan<- string) {
		chennel <- text
	}("hola", channel)
	fmt.Println(<-channel)

	/*
		Para el manejo de los canales podemos usar
		palabras reservadas del lenguaje como lo son
		"close", "range", "len", "cap"

		*	close(<name_channel>)
		permite cerrar el canal, esto es una buena
		practica ya que go no estaria esperando mas
		valores para el canal

		*	len, cap y range funcionan igual como se vio
			con los arrays, dan la cantida de valore que tiene el
			channel, la capacidad del channel y el range se usaria
			en una iteracion de valores del channel respectivamente

		NOTE: siempre que se vaya a iterar con range se debe de cerrar
		el chanel.
	*/

	var myChannel = make(chan string, 2)
	myChannel <- "hola"
	myChannel <- "amigo"

	fmt.Println(len(myChannel), cap(myChannel))
	close(myChannel)
	/*
		*	myChannel <- "error"

			si se ejecuta esta linea
			despues del close, aparece un error en go, diciendo
			de que el canal ya esta cerrado
	*/

	for message := range myChannel {
		fmt.Println(message)
	}

	/*
		Cuando se esta manejando muchos canales a la vez
		hay veces no se tiene certesa de que canal respondera
		primero ya que como son concurrentes esto puede pasar
		por eso se usa "select"

		este select se maneja dentro de un ciclo for i
		ya que se debe iterar la cantidad de canales y/o valores
		que se estan agregando a estos mismos

		lo que hace el select es que en caso de que sea haya
		resuelto un canal me ejecuta una porcion de codigo u
		otra.
	*/

	var email1Channel = make(chan string)
	var email2Channel = make(chan string)

	go message("fernando.zapata.live@gmail.com", email1Channel)
	go message("fernando-zapata@live.com", email2Channel)

	for i := 0; i < 2; i++ {
		select {
		case email1 := <-email1Channel:
			fmt.Println(email1)

		case email2 := <-email2Channel:
			fmt.Println(email2)

		}
	}
}
func message(text string, channel chan<- string) {
	channel <- text
}
