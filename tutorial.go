package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome! Do you want to play a game?")

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("Ok, you can play my game!")
	} else {
		fmt.Println("You are under 18 and can't play this game!")
		return // if age are under 18, program will stop
	}

	score := 0
	num_questions := 2

	fmt.Printf("What is better, the RTX 3060 or RTX 3070 for mining? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer+" "+answer2 == "RTX 3060" || answer+" "+answer2 == "rtx 3060" {
		fmt.Println("Correct!")
		score++ // score + 1
	} else {
		fmt.Println("Wrong answer!")
	}

	fmt.Println("How many cores does the Ryzen 9 3900h have? ")
	var cores int
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Wrong answer!")
	}

	fmt.Printf("You scored %v, out of %v!\n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100 // convert int to decimal
	fmt.Printf("You scored: %v%%.", percent)
}
