So The task-manager will be cli tool built in go.
It will have features like:

- Add a task
  - i can say "add <task name> or <task id>"
- List tasks
  - i can say "list", It will list all tasks with their status(pending/done)
- Remove a task
  - i can say "remove <task name> or <task id>"
- Mark a task as done
  - i can say "done <task name> or <task id>"
- Mark a task as not done
  - i can say "undo <task name> or <task id>"
- Clear all tasks
  - i can say "clear" or "reset"

all tasks will be stored in a json file.
the json file will be stored in the same directory as the cli tool.
