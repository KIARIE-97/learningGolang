package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)
//notification will be sebt via a channel to signal a job is done
type Notification struct{
	TaskId int
	TaskName string
	success bool
	Message string
	// Duration time.ParseDuration()
}
//channel
var notifications = make(chan Notification)
// now tracking of running / tasking being processed so as not to do same task twice
var runningTasks = make(map[int]bool)
var runMutex sync.Mutex

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func processTask(task Task) {
	//mark task as running
	runMutex.Lock()
	runningTasks[task.ID] = true
	runMutex.Unlock()


	fmt.Println("started procesing task:", task.Title, task.ID )

	go func () {

	time.Sleep(6* time.Second)

	err := Complete(task.ID)
	//notification
	notif := Notification {
		TaskId: task.ID,
		TaskName: task.Title,
		success: err == nil,
	}
	if err != nil {
		notif.Message = err.Error()
	}else {
		notif.Message = "task completed successfully"
	}
	//now sendn the notification via channel
	notifications <- notif

	runMutex.Lock()
	delete(runningTasks, task.ID)
	runMutex.Unlock()

	}()
fmt.Println("continue using the cli while this runs")
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
		fmt.Println("[P] process a task")
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

		case "P" :
			idStr, _ := getInput("Enter task id to mark as complete",reader)
			id, err := strconv.Atoi(idStr)
			if err!= nil {
				fmt.Println("invalid Id.please enter a number")
				continue
			}
			task, err := getById(id)
			    if err != nil {
                fmt.Printf("❌ %v\n", err)
                continue
            }
			go processTask(*task)	

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
		// err = saveAllTask([]Task{})
		// if err != nil {
		// 	fmt.Printf("error initializing: %v\n", err)
		// 	return
		// }
		// fmt.Println("welcome! created new task file")
		//seed
		  AddTask("Test task 1", "for testing concurrency")
        AddTask("Test task 2", "another test")
	}
	//start the cli
	Options()
}