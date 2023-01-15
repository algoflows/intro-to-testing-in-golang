package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// print a welcome message
	intro()

	// create a channel to receive input from the user
	doneChan := make(chan bool)

	// start a go routine to read from the channel and print the result
	go readUserInput(doneChan)

	// block until the doneChannel receives gets a value
	<-doneChan

	// close the channel
	close(doneChan)

	// say goodbye
	fmt.Println("Goodbye!")
}

func readUserInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		res, done := checkNumbers(scanner)
		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()
	input := scanner.Text()

	if input == "q" {
		return "Goodbye!", true
	}

	n, err := strconv.Atoi(input)
	if err != nil {
		return "Please enter a number or 'q' to quit the program", false
	}

	_, msg := isPrime(n)
	return msg, false
}

func intro() {
	fmt.Println("Welcome to the prime number checker!")
	fmt.Println("Enter a number to see if it is prime.")
	fmt.Println("Enter 'q' to quit.")
}

func prompt() {
	fmt.Print("-> ")
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	// negative numbers are not prime
	if n < 0 {
		return false, "Negative numbers are not prime, by definition!"
	}

	// use the modulus operator repeatedly to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not a prime number
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d", n, i)
		}
	}
	return true, fmt.Sprintf("%d is a prime number", n)
}