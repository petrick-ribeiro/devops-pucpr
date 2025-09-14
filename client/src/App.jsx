import { useEffect, useState } from 'react'
import './App.css'
// custom components
import { CreateForm } from './components/CreateForm'
import { EditForm } from './components/EditForm'
import { TaskList } from './components/TaskList'
// API functions
import {
  addTask,
  deleteTask,
  getTasks,
  updateTask,
  updateTaskDone
} from './services/api'

function App() {
  const [tasks, setTasks] = useState([])
  const [editedTask, setEditedTask] = useState(null)
  const [isEditing, setIsEditing] = useState(false)

  useEffect(() => {
    async function fetchData() {
      try {
        const data = await getTasks()
        setTasks(data)
      } catch (error) {
        console.error(error)
      }
    } fetchData()
  }, [])

  const handleAddTask = async (task) => {
    try {
      const newTask = await addTask(task)
      setTasks(prevState => [...prevState, newTask])
    } catch (error) {
      console.error(error)
    }
  }

  const handleToggleTask = async (task) => {
    try {
      const data = await updateTaskDone(task)
      setTasks(prevState => prevState.map(t => (
        t.id === task.id
          ? data
          : t
      )))
    } catch (error) {
      console.error(error)
    }
  }

  const enterEditMode = (task) => {
    setEditedTask(task)
    setIsEditing(true)
  }

  const closeEditMode = () => {
    setIsEditing(false)
  }

  const handleUpdateTask = async (task) => {
    try {
      const data = await updateTask(task)
      setTasks(prevState => prevState.map(t => (
        t.id === task.id
          ? data
          : t
      )))

      closeEditMode()
    } catch (error) {
      console.error(error)
    }
  }

  const handleDeleteTask = async (id) => {
    try {
      await deleteTask(id)
      setTasks(prevState => prevState.filter(task => task.id !== id))
    } catch (error) {
      console.error(error)
    }
  }

  return (
    <div className='app'>
      {
        isEditing && (
          <EditForm
            isEditing={isEditing}
            setIsEditing={setIsEditing}
            editedTask={editedTask}
            handleUpdateTask={handleUpdateTask}
            closeEditMode={closeEditMode}
          />
        )
      }

      <CreateForm handleAddTask={handleAddTask} />

      {
        tasks.length > 0 &&
        <TaskList
          tasks={tasks}
          handleToggleTask={handleToggleTask}
          enterEditMode={enterEditMode}
          handleDeleteTask={handleDeleteTask}
        />
      }
    </div>
  )
}

export default App
