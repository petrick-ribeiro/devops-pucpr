package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/petrick-ribeiro/devops-pucpr/types"
)

func (s *APIServer) handleTodo(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetTodo(w, r)
	}

	if r.Method == "POST" {
		return s.handlePostTodo(w, r)
	}

	return fmt.Errorf("method %s not supported", r.Method)
}

func (s *APIServer) handleTodoByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetTodoByID(w, r)
	}

	if r.Method == "PUT" {
		return s.handlePutTodoByID(w, r)
	}

	if r.Method == "DELETE" {
		return s.handleDeleteByID(w, r)
	}

	return fmt.Errorf("method %s not supported", r.Method)
}

func (s *APIServer) handleGetTodo(w http.ResponseWriter, r *http.Request) error {
	todos, err := s.storage.GetAll()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, todos)
}

func (s *APIServer) handleGetTodoByID(w http.ResponseWriter, r *http.Request) error {
	id, err := GetID(r)
	if err != nil {
		return err
	}

	todo, err := s.storage.Get(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, todo)
}

func (s *APIServer) handlePostTodo(w http.ResponseWriter, r *http.Request) error {
	req := new(types.CreateTodoRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	todo := types.NewTodo(req.Title, req.Description)

	if err := s.storage.Insert(todo); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, todo)
}

func (s *APIServer) handlePutTodoByID(w http.ResponseWriter, r *http.Request) error {
	id, err := GetID(r)
	if err != nil {
		return err
	}

	req := new(types.UpdateTodoRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	todo := types.UpdateTodo(req.Title, req.Description, req.Done)

	updatedTodo, err := s.storage.Update(todo, id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, updatedTodo)
}

func (s *APIServer) handleDeleteByID(w http.ResponseWriter, r *http.Request) error {
	id, err := GetID(r)
	if err != nil {
		return err
	}

	todo, err := s.storage.Delete(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, todo)
}
