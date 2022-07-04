package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"fmt"
	"os"
)

type Handler struct {
	time string
}

func main() {

	t := time.Now()

	ts := t.String()

	hand := &Handler{time: ts}

	router := mux.NewRouter()

	router.HandleFunc("/test", hand.Test).Methods("GET")

	srv := &http.Server{
		Addr:         os.Getenv("HOST") + ":8080",
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server Start")

	srv.ListenAndServe()

	fmt.Println("Server Stop")

}

func (h *Handler) Test(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)
	w.Write([]byte(h.time))
}
