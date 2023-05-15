// Created By: Lamees Hemdan
// Created On: May 2023
//
// This file contains the main function for the go_app application.

// this program creates a random number, either positive or negative.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("This program creates a random number, either positive or negative.")

	fmt.Println("Choose one of the following options:")
	fmt.Println("1. Positive")
	fmt.Println("2. Negative")

	var userChoice string
	var randomNumber int
	var positiveNumber = "positive"
	var negativeNumber = "negative"

	fmt.Print("Enter your choice: ")
	fmt.Scanln(&userChoice)

	if userChoice == positiveNumber {
		rand.Seed(time.Now().UnixNano())
		min := 1
		max := 6
		randomNumber = rand.Intn(max-min) + min
	} else if userChoice == negativeNumber {
		rand.Seed(time.Now().UnixNano())
		min := -6
		max := -1
		randomNumber = rand.Intn(max-min) + min
	} else {
		fmt.Println("Invalid input.")
	}
	// output
	fmt.Println("The random number is: ", randomNumber)
	fmt.Println("\nDone.")
}
