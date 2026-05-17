package task

import "testing"

func TestCompleteTask(t *testing.T) {
	err := CompleteTask(2)

	if err != nil {
		t.Errorf("AddTask failed: %v", err)
	}
}