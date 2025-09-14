// component import
import { List } from "@mui/material"
import { TaskItem } from "./TaskItem"

export const TaskList = ({ tasks, handleToggleTask, enterEditMode, handleDeleteTask }) => {
  return (
    <List
      sx={{
        display: 'grid',
        justifyContent: 'center',
        alignContent: 'center',
        width: '100%',
        bgcolor: 'background.paper'
      }}
    >
      {
        tasks.sort((a, b) => b.id - a.id).map((task) => {
          const labelId = `checkbox-list-label-${task.id}`
          return (
            <TaskItem
              key={task.id}
              task={task}
              labelId={labelId}
              handleToggleTask={handleToggleTask}
              enterEditMode={enterEditMode}
              handleDeleteTask={handleDeleteTask}
            />
          )
        })
      }
    </List>
  )

}
