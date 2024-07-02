package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	kv := NewKVStore()                  // Create a new instance of KVStore
	reader := bufio.NewReader(os.Stdin) // Create a new reader for user input
	fmt.Println("Please, insert your name and your age: ")
	fmt.Println(" Example: Name age")
	fmt.Println(`.........................`)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')    //Read the user input until newline
		input = strings.TrimSpace(input)       // Trim leading and trailing whitespace
		parts := strings.SplitN(input, " ", 3) // Split the input into up to 3 parts

		if len(parts) < 1 { // Check if the command is valid
			fmt.Println("Invalid command")
			continue
		}

		switch parts[0] { // Determine the command to execute
		case "SET":
			if len(parts) != 3 {
				fmt.Println("Usage: SET Name Age")
				continue
			}
			kv.Set(parts[1], parts[2]) // Set the key-value pair
			fmt.Println("OK")
		case "GET":
			if len(parts) != 2 {
				fmt.Println("Usage: GET Name")
				continue
			}
			value, exists := kv.Get(parts[1]) // Get the value for the given key
			if exists {
				fmt.Println(value)
			} else {
				fmt.Println("Name not found")
			}
		case "DELETE":
			if len(parts) != 2 {
				fmt.Println("Usage: DELETE Name")
				continue
			}
			kv.Delete(parts[1]) // Delete the key-value pair
			fmt.Println("OK")
		case "SAVE":
			if len(parts) != 2 {
				fmt.Println("Usage: SAVE filename")
				continue
			}
			if err := kv.SaveToFile(parts[1]); err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Println("Ok")
			}
		case "LOAD":
			if len(parts) != 2 {
				fmt.Println("Usage: LOAD filenam")
				continue
			}
			if err := kv.LoadFromFile(parts[1]); err != nil {
				fmt.Println("Error: ", err) // PRint an error message if any
			} else {
				fmt.Println("OK")
			}
		case "EXIT":
			return // Exit the program
		default:
			fmt.Println("Unknown command") // Print a message for unknown commands
		}
	}
}
