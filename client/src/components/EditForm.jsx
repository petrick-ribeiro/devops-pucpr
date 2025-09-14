import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  TextField
} from "@mui/material";
import { useState } from "react";

export const EditForm = ({ isEditing, setIsEditing, editedTask, handleUpdateTask, closeEditMode }) => {
  const [updatedTitle, setUpdatedTitle] = useState(editedTask.title)
  const [updateDescription, setUpdatedDescription] = useState(editedTask.description)

  const handleFormSubmit = (e) => {
    e.preventDefault()
    handleUpdateTask({
      ...editedTask,
      title: updatedTitle,
      description: updateDescription
    })
    setIsEditing(false)
  }

  return (
    <Dialog open={isEditing} onClose={closeEditMode}>
      <DialogTitle>Edit Task</DialogTitle>
      <DialogContent
        sx={{
          display: "grid",
          justifyContent: 'center',
          alignContent: 'center',
          width: '310px'
        }}
      >
        <TextField
          id="title"
          variant="outlined"
          value={updatedTitle}
          onChange={(e) => setUpdatedTitle(e.target.value)}
          autoFocus
          sx={{
            width: '300px',
            padding: '5px'
          }}
        />

        <TextField
          type="text"
          id="description"
          variant="outlined"
          multiline
          maxRows={4}
          value={updateDescription}
          onChange={(e) => setUpdatedDescription(e.target.value)}
          sx={{
            width: '300px',
            padding: '5px'
          }}
        />
      </DialogContent>

      <DialogActions>
        <Button variant="contained" onClick={handleFormSubmit}>Update</Button>
      </DialogActions>
    </Dialog>
  )
}
