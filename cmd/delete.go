package cmd

import (
	"cli_task_manager/internal/task"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use: "delete [task id]",
	Short: "Deletes task with specified ID.",
	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid task id: %w", err)
		}
		err = task.DeleteTask(id)
		if err != nil {
			return fmt.Errorf("failed to delete task: %w", err)
		}

		fmt.Printf("Task %d deleted successfully!\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}