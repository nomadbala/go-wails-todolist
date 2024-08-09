import React, { useState, useEffect } from 'react';
import {
    CreateTask,
    DeleteTask,
    GetAllTasks,
    MarkTaskDone,
    UpdateTask,
} from '../wailsjs/go/main/App.js';

const App = () => {
    const [tasks, setTasks] = useState([]);
    const [newTask, setNewTask] = useState('');
    const [newPriority, setNewPriority] = useState('low');
    const [newDeadline, setNewDeadline] = useState('');
    const [editTaskId, setEditTaskId] = useState(null);
    const [editTitle, setEditTitle] = useState('');
    const [editPriority, setEditPriority] = useState('low');
    const [editDeadline, setEditDeadline] = useState('');
    const [showDeleteModal, setShowDeleteModal] = useState(false);
    const [taskToDelete, setTaskToDelete] = useState(null);

    useEffect(() => {
        fetchTasks()
    }, []);

    const fetchTasks = async () => {
        try {
            const tasks = await GetAllTasks();
            console.log('Fetched tasks:', tasks);
            // const formattedTasks = tasks.map(task => ({
            //     ...task,
            //     deadline: task.deadline ? new Date(task.deadline).toISOString() : null
            // }));
            // console.log('Formatted tasks:', formattedTasks);
            setTasks(tasks);
        } catch (error) {
            console.error('Error fetching tasks:', error);
        }
    };

    const formatDeadline = (date) => {
        return date ? date.replace('T', ' ') : '';
    };

    const formatDeadline2 = (deadline) => {
        if (deadline.Valid) {
            const deadlineDate = new Date(deadline.Time);
            return deadlineDate.toLocaleString('en-GB', { timeZone: 'UTC' });
        } else {
            return 'No deadline';
        }
    };

    const handleAddTask = async () => {
        try {
            if (newTask === "") {
                return;
            }
            const formattedDeadline = formatDeadline(newDeadline);
            console.log('Adding task with deadline:', formattedDeadline);
            await CreateTask(newTask, false, newPriority, formattedDeadline);
            setNewTask('');
            setNewPriority('low');
            setNewDeadline('');
            fetchTasks();
        } catch (error) {
            console.error('Error creating task:', error);
            alert('Failed to create task: ' + error);
        }
    };

    const handleToggleTask = async (task) => {
        try {
            await MarkTaskDone(task.id);
            fetchTasks();
        } catch (error) {
            console.error('Error marking task as done:', error);
            alert('Failed to update task status: ' + error.message);
        }
    };

    const handleDeleteTask = async (task) => {
        try {
            await DeleteTask(task.id);
            setShowDeleteModal(false);
            setTaskToDelete(null);
            fetchTasks();
        } catch (error) {
            console.error('Error deleting task:', error);
            alert('Failed to delete task: ' + error.message);
        }
    };

    const handleUpdateTask = async () => {
        try {
            const formattedDeadline = formatDeadline(editDeadline);
            console.log('Updating task with deadline:', formattedDeadline);
            await UpdateTask(editTaskId, editTitle, editPriority, formattedDeadline);
            setEditTaskId(null);
            setEditTitle('');
            setEditPriority('low');
            setEditDeadline('');
            fetchTasks();
        } catch (error) {
            console.error('Error updating task:', error);
            alert('Failed to update task: ' + error.message);
        }
    };

    const sortedTasks = [...tasks].sort((a, b) => {
        const priorities = { 'low': 1, 'medium': 2, 'high': 3 };
        return priorities[b.priority] - priorities[a.priority];
    });

    return (
        <div className="bg-white shadow-md rounded-md p-6">
            <button onClick={fetchTasks}>qwe</button>
            <h2 className="text-2xl font-bold mb-4">Todo List</h2>
            <div className="flex flex-col gap-4 mb-4">
                <div className="flex gap-4 mb-4">
                    <input
                        type="text"
                        value={newTask}
                        onChange={(e) => setNewTask(e.target.value)}
                        className="flex-1 border-2 border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                        placeholder="Add a new task"
                    />
                    <select
                        value={newPriority}
                        onChange={(e) => setNewPriority(e.target.value)}
                        className="border-2 border-gray-300 rounded-md px-3 py-2"
                    >
                        <option value="low">Low</option>
                        <option value="medium">Medium</option>
                        <option value="high">High</option>
                    </select>
                    <input
                        type="datetime-local"
                        value={newDeadline}
                        onChange={(e) => setNewDeadline(e.target.value)}
                        className="border-2 border-gray-300 rounded-md px-3 py-2"
                    />
                    <button
                        onClick={handleAddTask}
                        className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-md"
                    >
                        Add
                    </button>
                </div>
                {sortedTasks.map((task) => (
                    <div key={task.id}
                         className={`flex items-center justify-between bg-gray-100 border-2 border-gray-300 rounded-md p-3 ${task.completed ? 'bg-green-100' : 'bg-red-100'}`}>
                        <div className="flex items-center gap-2">
                            <input
                                type="checkbox"
                                checked={task.completed}
                                onChange={() => handleToggleTask(task)}
                                className="form-checkbox h-5 w-5 text-blue-500"
                            />
                            {editTaskId === task.id ? (
                                <div className="flex gap-2">
                                    <input
                                        type="text"
                                        value={editTitle}
                                        onChange={(e) => setEditTitle(e.target.value)}
                                        className={`flex-1 font-medium ${task.completed ? 'line-through text-gray-500' : ''}`}
                                    />
                                    <select
                                        value={editPriority}
                                        onChange={(e) => setEditPriority(e.target.value)}
                                        className="border-2 border-gray-300 rounded-md px-3 py-2"
                                    >
                                        <option value="low">Low</option>
                                        <option value="medium">Medium</option>
                                        <option value="high">High</option>
                                    </select>
                                    <input
                                        type="datetime-local"
                                        value={editDeadline}
                                        onChange={(e) => setEditDeadline(e.target.value)}
                                        className="border-2 border-gray-300 rounded-md px-3 py-2"
                                    />
                                    <button
                                        onClick={handleUpdateTask}
                                        className="bg-green-500 hover:bg-green-600 text-white font-bold py-1 px-2 rounded-md"
                                    >
                                        Save
                                    </button>
                                    <button
                                        onClick={() => setEditTaskId(null)}
                                        className="bg-gray-500 hover:bg-gray-600 text-white font-bold py-1 px-2 rounded-md"
                                    >
                                        Cancel
                                    </button>
                                </div>
                            ) : (
                                <span
                                    className={`flex-1 font-medium ${task.completed ? 'line-through text-gray-500' : ''}`}>
                                    {task.title} - {task.priority} - {task.deadline ? formatDeadline2(task.deadline) : 'No deadline'}
                                </span>
                            )}
                        </div>
                        <div className="flex gap-2">
                            {editTaskId === task.id ? (
                                <button
                                    onClick={() => setEditTaskId(null)}
                                    className="bg-gray-500 hover:bg-gray-600 text-white font-bold py-1 px-2 rounded-md"
                                >
                                    Cancel
                                </button>
                            ) : (
                                <button
                                    onClick={() => {
                                        setEditTaskId(task.id);
                                        setEditTitle(task.title);
                                        setEditPriority(task.priority);
                                        setEditDeadline(task.deadline ? new Date(task.deadline).toISOString().slice(0, 16) : '');
                                    }}
                                    className="bg-yellow-500 hover:bg-yellow-600 text-white font-bold py-1 px-2 rounded-md"
                                >
                                    Edit
                                </button>
                            )}
                            <button
                                onClick={() => {
                                    setTaskToDelete(task);
                                    setShowDeleteModal(true);
                                }}
                                className="bg-red-500 hover:bg-red-600 text-white font-bold py-1 px-2 rounded-md"
                            >
                                Delete
                            </button>
                        </div>
                    </div>
                ))}
            </div>
            {showDeleteModal && (
                <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
                    <div className="bg-white p-4 rounded-md shadow-md">
                        <h3 className="text-lg font-bold mb-2">Confirm Delete</h3>
                        <p>Are you sure you want to delete this task?</p>
                        <div className="flex gap-2 mt-4">
                            <button
                                onClick={() => handleDeleteTask(taskToDelete)}
                                className="bg-red-500 hover:bg-red-600 text-white font-bold py-2 px-4 rounded-md"
                            >
                                Delete
                            </button>
                            <button
                                onClick={() => setShowDeleteModal(false)}
                                className="bg-gray-300 hover:bg-gray-400 text-black font-bold py-2 px-4 rounded-md"
                            >
                                Cancel
                            </button>
                        </div>
                    </div>
                </div>
            )}
        </div>
    );
};

export default App;