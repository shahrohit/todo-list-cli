# TODO LIST

A simple command-line interface (CLI) To-Do list app built with Go. This app allows you to add, list, complete, and delete tasks from your To-Do list, all from the terminal.

## Features

- Add tasks: Quickly add tasks to your To-Do list.
- List tasks: View all tasks, showing their completion status.
- Mark tasks as completed: Mark tasks as completed to keep track of progress.
- Delete tasks: Remove tasks from the list when no longer needed.
- Colorful output: Pending tasks appear in yellow, completed tasks appear in green for easier identification.

## Installation

### Prerequisites

- [Go](https://go.dev/) installed on your system.

### Clone the Repository

```bash
git clone https://github.com/shahrohit/todo-list.git
cd todo-list
```

### Install Dependencies

Install the required dependencies (in this case, github.com/fatih/color for colored output):

```bash
go mod tidy
```

### Build the Application

To build the app, simply run the following command:

```bash
go build -o todo
```

## Usage

### Add Todo

To add a new task, use the `-add` flag:

```bash
./todo -add "My Todo"
```

### List All Tasks

To list all tasks, use the `-list` flag:

```bash
./todo -list
```

### Mark a Task as Completed

To mark a task as completed, use the `-complete` flag followed by the `task ID`:

```bash
./todo -complete 1
```

### Delete a Task

To delete task, use the `-delete` flag followed by the `task ID`:

```bash
./todo -delete 1
```

## Example

Here’s a complete example of how you would use the To-Do app:

```bash
# Add some tasks
./todo -add "Buy Apple"
# Output:
# Task added: Buy Apple

./todo -add "Play Games"
# Output:
# Task added: Play Games


# List all tasks
./todo -list
# Output:
# 1. Buy Apple [Pending]
# 2. Play Games [Pending]

# Mark a task as completed
./todo -complete 1
# Output:
# Taks Compledted : 1

# List all tasks
./todo -list
# Output:
# 1. Buy Apple [Completed]
# 2. Play Games [Pending]

# Delete a task
./todo -delete 2
# Output:
# Task deleted: 2
```
