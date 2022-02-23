package assignment2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func CheckIPAddress() (bool, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var input string

	fmt.Println("\n------------ Program IP Address Checker. ------------")
	fmt.Print("Enter IP Address: ")

	if scanner.Scan() {
		input = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return false, err
	}

	result := validateIPAddress(input)

	return result, nil
}

func validateIPAddress(input string) bool {
	re := regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	return re.MatchString(input)
}
