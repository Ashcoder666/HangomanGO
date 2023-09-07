package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	var username string
	words := []string{"apple", "orange", "blueberry", "ashir"}
	max := len(words)
	hangCount := 0
	randomNumber := rng.Intn(max-0) - 0
	fmt.Println("Welcome to Hangman Game")
	fmt.Println("Enter your name")
	fmt.Scan(&username)
	fmt.Println("Get ready to guess this fruit")

	var dashString string = ""
	var guessLetter string
	length := len(words[randomNumber])
	for i := 0; i < length; i++ {
		dashString = dashString + "_"
	}
	fmt.Printf("Number of Letters of word is %v \n ", length)
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
		if j := strings.Contains(words[randomNumber], guessLetter); j == false {
			hangCount++
			filePath := fmt.Sprintf("./states/state%d.txt", hangCount)
			file, err := os.Open(filePath)
			if err != nil {
				log.Fatal(err)
			}

			defer file.Close()

			scanner := bufio.NewScanner(file)

			// Iterate over each line in the file.
			for scanner.Scan() {
				line := scanner.Text()
				fmt.Println(line)
			}

			// Check for any scanning errors.
			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading file:", err)
			}
		}

		if hangCount == 9 {
			fmt.Println("GAME OVER")
			break
		}
	}

}
