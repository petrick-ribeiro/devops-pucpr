import axios from "axios";

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
  timeout: 5000,
});

export const getTasks = async () => {
  const response = await api.get("/todo");

  if (response.status === 200) {
    return response.data;
  }

  throw new Error("Failed to get the tasks");
};

export const addTask = async (task) => {
  const response = await api.post("/todo", {
    title: task.title,
    description: task.description,
  });

  if (response.status === 200) {
    return response.data;
  }

  throw new Error("Failed to add the task");
};

export const updateTaskDone = async (task) => {
  const response = await api.put(`/todo/${task.id}`, { done: task.done });

  if (response.status === 200) {
    return response.data;
  }

  throw new Error("Failed to toggle task");
};

export const updateTask = async (task) => {
  const response = await api.put(`/todo/${task.id}`, {
    title: task.title,
    description: task.description,
    done: task.done,
  });

  if (response.status === 200) {
    return response.data;
  }

  throw new Error("Failed to update the task");
};

export const deleteTask = async (id) => {
  const response = await api.delete(`/todo/${id}`);

  if (response.status === 200) {
    return true;
  }

  throw new Error("Failed to delete the task");
};
