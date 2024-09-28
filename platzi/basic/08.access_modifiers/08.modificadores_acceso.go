package main

/*
	TODO MODIFICADORES DE ACCESO
	Para acceder a un struct de otro archivo este debe de
	estar en mayuscula, go no va a dejar hacer ninguna
	importacion si no se hace "go mod init" en el proyecto

	go siempre toma la carpeta SRC para sacar desde ahí
	las rutas de los paquetes

	se puede usar aliases en el llamado del paquete, colocandolo
	al principio de cada paquete
*/

import (
	"study/platzi/basic/packages"
	f "fmt"
)

func main() {
	var car Packages.Car
	car.Brand = "toryota"
	car.Year = 2032
	f.Println(car)

	Packages.Saludo("David")
}
