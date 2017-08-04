package main

import "fmt"

func ex4() {
	var smallNumber, bigNumber int
	fmt.Print("Enter small number: ")
	fmt.Scan(&smallNumber)
	fmt.Print("Enter big number: ")
	fmt.Scan(&bigNumber)
	if bigNumber >= smallNumber {
		fmt.Println("Remainder is:", bigNumber%smallNumber)
	} else {
		fmt.Println("Invalid")
		ex4()
	}
}
