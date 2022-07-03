package main

import (
	"bufio"
	"email-checker-tool/input"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Type the email: ")
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input %v\n", err)
	}

	i := input.NewEmailInput(scanner.Text())
	email, err := i.NewEmail()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Local and domain email are valid")
	fmt.Println(email)
}
