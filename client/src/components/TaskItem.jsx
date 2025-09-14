import { Fragment, useState } from "react";
import {
  Checkbox,
  IconButton,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Typography
} from "@mui/material"
import DeleteOutlinedIcon from '@mui/icons-material/DeleteOutlined';
import ModeEditOutlineIcon from '@mui/icons-material/ModeEditOutline';
import formatData from '../util'

export const TaskItem = ({ task, labelId, handleToggleTask, enterEditMode, handleDeleteTask }) => {
  const [isDone, setIsDone] = useState(task.done)

  const handleCheckBoxChange = () => {
    setIsDone(!isDone)
    handleToggleTask({ ...task, done: !isDone })
  }

  return (
    <ListItem
      divider
      secondaryAction={
        <Fragment>
          <IconButton
            edge="start"
            aria-label="edit"
            onClick={() => enterEditMode(task)}
          >
            <ModeEditOutlineIcon />
          </IconButton>
          <IconButton
            edge="end" aria-label="delete"
            onClick={() => handleDeleteTask(task.id)}>
            <DeleteOutlinedIcon />
          </IconButton>
        </Fragment>
      }
    >
      <ListItemButton
        onClick={handleCheckBoxChange}
      >
        <ListItemIcon>
          <Checkbox
            edge="start"
            checked={isDone}
            color="success"
            inputProps={{ 'aria-labelledby': labelId }}
          />
        </ListItemIcon>

        <ListItemText
          id={labelId}
          primary={
            task.title
          }
          secondary={
            <Fragment>
              <Typography
                component="span"
                variant="body2"
                color="text.primary"
                sx={{
                  display: 'inline'
                }}
              >
                {task.description}
                <br />
              </Typography>
              {`created ${formatData(task.created_at)}`}
            </Fragment>
          }

        />

      </ListItemButton>
    </ListItem>
  )
}
