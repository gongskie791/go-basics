package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetStringInput(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func GetIntInput(prompt string) int {
	for {
		input := GetStringInput(prompt)
		if num, err := strconv.Atoi(input); err == nil {
			return num
		}
		fmt.Println("Please enter a valid number")
	}
}

func GetFloatInput(prompt string) float64 {
	for {
		input := GetStringInput(prompt)
		if num, err := strconv.ParseFloat(input, 64); err == nil {
			return num
		}
		fmt.Println("Please enter a valid number")
	}
}

func Confirm(prompt string) bool {
	for {
		input := GetStringInput(prompt + " (y/n): ")
		input = strings.ToLower(input)
		if input == "y" || input == "yes" {
			return true
		}
		if input == "n" || input == "no" {
			return false
		}
		fmt.Println("Please enter y or n")
	}
}
