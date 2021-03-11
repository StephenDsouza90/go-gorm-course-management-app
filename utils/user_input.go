package utils

import (
	"bufio"
	"os"
)

// UserInput : ...
func UserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput := scanner.Text()
	return userInput
}
