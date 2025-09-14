import { Box, Button, TextField } from "@mui/material";
import { useState } from "react";

export const CreateForm = ({ handleAddTask }) => {
  const [title, setTitle] = useState("")
  const [description, setDescription] = useState("")

  const handleFormSubmit = (e) => {
    e.preventDefault()

    handleAddTask({
      title: title,
      description: description,
    })

    setTitle("")
    setDescription("")
  }

  return (
    <Box
      component="form"
      noValidate
      autoComplete="off"
      sx={{
        display: 'grid',
        justifyContent: 'center',
        alignContent: 'center',
        height: '20vh',
      }}
    >
      <TextField
        id="title"
        placeholder="Type a title..."
        variant="outlined"
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        autoFocus

        sx={{
          width: '300px',
          padding: '5px',
        }}
      />

      <TextField
        type="text"
        id="description"
        placeholder="Type a description..."
        variant="outlined"
        multiline
        maxRows={2}
        value={description}
        onChange={(e) => setDescription(e.target.value)}

        sx={{
          width: '300px',
          padding: '5px',
        }}
      />

      <Button
        variant="contained"
        onClick={handleFormSubmit}
      >
        Add
      </Button>
    </Box>
  )
}
