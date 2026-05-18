package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

//small letter fields are auomaticallay private so for export make them capital
//fields must start with capital letters for JSON export
type Task struct {
	ID          int		`json:"id"`
	Title       string	`json:"title"`
	Description string	`json:"description"`
	Completed   bool	`json:"completed"`
}


var filename string = "tasks.json"

//load task- waiting area avoids overwriting
func loadTasks() ([]Task, error) {
	data, err := os.ReadFile(filename)
	if err !=nil {
		if os.IsNotExist(err) {
			return []Task{}, nil //no file yet return empty string
		}
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

// addtask
func AddTask(title string, desc string) (Task, error){
	//load existing task
	tasks, err := loadTasks()
	if err!=nil {
		return Task{}, err
	}

	//gen new id
	var nextId int = 1
	if len(tasks) > 0 {
		nextId = tasks[len(tasks)-1].ID +1
	}

	//create new task
	t := Task{
		ID:          nextId,
		Title:       title,
		Description: desc,
		Completed:   false,
	}
	tasks = append(tasks, t)
	err = saveAllTask(tasks)
	return t, err
}

// save to json
func  saveAllTask(tasks []Task) error{
	data, err := json.MarshalIndent(tasks, "", " ")
	if err !=nil {
		return err
	}

 return os.WriteFile(filename, data, 0644)

}

//mark complete
func (t *Task) Complete(id int) error {
	tasks,err := loadTasks()
	if err != nil {
		return fmt.Errorf("failed to load tasks: %w",err)
	}

	found := false

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			found = true
			fmt.Printf("✅ Task #%d marked as complete\n", id)
            break
		}
	}

	if !found {
        return fmt.Errorf("task with ID %d not found", id)
    }
	return saveAllTask(tasks)
}

//removetask - use of slice
func RemoveTask(id int) error {
	//load current tasks
	tasks, err := loadTasks()
	if err != nil {
		return fmt.Errorf("failed to load tasks: %w",err)
	}
	//find and remove the task
	originalLength := len(tasks)

	index := -1
	for i, task := range tasks{
		if task.ID ==id {
			index = i
			break
		}
	}
	if index == -1 {
		return  fmt.Errorf("task with id %d not found", id)
	}
	//remove
	tasks = append(tasks[index:], tasks[index+1:]...)

	//check if anything was removed
	if len(tasks) == originalLength {
		return errors.New("no task was removed")
	}
	err = saveAllTask(tasks)
	    if err != nil {
        return fmt.Errorf("failed to save tasks after removal: %w", err)
    }
    
    fmt.Printf("✅ Successfully removed task with ID %d\n", id)
    return nil

}
//undo last removal
