package server

import (
	"net/http"
)

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is healthy"))
	})

	// 예: mux.HandleFunc("/api/users", usersHandler)
	// 예: mux.HandleFunc("/api/diaries", diariesHandler)
}
