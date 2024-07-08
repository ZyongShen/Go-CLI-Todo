package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"bufio"
)

func startFile() {
	f, err := os.OpenFile("todo.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Failed to open file")
	}
	defer f.Close() // close once function is over

	fmt.Println("Successfully started or opened a file")

	// read file
	b, err := os.ReadFile("todo.txt")
	if (err != nil) {
		log.Fatal(err)
		fmt.Println("Failed to read file")
	} else {
		instructions := "Add X to [ ] to mark as completed"
		fileContents := string(b)

		// if does not contain instructions, then write instructions, else do nothing
		if (!strings.Contains(fileContents, instructions)) {
			// reset file
			f.Truncate(0)

			// write existing content
			f.WriteString(fileContents)
		}

		fmt.Println("Successfully started or opened a file")
	}
}

func displayTodo() {
	count := 1
	b, err := os.ReadFile("todo.txt")
	if (err != nil) {
		log.Fatal(err)
		fmt.Println("Faied to read file")
	}
	fileContents := string(b)
	contentsList := strings.Split(fileContents, "\n")

	fmt.Println("\n")
	fmt.Println("Your todo list")
	for _, value := range contentsList {
		if (value != "") {
			fmt.Println(strconv.Itoa(count) + "." + value)
			count += 1
		}
	}
}

// add a new task
func addTask(task string) {
	// check if file exists or create a new file first
	// load the file
	f, err := os.OpenFile("todo.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Failed to open file")
	}

	defer f.Close() // close once function is over

	task = "[ ] " + task + "\n"
	_, err = f.WriteString(task)
	if (err != nil) {
		log.Fatal(err)
		fmt.Println("Failed to write to file")
	}

	fmt.Println("Successfully wrote to file")
}

// parse the files and remove lines starting with 
func removeTask(taskNumber int) {
	idx := taskNumber - 1
	f, err := os.OpenFile("todo.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if (err != nil) {
		log.Fatal(err)
		fmt.Println("Failed to open file")
	}

	defer f.Close()

	b, err := os.ReadFile("todo.txt")
	if (err != nil) {
		log.Fatal(err)
	}
	fileContents := string(b)
	contentsList := strings.Split(fileContents, "\n")
	contentsList = append(contentsList[:idx], contentsList[idx+1:]...)
	result := ""
	for _, value := range contentsList {
		result += (value + "\n")
	}
	f.Truncate(0)
	f.WriteString(result)
}

func completeTask(taskNumber int) {
	idx := taskNumber - 1

	f, err := os.OpenFile("todo.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if (err != nil) {
		log.Fatal(err)
		fmt.Println("Failed to open file")
	}

	defer f.Close()

	b, err := os.ReadFile("todo.txt")
	if (err != nil) {
		log.Fatal(err)
	}
	fileContents := string(b)
	contentsList := strings.Split(fileContents, "\n")

	content := contentsList[idx]
	content = strings.Replace(content, "[ ]", "[X]", -1)
	contentsList[idx] = content

	result := ""

	for _, value := range contentsList {
		result += (value + "\n")
	}
	f.Truncate(0)
	f.WriteString(result)
}

func main() {

	startFile()

	run := true
	var action string
	var task string
	var taskNumber int

	scanner := bufio.NewScanner(os.Stdin)

	for run {
		fmt.Println("\n1. Add task\n2. Complete task\n3. Remove task\n4. Display tasks\n5. Exit")
		fmt.Print("Selection (enter number or enter Exit): ")
		scanner.Scan()
		action = scanner.Text()
		fmt.Println("Your action: ", action)
		switch action {
			case "1":
				fmt.Print("Enter task: ")
				scanner.Scan()
				task = scanner.Text()
				addTask(task)

				displayTodo()
			case "2":
				fmt.Print("Enter task number: ")
				fmt.Scan(&taskNumber)
				completeTask(taskNumber)
				
				displayTodo()
			case "3":
				fmt.Print("Enter task number: ")
				fmt.Scan(&taskNumber)
				removeTask(taskNumber)

				displayTodo()
			case "4":
				displayTodo()
			case "Exit":
				run = false
		}
	}
}