package utility

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// retrieveUserInput is a function that can retrieve user input in form of string. By default,
// it will prompt the user. In test, you can replace this with code that returns the appropriate response.
var retrieveUserInput = func(message string) (string, error) {
	return readUserInput(os.Stdin, message)
}

// readUserInput is a io.Reader to read user input from. It is meant to allow simplified testing
// as to-be-read inputs can be injected conveniently.
func readUserInput(in io.Reader, message string) (string, error) {
	reader := bufio.NewReader(in)
	YellowConfirm("Are you sure you want to " + message + " (y/N) ? ")
	answer, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	answer = strings.TrimRight(answer, "\r\n")

	return strings.ToLower(answer), nil
}

// AskForConfirm parses and verifies user input for confirmation.
func AskForConfirm(message string) error {
	answer, err := retrieveUserInput(message)
	if err != nil {
		Error("Unable to parse users input: %s", err)
	}

	if answer != "y" && answer != "ye" && answer != "yes" {
		return fmt.Errorf("invalid user input")
	}

	return nil
}

// AskForConfirmDelete builds a message to ask the user to confirm delete
// a resource and then sends it through to AskForConfirm to
// parses and verifies user input.
func AskForConfirmDelete(resourceType string) error {
	message := fmt.Sprintf("delete this %s?", resourceType)
	err := AskForConfirm(message)
	if err != nil {
		return err
	}

	return nil
}
