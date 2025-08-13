# Challenge: Task Manager CRUD + React Performance Fix

## Part 1: Golang CRUD API (7-8 minutes)

Instructions: "Use AI to help you create a simple CRUD API in Go for a Task management system"
Requirements:
```go
// Task model
type Task struct {
    ID          string    `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Completed   bool      `json:"completed"`
    CreatedAt   time.Time `json:"created_at"`
}

// Endpoints needed:
// GET    /tasks        - List all tasks (with optional ?completed=true/false filter)
// GET    /tasks/:id    - Get single task
// POST   /tasks        - Create task
// PUT    /tasks/:id    - Update task
// DELETE /tasks/:id    - Delete task
```

## Part 2: React Performance Fix (7-8 minutes)
Give them this problematic React component:

```javascript
// TaskList.jsx - This component has multiple performance issues
import React, { useState, useEffect } from 'react';

function TaskList() {
  const [tasks, setTasks] = useState([]);
  const [filter, setFilter] = useState('all');
  const [searchTerm, setSearchTerm] = useState('');

  useEffect(() => {
    fetch('http://localhost:8080/tasks')
      .then(res => res.json())
      .then(data => setTasks(data));
  });

  const deleteTask = (id) => {
    fetch(`http://localhost:8080/tasks/${id}`, { method: 'DELETE' })
      .then(() => {
        setTasks(tasks.filter(t => t.id !== id));
      });
  };

  const filteredTasks = tasks.filter(task => {
    console.log('Filtering...'); // This logs way too often
    const matchesFilter = filter === 'all' || 
      (filter === 'completed' && task.completed) ||
      (filter === 'pending' && !task.completed);
    
    const matchesSearch = task.title.toLowerCase()
      .includes(searchTerm.toLowerCase());
    
    return matchesFilter && matchesSearch;
  });

  return (
    <div>
      <input 
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
        placeholder="Search tasks..."
      />
      
      <select value={filter} onChange={(e) => setFilter(e.target.value)}>
        <option value="all">All</option>
        <option value="completed">Completed</option>
        <option value="pending">Pending</option>
      </select>

      {filteredTasks.map(task => (
        <TaskItem 
          task={task} 
          onDelete={deleteTask}
          // Problem 5: Creating new object as prop
          style={{ padding: '10px', margin: '5px' }}
        />
      ))}
    </div>
  );
}

function TaskItem({ task, onDelete, style }) {
  console.log('TaskItem rendering:', task.id);
  
  return (
    <div style={style}>
      <h3>{task.title}</h3>
      <p>{task.description}</p>
      <button onClick={() => onDelete(task.id)}>Delete</button>
    </div>
  );
}
``` 