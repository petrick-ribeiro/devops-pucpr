package api

import (
	"fmt"
	"log"
	"net/http"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/petrick-ribeiro/devops-pucpr/storage"
)

type APIServer struct {
	listenAddr string
	storage    storage.Storage
}

func NewAPIServer(listenAddr string, db storage.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		storage:    db,
	}
}

func (s *APIServer) Run() error {
	r := mux.NewRouter()

	// Handlers
	r.HandleFunc("/todo", makeHTTPHandleFunc(s.handleTodo))
	r.HandleFunc("/todo/{id}", makeHTTPHandleFunc(s.handleTodoByID))

	// Set CORS
	headers := gohandlers.AllowedHeaders([]string{
		"X-Requested-With",
		"Content-Type",
		"Authorization",
	})
	methods := gohandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := gohandlers.AllowedOrigins([]string{"*"})
	ch := gohandlers.CORS(headers, methods, origins)

	// Set Server
	srv := http.Server{
		Handler: ch(r),
		Addr:    fmt.Sprintf("0.0.0.0:%s", s.listenAddr),
	}

	log.Printf("Starting the server on port %s", s.listenAddr)

	// Start the Server
	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
