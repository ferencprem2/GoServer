package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SendResponse(w http.ResponseWriter, i any) {
	data, err := json.Marshal(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func DecodeRequest(w http.ResponseWriter, r *http.Request, i any) bool {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}
	return true
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Middleware", r.URL)
			next.ServeHTTP(w, r)

		})
}

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/todos", Controller_Todos).Methods("GET", "POST")
	mux.HandleFunc("/todos/{id:[0-9]+}", Controller_Todos_Id).Methods("GET", "PATCH", "DELETE")

	//mux.Use(Middleware)

	http.ListenAndServe(":7777", mux)
}
