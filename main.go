package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	var username string
	words := []string{"apple", "orange", "blueberry", "ashir"}
	max := len(words)
	randomNumber := rng.Intn(max-0) - 0
	fmt.Println(randomNumber)
	fmt.Println("Welcome to Hangman Game")
	fmt.Println("Enter your name")
	fmt.Scan(&username)
	fmt.Println("Get ready to guess this fruit")

	var dashString string = ""
	var guessLetter string
	for i := 0; i < len(words[randomNumber]); i++ {
		dashString = dashString + "_"
	}

	for dashString != words[randomNumber] {

		for i := 0; i < len(words[randomNumber]); i++ {
			if string(guessLetter) == string(words[randomNumber][i]) {
				dashString = dashString[:i] + guessLetter + dashString[i+1:]
			}
		}

		fmt.Println(dashString)
		if dashString == words[randomNumber] {
			fmt.Printf("congrats %v your guess for word %v is correct \n", username, dashString)
			break
		}
		fmt.Println("Enter your letter")
		fmt.Scan(&guessLetter)
	}

}
