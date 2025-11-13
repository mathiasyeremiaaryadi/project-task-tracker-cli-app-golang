
# Project Title

This is a CLI application for task tracker to manage your todo task. 

Project from: https://roadmap.sh/projects/task-tracker



## Features

- Add task
- Update task description
- Delete task
- Mark task status (in-progress/done)
- List task by the status (in-progress/done) 


## Documentation

### Add Task
```
task-cli add "Buy groceries"
```

### Updating and deleting tasks
```
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1
```

### Marking a task as in progress or done
```
task-cli mark-in-progress 1
task-cli mark-done 1
```

### Listing all tasks
```
task-cli list
```

### Listing tasks by status
```
task-cli list done
task-cli list todo
task-cli list in-progress
```


## Run Locally

Clone the project

```bash
https://github.com/mathiasyeremiaaryadi/project-task-tracker-cli-app-golang.git
```

Go to the project directory

```bash
  cd my-project
```

Install dependencies

```bash
go build
```

Start the server

```bash
./task-cli [command]
```
