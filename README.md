# Number Guessing Game (CLI)

This is a simple command-line number guessing game built using **Go (Golang)**. The application generates a random number between 1 and 100, and the player must guess it within a limited number of attempts based on the selected difficulty level.

This project is a solution to the challenge found on [roadmap.sh/projects/number-guessing-game](https://roadmap.sh/projects/number-guessing-game).

## Features

- **Random Number Generation:** The computer randomly picks a number between 1 and 100 at the start of each game.
- **Difficulty Levels:** Players can select their preferred difficulty, which alters their total attempts:
  - **Easy:** 10 chances
  - **Medium:** 5 chances
  - **Hard:** 3 chances
- **Dynamic Feedback:** The game informs you if your guess is too high or too low.
- **Input Validation:** Prevents loss of attempts if a non-numeric string or unexpected text is accidentally entered.

## Requirements

To run this project locally, you need to have **Go** installed on your system.

- [Download and install Go](https://go.dev)

## How to Run

1. Clone this repository to your local machine:
   ```bash
   git clone https://github.com
   ```

2. Navigate into the project directory:
   ```bash
   cd YOUR_REPOSITORY_NAME
   ```

3. Execute the code directly using the Go CLI:
   ```bash
   go run main.go
   ```

## Project Structure

```text
├── main.gotext    # Main game engine and application entry point
└── README.md      # Project documentation
```

## Sample Output

```text
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.

Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)
Enter your choice: 2

Enter your guess: 50
Incorrect! The number is less than 50

Enter your guess: 25
Incorrect! The number is greater than 25

Enter your guess: 35
Incorrect! The number is less than 35

Enter your guess: 30
Congratulations! You guessed the correct number in 4 attempts.
```
