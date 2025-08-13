package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"task-api/handlers"
	"task-api/models"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateTask(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")

	task := models.Task{Title: "Test Task", Description: "Test Description"}
	body, _ := json.Marshal(task)

	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var createdTask models.Task
	json.Unmarshal(rr.Body.Bytes(), &createdTask)

	if createdTask.Title != task.Title {
		t.Errorf("handler returned unexpected body: got %v want %v",
			createdTask.Title, task.Title)
	}
}

func TestListTasks(t *testing.T) {
	// Ensure there's at least one task
	handlers.CreateTaskForTest(models.Task{Title: "Test Task", Description: "Test Description"})

	r := mux.NewRouter()
	r.HandleFunc("/tasks", handlers.ListTasks).Methods("GET")

	req, _ := http.NewRequest("GET", "/tasks", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var tasks []models.Task
	json.Unmarshal(rr.Body.Bytes(), &tasks)

	if len(tasks) == 0 {
		t.Errorf("handler returned empty list of tasks")
	}
}

func TestGetTask(t *testing.T) {
	createdTask := handlers.CreateTaskForTest(models.Task{Title: "Test Task", Description: "Test Description"})

	r := mux.NewRouter()
	r.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET")

	req, _ := http.NewRequest("GET", "/tasks/"+createdTask.ID, nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var task models.Task
	json.Unmarshal(rr.Body.Bytes(), &task)

	if task.ID != createdTask.ID {
		t.Errorf("handler returned wrong task: got %v want %v",
			task.ID, createdTask.ID)
	}
}

func TestUpdateTask(t *testing.T) {
	createdTask := handlers.CreateTaskForTest(models.Task{Title: "Test Task", Description: "Test Description"})

	r := mux.NewRouter()
	r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")

	updatedTask := models.Task{Title: "Updated Task", Description: "Updated Description", Completed: true}
	body, _ := json.Marshal(updatedTask)

	req, _ := http.NewRequest("PUT", "/tasks/"+createdTask.ID, bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var task models.Task
	json.Unmarshal(rr.Body.Bytes(), &task)

	if task.Title != updatedTask.Title {
		t.Errorf("handler did not update title: got %v want %v",
			task.Title, updatedTask.Title)
	}

	if task.Completed != updatedTask.Completed {
		t.Errorf("handler did not update completed status: got %v want %v",
			task.Completed, updatedTask.Completed)
	}
}

func TestDeleteTask(t *testing.T) {
	createdTask := handlers.CreateTaskForTest(models.Task{Title: "Test Task", Description: "Test Description"})

	r := mux.NewRouter()
	r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	req, _ := http.NewRequest("DELETE", "/tasks/"+createdTask.ID, nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}

	// Verify the task is actually deleted
	r.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET")
	req, _ = http.NewRequest("GET", "/tasks/"+createdTask.ID, nil)
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("expected task to be deleted, but it was found")
	}
}

// Helper function to be added in handlers/handlers.go for testing purposes
// This is a workaround because we don't have a persistent database
func init() {
	// This is a dummy function to be replaced by a proper implementation
	// in the handlers package, to allow tests to create tasks.
}

// CreateTaskForTest creates a task for testing purposes and returns it.
// This function needs to be added to the handlers package.
// We are defining it here to show what is needed.
// In a real application, you would use a test database.

/*
// Add this to handlers/handlers.go
var (
	testNextID = 100 // Start test IDs from a different range
)
func CreateTaskForTest(task models.Task) models.Task {
	mu.Lock()
	defer mu.Unlock()

	task.ID = strconv.Itoa(testNextID)
	testNextID++
	task.CreatedAt = time.Now()
	tasks[task.ID] = task
	return task
}
*/
