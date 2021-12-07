package main

import "fmt"

func main() {
	var name string
	fmt.Printf("Enter your name (one-word only): ")
	fmt.Scan(&name)
	fmt.Printf("Hello %v, welcome to the Quiz Game\n", name)

	fmt.Printf("Enter your age: ")
	var age int
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Printf("You are eligible to play the quiz\n")
	} else {
		fmt.Printf("You are not eligible to play the quiz\n")
		return
	}

	score := 0
	num_questions := 2

	// If it reaches here, then the user is eligible to play the quiz
	// Question 1
	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer + answer2 == "RTX 3090" || answer + " " + answer2 == "rtx 3090" {
		fmt.Printf("Correct!\n")
		score++
	} else {
		fmt.Printf("Wrong!\n")
	}

	// Question 2
	fmt.Printf("How many cores does the Ryzen 9 3900x have? ")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Printf("Correct!\n")
		score++
	} else {
		fmt.Printf("Wrong!\n")
	}

	fmt.Printf("Your score is %v out of %v questions\n", score, num_questions)

	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("You scored: %v%%\n", percent)

	fmt.Printf("Thanks for playing our game. Please give us a rating from 1-5:")
	var rating uint
	fmt.Scan(&rating)

}

