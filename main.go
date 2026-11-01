package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")

	// Correctly initialize random number between 1 and 100
	num := rand.Intn(100) + 1
	maxTries := 5

	fmt.Println("\nPlease select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")
	fmt.Print("Enter your choice: ")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("invalid string!")
		return
	}

	input = strings.TrimSpace(input)

	choices, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Choose a number for your choices!")
	}

	switch choices {
	case 1:
		maxTries = 10
	case 2:
		maxTries = 5
	case 3:
		maxTries = 3
	default:
		fmt.Printf("Invalid choice using default %d\n", maxTries)
	}

	for tries := 1; tries <= maxTries; tries++ {
		fmt.Print("Enter your guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("invalid string!")
			return
		}

		input = strings.TrimSpace(input)

		userTry, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid!, try again")
			tries--
			continue
		}

		if userTry == num {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", tries)
			return
		} else if userTry > num {
			fmt.Printf("Incorrect! The number is less than %d\n", userTry)
			continue
		} else {
			fmt.Printf("Incorrect! The number is greater than %d\n", userTry)
			continue
		}
	}

	fmt.Printf("GAME OVER! tries: %d\n", maxTries)
}
