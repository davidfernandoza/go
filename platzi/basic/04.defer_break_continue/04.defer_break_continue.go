package main

import "fmt"

func main() {
	/*
		TODO DEFER
		esta palabra reservada del lenguaje permite ejecutar codigo
		siempre antes de la muerte de una funcion EJ:

		func saludar {
			defer fmt.Println("Hola")
			fmt.Println("Mundo")	defer fmt.Println("Hola")
		}

		se ejecuta primero "Mundo" y por ultimo el "Hola"
		lo unico que si no hace es ejecutarse despues del
		return
	*/
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	/*
		TODO CONTINUE
		La instrucción continue se usa cuando se busca omitir la
		parte restante del bucle, esto seguira con el otro valor
		a iterar EJ:

		* 	for i := 0; i < 3; i++ {
				if i == 2 {
					continue
				}
				fmt.Println(i)
		  	}
	*/

	for i := 0; i < 3; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

	/*
		TODO BREAK
		La instruccion break detiene la ejecucion del ciclo y
		este no seguira ejecutando EJ:

		* 	for i := 0; i < 3; i++ {
				fmt.Println(i)
				if i == 2 {
					break
				}
			  }
	*/
	for i := 0; i < 4; i++ {
		fmt.Println(i)
		if i == 2 {
			break
		}
	}
}
