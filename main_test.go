package main

import (
	"testing"
	"time"
)

// TestAddTask ensures that adding a task increases the total count
func TestAddTask(t *testing.T) {
	initialTasks := loadTasks()
	addTask("Test task " + time.Now().Format(time.RFC3339))
	updatedTasks := loadTasks()

	if len(updatedTasks) != len(initialTasks)+1 {
		t.Errorf("Expected %d tasks, got %d", len(initialTasks)+1, len(updatedTasks))
	}
}
