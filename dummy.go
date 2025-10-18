package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

var (
	todos = []Todo{}
	mu    sync.Mutex
	id    int
)

func main() {
	http.HandleFunc("/todos", handleTodos)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mu.Lock()
	defer mu.Unlock()

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(todos)
	case http.MethodPost:
		var t Todo
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		id++
		t.ID = id
		todos = append(todos, t)
		json.NewEncoder(w).Encode(t)
	default:
		http.Error(w, "Method not allowed", 405)
	}
}
