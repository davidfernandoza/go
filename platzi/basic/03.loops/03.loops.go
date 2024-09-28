package main

import "fmt"

func main() {
	/*
		TODO FOR
		En go solo existe una sola de iterar y es con for,no hay mas, pero con
		este se puede remplazar un while, do while y foreach EJ:

		* 	for i := 0; i < <value>; i++ {...} // for
		* 	for <variable> == true {...} // while
		* 	for true {
				if <validation> {
					break
				}
			} // do while
		* for i, value := range slice {...} // foreach
	*/

	// for i
	for i := 0; i <= 10; i++ {
		fmt.Println("for:", i)
	}

	//while
	var iterator int = 0
	for iterator <= 12 { //hace la operacion siempre y cuando este en true
		fmt.Println("while:", iterator)
		iterator++
	}

	//do while
	iterator = 0
	for true {
		if iterator == 12 {
			break
		}
		fmt.Println("do while:", iterator)
		iterator++
	}

	//foreach
	var slice = []int{1, 2, 3, 4, 5, 6, 7}
	for i, value := range slice {
		fmt.Printf("foreach value: %d, index: %d \n", value, i)
	}

	/*
		Si el index no se va usar se hace lo mismo que se hace en las
		funciones con multiple return al momento de retornar diferentes
		valores, se coloca un guion al piso para que go no pida el uso
		de esta variable
	*/
	for _, value := range slice {
		fmt.Printf("foreach value: %d, without index\n", value)
	}
}
