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

func createTask() Task{
	reader := bufio.NewReader(os.Stdin)
	title, _ := getInput("whats the task title:", reader)
	desc, _ := getInput("give some description:", reader)

	t, err:= AddTask(title, desc)
	if err != nil {
        panic(err)
    }
	return t
}
//c- mark complete r-remove a task  u- undo removal
func Options(t Task) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("how would like to proceed \n (C- mark complete) \n <R-remove a task> \n <X- close prompt window>", reader)

	switch opt {
	case "C" :
		id, _ := getInput("please enter id of task to delete",reader)

		i64, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println("id must be a number")
			Options(t)
		}
		t.Complete(int(i64))

		fmt.Println("")
	
	case "R" :
		id, _ := getInput("please enter id of task to delete",reader)

		i64, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println("id must be a number")
			break
		}

		err = RemoveTask(int(i64))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("")
	
	case "X" :
		fmt.Println("")
	
	default:
		fmt.Println("that was not a valid option...")
		Options(t)	
	}
}

func main() {
	tasks := createTask()
	// Options(tasks)
	fmt.Println(tasks)
}