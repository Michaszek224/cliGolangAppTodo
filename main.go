package main

import (
	"fmt"
	"os"
	"strconv"
	"todo/todo"
)

func main() {
	todo.InitDB("tasks.db")

	if len(os.Args) < 2 {
		fmt.Println("Usage: todo [arguments]")
		fmt.Println("Arguments:")
		fmt.Println("---add + name")
		fmt.Println("---list")
		fmt.Println("---done + id")
		fmt.Println("---delete + id")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please use:")
			fmt.Println("--add + name")
			return
		}
		name := os.Args[2]
		todo.AddTask(name)
	case "list":
		todo.GetTasks()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Please use:")
			fmt.Println("--done + id")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		todo.MarkDone(id)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please use:")
			fmt.Println("--delete + id")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		todo.DeleteTask(id)
	default:
		fmt.Println("Unkown command")
	}
}
