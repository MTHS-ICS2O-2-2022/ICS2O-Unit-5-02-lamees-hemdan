// Created By: Lamees Hemdan
// Created On: May 2023
//
// This file contains the main function for the go_app application.

// this program creates a random number, either positive or negative.

package main

import (
	"fmt"
	"math.random"
	"math.floor"
)

func main() {

	fmt.Println("This program creates a random number, either positive or negative.")

	fmt.Println("Choose one of the following options:")
	fmt.Println("1. Positive")
	fmt.Println("2. Negative")

	var choice int
	var positiveButtonChecked bool
	var randomNumber int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	if positiveButtonChecked == true {
		// positive
		randomNumber = math.floor(math.random() * 6)
	} else {
		// negative
		randomNumber = math.floor(math.random() * -6)
	}

	fmt.Println("The random number is: ", randomNumber)
	fmt.Println("\nDone.")
}
