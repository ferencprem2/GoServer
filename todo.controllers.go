package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TodoEditor struct {
	Text string
	Done bool
}

type TodoItem struct {
	TodoId int
	Text   string
	Done   bool
}

func Controller_Todos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		result, err := GetTodoItems()
		if err != nil {
			http.Error(w, "invalid id", http.StatusInternalServerError)
			return
		}
		SendResponse(w, result)
	case http.MethodPost:
		var request TodoEditor
		if !DecodeRequest(w, r, &request) {
			return
		}
		err := CreateTodoItem(request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		SendResponse(w, struct{}{})

	}
}
func Controller_Todos_Id(w http.ResponseWriter, r *http.Request) {
	var id int
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		result, err := GetTodoItem(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, result)

	case http.MethodPatch:
		var request TodoEditor
		if !DecodeRequest(w, r, &request) {
			return
		}
		err := UpdateTodoItem(request, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		SendResponse(w, struct{}{})

	case http.MethodDelete:
		err := DeleteTodoItem(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		SendResponse(w, struct{}{})
	}
}
