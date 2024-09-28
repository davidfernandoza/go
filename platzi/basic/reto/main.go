package main

import "fmt"

func main() {

	for {
		const message string = "Select an option: \n [1] Number is par? \n [2] Login \n [3] Exit \n =>"
		var selectOption int = 0
		fmt.Print(message)
		fmt.Scanln(&selectOption)
		if selectOption == 1 {
			numberIsPar()
			break
		} else if selectOption == 2 {
			login()
			break
		} else if selectOption == 3 {
			fmt.Printf("***Good bye*** \n")
			break
		}
		fmt.Printf("***Selection invalid*** \n")

	}

}

func numberIsPar() bool {
	var number int = 0
	fmt.Printf("***Your as selected option: number is par***\n")

	fmt.Printf("write a number: \n =>")
	fmt.Scanln(&number)
	if number == 0 {
		fmt.Printf("Number is zero")
		return false
	} else if number%2 == 0 {
		fmt.Printf("Number is par")
	} else {
		fmt.Printf("Number is not par")
	}
	return true
}

func login() bool {
	var user string = ""
	var password string = ""
	var credentialUser, credentialPassword string = "Dario", "12345"
	fmt.Printf("***Your as selected option: login***\n")
	fmt.Printf("user: \n =>")
	fmt.Scanln(&user)
	fmt.Printf("password: \n =>")
	fmt.Scanln(&password)

	if user == credentialUser && password == credentialPassword {
		fmt.Printf("***Login successfull***")
		return true
	}
	fmt.Printf("***Login faild***")
	return false

}
