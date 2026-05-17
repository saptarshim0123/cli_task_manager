package cmd

import (
	"cli_task_manager/internal/task"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [task id]",
	Short: "Change selected task status to complete",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid task id: %w", err)
		}
		err = task.CompleteTask(id)
		if err != nil {
			return fmt.Errorf("failed to complete task: %w", err)
		}

		fmt.Printf("✅ Task %d marked as complete!\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
