package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"cli_task_manager/internal/task"
)

var addCmd = &cobra.Command{
	Use: "add [task description]",
	Short: "Add a new task to your list",
	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		desc := args[0]
		err := task.AddTask(desc)
		if err != nil {
			return fmt.Errorf("failed to add task: %w", err)
		}

		fmt.Printf("✅ Added task: %q\n", desc)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}