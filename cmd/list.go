package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"cli_task_manager/internal/task"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "List tasks",
	Args: cobra.ExactArgs(0),

	RunE: func(cmd *cobra.Command, args []string) error {
		err := task.ListTasks()
		if err != nil {
			return fmt.Errorf("failed to list tasks: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}