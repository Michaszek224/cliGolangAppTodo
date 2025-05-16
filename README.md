# Go CLI Todo List

A simple command-line todo list application written in Go, using SQLite for storage.

## Features

- Add tasks
- List tasks
- Mark tasks as done
- Delete tasks

## Usage

- Add a new task <br>
<code>go run main.go add "Your task description"</code>

- List all tasks <br>
<code>go run main.go list</code>

- Mark task as done by ID <br>
<code>go run main.go done <task_id></code>

- Delete task by ID <br>
<code>go run main.go delete <task_id></code>
