package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/petrick-ribeiro/devops-pucpr/types"
)

type mockStorage struct {
	todos []*types.Todo
	err   error
}

func (m *mockStorage) GetAll() ([]*types.Todo, error) {
	return m.todos, m.err
}
func (m *mockStorage) Get(id uint64) (*types.Todo, error) {
	if m.err != nil {
		return nil, m.err
	}
	return &types.Todo{ID: uint(id), Title: "Test", Description: "Desc"}, nil
}
func (m *mockStorage) Insert(todo *types.Todo) error {
	return m.err
}
func (m *mockStorage) Update(todo *types.Todo, id uint64) (*types.Todo, error) {
	if m.err != nil {
		return nil, m.err
	}
	return todo, nil
}
func (m *mockStorage) Delete(id uint64) (*types.Todo, error) {
	if m.err != nil {
		return nil, m.err
	}
	return &types.Todo{ID: uint(id), Title: "Deleted"}, nil
}

func TestHandleGetTodo(t *testing.T) {
	s := &APIServer{storage: &mockStorage{todos: []*types.Todo{{ID: 1, Title: "Test"}}}}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/todo", nil)
	err := s.handleGetTodo(w, r)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestHandleGetTodo_Error(t *testing.T) {
	s := &APIServer{storage: &mockStorage{err: errors.New("fail")}}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/todo", nil)
	err := s.handleGetTodo(w, r)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestHandlePostTodo(t *testing.T) {
	s := &APIServer{storage: &mockStorage{}}
	reqBody := types.CreateTodoRequest{Title: "New", Description: "Desc"}
	b, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/todo", bytes.NewReader(b))
	err := s.handlePostTodo(w, r)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestHandlePostTodo_BadJSON(t *testing.T) {
	s := &APIServer{storage: &mockStorage{}}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/todo", bytes.NewReader([]byte("bad json")))
	err := s.handlePostTodo(w, r)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestHandleGetTodoByID(t *testing.T) {
	s := &APIServer{storage: &mockStorage{}}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/todo/1", nil)
	r = muxSetID(r, "1")
	err := s.handleGetTodoByID(w, r)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestHandleGetTodoByID_BadID(t *testing.T) {
	s := &APIServer{storage: &mockStorage{}}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/todo/abc", nil)
	r = muxSetID(r, "abc")
	err := s.handleGetTodoByID(w, r)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestHandlePutTodoByID(t *testing.T) {
	s := &APIServer{storage: &mockStorage{}}
	reqBody := types.UpdateTodoRequest{Title: "Up", Description: "Desc", Done: true}
	b, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/todo/1", bytes.NewReader(b))
	r = muxSetID(r, "1")
	err := s.handlePutTodoByID(w, r)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestHandlePutTodoByID_BadID(t *testing.T) {
	s := &APIServer{storage: &mockStorage{}}
	reqBody := types.UpdateTodoRequest{Title: "Up", Description: "Desc", Done: true}
	b, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/todo/abc", bytes.NewReader(b))
	r = muxSetID(r, "abc")
	err := s.handlePutTodoByID(w, r)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestHandleDeleteByID(t *testing.T) {
	s := &APIServer{storage: &mockStorage{}}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/todo/1", nil)
	r = muxSetID(r, "1")
	err := s.handleDeleteByID(w, r)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestHandleDeleteByID_BadID(t *testing.T) {
	s := &APIServer{storage: &mockStorage{}}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/todo/abc", nil)
	r = muxSetID(r, "abc")
	err := s.handleDeleteByID(w, r)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func muxSetID(r *http.Request, id string) *http.Request {
	return mux.SetURLVars(r, map[string]string{"id": id})
}
