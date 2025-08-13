package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"task-api/models"
	"time"

	"github.com/gorilla/mux"
	gorillaHandlers "github.com/gorilla/handlers"
)

// In-memory data store
var (
	tasks  = make(map[string]models.Task)
	nextID = 1
	mu     sync.Mutex
)

// ListTasks handles GET /tasks
// It can filter tasks by the 'completed' query parameter.
// Example: GET /tasks?completed=true
func ListTasks(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	completedFilter := r.URL.Query().Get("completed")
	if completedFilter == "" {
		taskList := make([]models.Task, 0, len(tasks))
		for _, task := range tasks {
			taskList = append(taskList, task)
		}
		json.NewEncoder(w).Encode(taskList)
		return
	}

	completed, err := strconv.ParseBool(completedFilter)
	if err != nil {
		http.Error(w, "Invalid 'completed' query parameter", http.StatusBadRequest)
		return
	}

	filteredTasks := make([]models.Task, 0)
	for _, task := range tasks {
		if task.Completed == completed {
			filteredTasks = append(filteredTasks, task)
		}
	}
	json.NewEncoder(w).Encode(filteredTasks)
}

// GetTask handles GET /tasks/{id}
// It retrieves a single task by its ID.
func GetTask(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	vars := mux.Vars(r)
	id := vars["id"]

	task, ok := tasks[id]
	if !ok {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(task)
}

// CreateTask handles POST /tasks
// It creates a new task.
// Example: POST /tasks with JSON body: {"title": "New Task", "description": "A new task"}
func CreateTask(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task.ID = strconv.Itoa(nextID)
	nextID++
	task.CreatedAt = time.Now()
	tasks[task.ID] = task

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// UpdateTask handles PUT /tasks/{id}
// It updates an existing task.
// Example: PUT /tasks/1 with JSON body: {"title": "Updated Task", "completed": true}
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := tasks[id]; !ok {
		http.NotFound(w, r)
		return
	}

	var updatedTask models.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Preserve original creation date and ID
	originalTask := tasks[id]
	updatedTask.ID = id
	updatedTask.CreatedAt = originalTask.CreatedAt
	tasks[id] = updatedTask

	json.NewEncoder(w).Encode(updatedTask)
}

// DeleteTask handles DELETE /tasks/{id}
// It deletes a task by its ID.
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := tasks[id]; !ok {
		http.NotFound(w, r)
		return
	}

	delete(tasks, id)
	w.WriteHeader(http.StatusNoContent)
}

// CreateTaskForTest is a helper function for testing purposes.
// It creates a task and adds it to the in-memory store.
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


func CreateCORSHandler() func(http.Handler) http.Handler {

	// --- CORS Configuration ---
	// Define the allowed origins. You can use "*" for public APIs,
	// but it's better to be specific for security.
	allowedOrigins := gorillaHandlers.AllowedOrigins([]string{"*"})

	// Define the allowed HTTP methods
	allowedMethods := gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	// Define the allowed headers
	allowedHeaders := gorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	
	// Optional: Allow credentials, such as cookies, to be sent
	allowCredentials := gorillaHandlers.AllowCredentials()

	// Create the CORS handler with our options
	return gorillaHandlers.CORS(allowedOrigins, allowedMethods, allowedHeaders, allowCredentials)
}
