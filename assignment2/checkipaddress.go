package assignment2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func CheckIPAddress() error {
	scanner := bufio.NewScanner(os.Stdin)
	var input string

	fmt.Println("\n------------ Program IP Address Checker. ------------")
	fmt.Print("Enter IP Address: ")

	if scanner.Scan() {
		input = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	result := validateIPAddress(input)

	fmt.Printf("Result: %t", result)
	fmt.Println("\n------------------------------------------------------")

	return nil
}

func validateIPAddress(input string) bool {
	re := regexp.MustCompile(`^([1-9]{1,3})\.([0-9]{1,3})\.([0-9]{1,3})\.([0-9]{1,3})$`)
	return re.MatchString(input)
}
