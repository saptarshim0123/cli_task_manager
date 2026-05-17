package task

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

func AddTask(desc string) error {
	filepath, err := GetStoragePath()
	if err != nil {
		return err
	}

	// load existing tasks
	tasks, err := Load(filepath)
	if err != nil {
		return err
	}

	// generate new id
	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	// build new Task struct
	newTask := Task{
		ID:          newID,
		Description: desc,
		IsDone:      false,
		CreatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	return Save(filepath, tasks)
}

func ListTasks() error {
	filepath, err := GetStoragePath()
	if err != nil {
		return err
	}

	// load existing tasks
	tasks, err := Load(filepath)
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("🎉 No tasks found. You're all caught up!")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	fmt.Fprintln(w, "ID\tSTATUS\tDESCRIPTION\tCREATED")
	fmt.Fprintln(w, "--\t------\t-----------\t-------")

	for _, t := range tasks {
		status := "  [ ]"
		if t.IsDone {
			status = "  [X]"
		}

		dateStr := t.CreatedAt.Format("2006-01-02 15:04")
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", t.ID, status, t.Description, dateStr)
	}

	w.Flush()
	return nil
}

func CompleteTask(id int) error {
	filepath, err := GetStoragePath()
	if err != nil {
		return err
	}

	// load existing tasks
	tasks, err := Load(filepath)
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].IsDone = true
			return Save(filepath, tasks)
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

func DeleteTask(id int) error {
	filepath, err := GetStoragePath()
	if err != nil {
		return err
	}

	// load existing tasks
	tasks, err := Load(filepath)
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return Save(filepath, tasks)
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}
