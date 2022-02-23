package assignment1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func CheckFirstNumeric() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var input string

	fmt.Println("\n------- Program Check First Numeric of String. -------")
	fmt.Print("Enter String: ")

	if scanner.Scan() {
		input = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	result := findFirstNumeric(input)
	if result == "" {
		result = "Numeric not found."
	}

	return result, nil
}

func findFirstNumeric(input string) string {
	re := regexp.MustCompile("[0-9]{1}")
	return re.FindString(input)
}
