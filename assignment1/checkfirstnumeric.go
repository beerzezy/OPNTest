package assignment1

import (
	"fmt"
	"regexp"
)

func CheckFirstNumeric() {
	var input string

	fmt.Println("\n------- Program Check First Numeric of String. -------")
	fmt.Print("Enter String: ")
	fmt.Scanln(&input)

	result := findFirstNumeric(input)

	if result == "" {
		result = "not found numeric."
	}
	fmt.Printf("Result: %s", result)
	fmt.Println("\n------------------------------------------------------")
}

func findFirstNumeric(input string) string {
	re := regexp.MustCompile("[0-9]{1}")
	return re.FindString(input)
}
