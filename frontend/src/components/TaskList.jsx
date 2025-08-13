// TaskList.jsx - This component has multiple performance issues
import React, { useState, useEffect, useMemo, useCallback } from 'react';

const taskItemStyle = { padding: '10px', margin: '5px' };

export function TaskList() {
  const [tasks, setTasks] = useState([]);
  const [filter, setFilter] = useState('all');
  const [searchTerm, setSearchTerm] = useState('');

  useEffect(() => {
    fetch('http://localhost:3001/tasks')
      .then(res => res.json())
      .then(data => setTasks(data));
  }, []);

  const deleteTask = useCallback((id) => {
    fetch(`http://localhost:8080/tasks/${id}`, { method: 'DELETE' })
      .then(() => {
        setTasks(prevTasks => prevTasks.filter(t => t.id !== id));
      });
  }, []);

  const filteredTasks = useMemo(() => {
    console.log('Filtering...');
    return tasks.filter(task => {
      const matchesFilter = filter === 'all' ||
        (filter === 'completed' && task.completed) ||
        (filter === 'pending' && !task.completed);

      const matchesSearch = task.title.toLowerCase()
        .includes(searchTerm.toLowerCase());

      return matchesFilter && matchesSearch;
    });
  }, [tasks, filter, searchTerm]);

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
          key={task.id}
          task={task} 
          onDelete={deleteTask}
          style={taskItemStyle}
        />
      ))}
    </div>
  );
}

export const TaskItem = React.memo(function TaskItem({ task, onDelete, style }) {
  console.log('TaskItem rendering:', task.id);
  
  return (
    <div style={style}>
      <h3>{task.title}</h3>
      <p>{task.description}</p>
      <button onClick={() => onDelete(task.id)}>Delete</button>
    </div>
  );
});
