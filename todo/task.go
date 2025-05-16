package todo

import "fmt"

func AddTask(name string) {
	stmt, _ := DB.Prepare(("INSERT INTO tasks (name) VALUES (?)"))
	_, err := stmt.Exec(name)
	if err != nil {
		fmt.Println("Error adding task:", err)
	} else {
		fmt.Println("Task added")
	}
}

func GetTasks() {
	rows, _ := DB.Query("SELECT id, name, done FROM tasks")
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var done bool
		rows.Scan(&id, &name, &done)
		status := " "
		if done {
			status = "done"
		}
		fmt.Printf("[%s] %d,: %s\n", status, id, name)
	}
}

func MarkDone(id int) {
	stmt, err := DB.Prepare("UPDATE tasks SET done = 1 WHERE id = ?")
	if err != nil {
		fmt.Println("Error preparing update statement:", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		fmt.Println("Error marking task as done:", err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		fmt.Println("No task found with that ID.")
	} else {
		fmt.Println("Task marked as done.")
	}
}

func DeleteTask(id int) {
	stmt, err := DB.Prepare("DELETE FROM tasks WHERE id = ?")
	if err != nil {
		fmt.Println("Error preparing delete statement:", err)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		fmt.Println("Error executing delete:", err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		fmt.Println("No task found with that ID.")
	} else {
		fmt.Println("Task deleted.")
	}
}
