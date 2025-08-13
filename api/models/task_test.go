package models

import (
    "testing"
    "time"
)

func TestNewTask(t *testing.T) {
    title := "Test Task"
    description := "Test Description"
    now := time.Now()

    task := Task{
        ID:          "1",
        Title:       title,
        Description: description,
        Completed:   false,
        CreatedAt:   now,
    }

    if task.Title != title {
        t.Errorf("Expected title %s, but got %s", title, task.Title)
    }

    if task.Description != description {
        t.Errorf("Expected description %s, but got %s", description, task.Description)
    }

    if task.Completed != false {
        t.Errorf("Expected completed to be false, but got %t", task.Completed)
    }

    if !task.CreatedAt.Equal(now) {
        t.Errorf("Expected created at %v, but got %v", now, task.CreatedAt)
    }
}
