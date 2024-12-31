package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/shahrohit/todo-list/task"
)


func main(){
	taskManager := task.TaskManager{FilePath: "tasks.json"}

	
	if err := taskManager.Load(); err != nil {log.Fatal(err)}
	add := flag.String("add", "", "Add new task")
	list := flag.Bool("list", false, "List all the task")
	complete := flag.Int("complete", 0, "Mark a task as complete")
	delete := flag.Int("delete", 0, "Delete a task")
	

	flag.Parse()

	switch {
	case *add != "":
		if err := taskManager.Add(*add); err != nil {
			fmt.Println("Error adding task:", err)
		} else {
			s := fmt.Sprintf("Task added: %s", *add)
			color.Green(s)
		}
	case *list:
		taskManager.List()
	case *complete > 0:
		if err := taskManager.MarkCompleted(*complete); err != nil {
			fmt.Println("Error completing task:", err)
		} else {
			s := fmt.Sprintf("Taks Compledted : %d", *complete)
			color.Green(s)
		}
	case *delete > 0:
		if err := taskManager.Delete(*delete); err != nil {
			fmt.Println("Error deleting task:", err)
		} else {
			s := fmt.Sprintf("Task deleted: %d", *delete)
			color.Red(s)
			
		}
	default:
		fmt.Println("Usage:")
		flag.PrintDefaults()
	}



	
}