# Task Tracker App

> This project is from [roadmap.sh](https://roadmap.sh/) [Task Tracker](https://roadmap.sh/projects/task-tracker) project.

A **Task Tracker** application is designed to help users manage and organize their daily tasks efficiently. This app includes features like task creation, categorization, updating, and deletion, ensuring an intuitive user experience.

## Features

- **Add Tasks:** Create new tasks with essential details such as title, description, due date, and priority.
- **Update Tasks:** Edit task details as requirements change.
- **Delete Tasks:** Remove tasks that are no longer needed.
- **Mark Tasks as Done:** Indicate tasks that have been completed.
- **Mark Tasks as In Progress:** Change the status of tasks that are currently being worked on.
- **List Tasks:** View all tasks, pending tasks, completed tasks, or tasks in progress.

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/hnfnfl/task-tracker.git
cd task-tracker
```

### 2. Use Makefile To Build

```bash
make build
```

### 3. Run the Application

```bash
./task-tracker
```

## Commands List

- **Add Task:** `add`
  - `-desc`: Task description
- **Update Task:** `update`
  - `-id`: Task ID
  - `-desc`: Task description
- **Delete Task:** `delete`
  - `-id`: Task ID
- **Mark Task as Done:** `done`
  - `-id`: Task ID
- **Mark Task as In Progress:** `in-progress`
  - `-id`: Task ID
- **List Tasks:** `list`
  - no flags: List all tasks
  - `-todo`: List pending tasks
  - `-in-progress`: List tasks in progress
  - `-done`: List completed tasks
