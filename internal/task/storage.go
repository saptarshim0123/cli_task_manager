package task

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
)

func Load(filename string) ([]Task, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Task{}, nil
		}
		return nil, err
	}

	defer file.Close()

	// Read bytes
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// If file is empty
	if len(bytes) == 0 {
		return []Task{}, nil
	}

	// Unmarshal bytes into Task
	var tasks []Task
	err = json.Unmarshal(bytes, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetStoragePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".todo.json"), nil
}

func Save(filename string, tasks []Task) error {
	// Marshal the tasks into formatted JSON
	bytes, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	// Write bytes to file
	err = os.WriteFile(filename, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
