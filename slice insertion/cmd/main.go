package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	m := make(map[int]string)
	var id int = 0 // Declare id outside the loop

	fmt.Println("Storage", m)

	for {
		fmt.Print("Enter your choice: \n")
		fmt.Print("1 : Create\n")
		fmt.Print("2 : Read\n")
		fmt.Print("3 : Update\n")
		fmt.Print("4 : Delete\n")
		fmt.Print("0 : Stop\n")

		// Read the whole line
		input, _ := reader.ReadString('\n')
		// Remove the newline character
		input = strings.TrimSpace(input)

		// Check if it's a single character
		if len(input) != 1 {
			fmt.Println("Incorrect choice - please enter a single character")
			continue
		}

		choice := input[0] // Get the first byte/character

		switch choice {
		case '1':
			fmt.Print("Please input your name: ")
			nameInput, _ := reader.ReadString('\n')
			nameInput = strings.TrimSpace(nameInput)
			m[id] = nameInput
			fmt.Printf("Added: ID %d = %s\n", id, nameInput)
			id++ // Increment after adding
		case '2':
			// Read logic
			if len(m) == 0 {
				fmt.Println("Map is empty")
			} else {
				for i, value := range m {
					fmt.Printf("ID: %d, Name: %s\n", i, value)
				}
			}
		case '3':
			// Update logic
			fmt.Print("Enter ID to update: ")
			idInput, _ := reader.ReadString('\n')
			idInput = strings.TrimSpace(idInput)

			// Convert string to int
			updateID, err := strconv.Atoi(idInput)
			if err != nil {
				fmt.Println("Invalid ID - please enter a number")
				break
			}

			// Check if ID exists in the map
			_, exists := m[updateID]
			if !exists {
				fmt.Printf("ID %d does not exist in the map\n", updateID)
			} else {
				// ID exists, ask for new name
				fmt.Print("Enter new name: ")
				newName, _ := reader.ReadString('\n')
				newName = strings.TrimSpace(newName)
				m[updateID] = newName
				fmt.Printf("Updated ID %d: %s\n", updateID, newName)
			}

		case '4':
			// Delete logic
			fmt.Print("Enter ID to delete: ")
			idInput, _ := reader.ReadString('\n')
			idInput = strings.TrimSpace(idInput)

			// Convert string to int
			deleteID, err := strconv.Atoi(idInput)
			if err != nil {
				fmt.Println("Invalid ID - please enter a number")
				break
			}

			// Check if ID exists before deleting
			_, exists := m[deleteID]
			if !exists {
				fmt.Printf("ID %d does not exist in the map\n", deleteID)
			} else {
				delete(m, deleteID)
				fmt.Printf("Deleted ID %d\n", deleteID)
			}
		case '0':
			// Stop
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Printf("Incorrect choice '%c'\n", choice)
		}
	}
}
