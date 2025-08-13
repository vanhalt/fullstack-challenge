// TaskList.jsx - This component has multiple performance issues
import React, { useState, useEffect } from 'react';

export function TaskList() {
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

export function TaskItem({ task, onDelete, style }) {
  console.log('TaskItem rendering:', task.id);
  
  return (
    <div style={style}>
      <h3>{task.title}</h3>
      <p>{task.description}</p>
      <button onClick={() => onDelete(task.id)}>Delete</button>
    </div>
  );
}
