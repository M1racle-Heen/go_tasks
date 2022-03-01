package main

import (
	"errors"
	"fmt"
)

func main() {
	messages := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	messages = append(messages, "3")
	// fmt.Println(messages[4:])
	printMessage(messages)
	fmt.Println(messages)
}

func printMessage(message []string) error {
	if len(message) == 0 {
		return errors.New("message is empty")
	}
	message = append(message, "3")
	message[3] = "12"
	// fmt.Println(message)
	return nil
}
