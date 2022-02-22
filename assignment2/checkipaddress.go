package assignment2

import (
	"fmt"
	"regexp"
)

func CheckIPAddress() {
	var input string

	fmt.Println("\n------------ Program IP Address Checker. ------------")
	fmt.Print("Enter IP Address: ")
	fmt.Scanln(&input)

	result := validateIPAddress(input)

	fmt.Printf("Result: %t", result)
	fmt.Println("\n------------------------------------------------------")
}

func validateIPAddress(input string) bool {
	re := regexp.MustCompile(`^([1-9]{1,3})\.([0-9]{1,3})\.([0-9]{1,3})\.([0-9]{1,3})$`)
	return re.MatchString(input)
}
