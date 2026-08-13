package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Vxplain/sample-repo/internal/tasks"
)

func main() {
	store := tasks.NewStore()
	_, _ = store.Add("open a pull request")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, store.List())
	})
	mux.HandleFunc("POST /tasks", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Title string `json:"title"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		task, err := store.Add(input.Title)
		if errors.Is(err, tasks.ErrEmptyTitle) {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		if err != nil {
			http.Error(w, "could not create task", http.StatusInternalServerError)
			return
		}
		writeJSON(w, http.StatusCreated, task)
	})

	log.Println("sample service listening on http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", mux))
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Printf("encode response: %v", err)
	}
}
