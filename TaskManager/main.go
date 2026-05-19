package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var filename string = "tasks.json"

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createTask() {
	reader := bufio.NewReader(os.Stdin)
	title, _ := getInput("whats the task title:", reader)
	desc, _ := getInput("give some description:", reader)

	_, err:= AddTask(title, desc)
	if err != nil {
        fmt.Printf("Error adding task: %v\n", err)
    }
}
//c- mark complete r-remove a task  u- undo removal
func Options() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n What would like to do?")
		fmt.Println("[C] mark task as complete")
		fmt.Println("[R] remove a task")
		fmt.Println("[A] add a task")
		fmt.Println("[X] exit")

		opt, _ := getInput("\n your choice",reader)
		opt = strings.ToUpper(opt)

		switch opt {
		case "C" :
			idStr, _ := getInput("Enter task id to mark as complete",reader)
			id, err := strconv.Atoi(idStr)
			if err!= nil {
				fmt.Println("invalid Id.please enter a number")
				continue
			}
			err = Complete(id)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			
		case "R" :
			idStr, _ := getInput("Enter task id to remove",reader)
			id, err := strconv.Atoi(idStr)
			if err!= nil {
				fmt.Println("invalid Id.please enter a number")
				continue
			}
	
			//confirm deletion
			confirm, _ := getInput("are you sure you want to delete task #%d? (y/N)", reader)
			if strings.ToLower(confirm) != "y" {
				fmt.Println("deletion cancelled")
				continue
			}

			err = RemoveTask(id)
			if err != nil {
				fmt.Println(err)
			}
		
		case "A" :
			createTask()

		case "X" :
			fmt.Println("\n goodbye >>>>")
			return
		default:
			fmt.Println("Invalid option.please choose C, R, A or X")
			Options()	
		}
	}


}

func main() {
	fmt.Println("=== TaskManager CLI ===")
	//create a new file if it doesnt exist
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		//create empty task file
		err = saveAllTask([]Task{})
		if err != nil {
			fmt.Printf("error initializing: %v\n", err)
			return
		}
		fmt.Println("welcome! created new task file")
	}
	//start the cli
	Options()
}