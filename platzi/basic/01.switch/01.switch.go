package main

import "fmt"

func main() {
	/*
		TODO SWITCH
		Es una estructura de control que funciona como if, else if y else en
		un solo enunciado, la diferencia con otros lenguajes es que a cada
		caso no hay que hacerle un break para que pare la ejecución, sino
		que se hace es un fallthrough esto le da el poder al caso siguiente de
		ejecutarse, EJ:

		* switch condition {
			case <value> :
				code...
				fallthrough // se ejecuta tambien el caso de valor_2
			case <value_2> :
				code...
			default:
				code...
		}
	*/
	var expr string = "hola"
	switch expr {
	case "hola":
		fmt.Println("hola")
		fallthrough
	case "chao":
		fmt.Println("chao")
	default:
		fmt.Println("nada")
	}

	/*
		tambien puede haber switch sin condicion y se condiciona en
		cada case
	*/

	switch {
	case expr == "hola":
		fmt.Println("hola")
		fallthrough
	case expr == "chao":
		fmt.Println("chao")
		fallthrough
	default:
		fmt.Println("nos vemos")

	}
}
