# Task Tracker CLI

A simple command-line application to manage your daily tasks, written in Go. It allows you to add, update, delete, list, and change the status of tasks, storing everything in a local JSON file.

---

## Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/youruser/task-tracker.git
   cd task-tracker
   ```

2. **Build the executable:**
   ```sh
   go build -o tt
   ```

---

## Usage

The main executable is `tt`. You can run any of the following commands from your terminal:

### 1. **Add a task**
```sh
./tt add "Go shopping"
```
**Output:**  
`the task with description Go shopping has been successfully saved. The id is :0`

---

### 2. **Update a task**
```sh
./tt update 0 "Go shopping and cook"
```
**Output:**  
`the task with ID 0 has been updated. The new description is Go shopping and cook.`

---

### 3. **Delete a task**
```sh
./tt delete 0
```
**Output:**  
`the task with ID 0 has been deleted.`

---

### 4. **Mark a task as "in progress"**
```sh
./tt mark-in-progress 1
```
**Output:**  
`the task with ID 1 has been updated to state in-progress.`

---

### 5. **Mark a task as "done"**
```sh
./tt mark-done 1
```
**Output:**  
`the task with ID 1 has been updated to state done.`

---

### 6. **List all tasks**
```sh
./tt list
```
**Sample output:**
```
ID   | Description                              | Status       | Created At           | Updated At
-----+------------------------------------------+--------------+----------------------+----------------------
0    | Go shopping and cook                     | done         | 26-07-2025 12:00:00  | 26-07-2025 12:10:00
1    | Call Carmen                              | in-progress  | 26-07-2025 12:05:00  | 26-07-2025 12:15:00
```

---

### 7. **List tasks by status**
```sh
./tt list done
./tt list todo
./tt list in-progress
```
**Output:**  
Only shows tasks with the specified status.

---

## Task properties

- **id**: Unique identifier for the task
- **description**: Short description of the task
- **status**: Task status (`todo`, `in-progress`, `done`)
- **createdAt**: Creation date and time
- **updatedAt**: Last update date and time

---

## Notes

- The `tasks.json` file is automatically created in the current directory if it does not exist.
- All commands display clear success or error messages.
- If you enter an incorrect command or argument, the program will show help or an error message.

---

## Example full workflow

```sh
./tt add "Read a book"
./tt list
./tt mark-in-progress 0
./tt update 0 "Read a Go book"
./tt mark-done 0
./tt list done
./tt delete 0
```

---
