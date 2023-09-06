package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	var username string
	words := []string{"Apple", "Orange", "Blueberry", "Ashir"}
	max := len(words)
	randomNumber := rng.Intn(max-0) - 0
	fmt.Println(randomNumber)
	fmt.Println("Welcome to Hangman Game")
	fmt.Println("Enter your name")
	fmt.Scan(&username)
	fmt.Println("Get ready to guess this fruit")

	var dashString string = ""
	for i := 0; i < len(words[randomNumber]); i++ {
		dashString = dashString + "_ "
	}
	fmt.Println(dashString)

}
