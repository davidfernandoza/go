package main

import "fmt"

type computer struct {
	brand string
	ram   int
	price int
}

/*
	TODO STRINGERS
	Es una funcion de un struct que se debe de declarar como
	"String", esta funcion cuando se declara cambia la forma
	de como se imprime el struct cuando se hace un
	fmt.Println(<sturct>), esto imprime lo que retorna la funcion
	pero antes de esto se debe de llamar para cambiar esta estructura
*/

func (myComputer computer) String() string {
	return fmt.Sprintf("Mi computador es %s tiene %d de ram y cuesta %d dolares", myComputer.brand, myComputer.ram,
		myComputer.price)
}

func main() {
	var myPc = computer{brand: "HP", ram: 16, price: 3200}
	fmt.Println(myPc.String())
}
