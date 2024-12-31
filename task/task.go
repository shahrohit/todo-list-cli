package task

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
)


type Task struct{
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

type TaskManager struct {
	FilePath string
	Tasks []Task
}

func (tm *TaskManager) Load() error {
	file, err := os.Open(tm.FilePath)
	if err != nil {
		if os.IsNotExist(err){
			tm.Tasks = []Task{}
			return nil
		}
		return fmt.Errorf("error loading task %v", err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tm.Tasks); err != nil {
		if(err.Error() != "EOF"){
			return fmt.Errorf("error decoding task : %v", err)
		}
	}

	return nil
}


func (tm *TaskManager) Save() error {
	file, err := os.Create(tm.FilePath)
	if err != nil {
		return fmt.Errorf("error saving tasks: %v",err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(&tm.Tasks); err != nil {
		return fmt.Errorf("error encoding data : %v", err)
	}

	return nil
}

// Add New Task
func (tm *TaskManager) Add(title string) error {

	newTask := Task{
		Title: title,
		Completed: false,
	}

	tm.Tasks = append(tm.Tasks, newTask)

	return tm.Save();
}

// List all task
func (tm *TaskManager) List() {

	if(len(tm.Tasks) == 0){
		fmt.Println("No Task.")
		return
	}

	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()


	for id,task := range tm.Tasks {
		status := yellow("pending")
		if task.Completed { status = green("completed")}

		fmt.Printf("%d. %s [%s]\n",id + 1, task.Title, status)
	}
}

// Mark as completed
func (tm *TaskManager) MarkCompleted(id int) error {
	if id <= 0 || id > len(tm.Tasks){
		return fmt.Errorf("task with Id %d not found", id)
	}
	tm.Tasks[id - 1].Completed = true
	return tm.Save()
}

// Delete Task
func (tm *TaskManager) Delete(id int) error {
	if id <= 0 || id > len(tm.Tasks){
		return fmt.Errorf("task with Id %d not found", id)
	}
	tm.Tasks = append(tm.Tasks[:id - 1], tm.Tasks[id:]...)
	return tm.Save()
}