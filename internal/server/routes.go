package server

import (
	"net/http"

	"github.com/jhphon0730/dairify/internal/database"
)

func SetupRoutes(mux *http.ServeMux, db *database.DB) {
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is healthy"))
	})

	// 예: mux.HandleFunc("/api/users", usersHandler)
	// 예: mux.HandleFunc("/api/diaries", diariesHandler)
}
