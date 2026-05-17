package task

import (
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done"`
	CreatedAt   time.Time `json:"created_at"`
}
