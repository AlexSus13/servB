package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
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
		Addr:         ":8080",
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server Start")

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("1ERROR: ", err)
	}

	fmt.Println("Server Stop")

}

func (h *Handler) Test(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)
	w.Write([]byte(h.time))
}
