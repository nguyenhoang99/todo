package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/friendsofgo/errors"
	"github.com/go-chi/chi/v5"
	"github.com/volatiletech/null/v8"
	"io"
	"net/http"
	"strconv"
)

const (
	todoIDURLParam = "id"
)

type CreateTodoRequest struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

type TodoRequest struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func (h *Handler) CreateHandlerTodo(w http.ResponseWriter, r *http.Request) {
	data := CreateTodoRequest{}
	reqBytes, err := io.ReadAll(r.Body)
	if err = json.Unmarshal(reqBytes, &data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(getHttpStatus(err))
		return
	}
	defer r.Body.Close()
	todo, err := h.todoCtrl.CreateTodo(r.Context(), data.Title, null.StringFrom(data.Status))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(getHttpStatus(err))
		return
	}
	todoJson, err := json.Marshal(todo)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(todoJson)
}

func (h *Handler) GetHandlerTodoByID(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, todoIDURLParam)
	id, err := strconv.Atoi(todoID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	todo, err := h.todoCtrl.GetTodoByID(r.Context(), id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(getHttpStatus(err))
		return
	}
	todoJson, err := json.Marshal(todo)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(getHttpStatus(err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(todoJson)
}

func (h *Handler) GetAllTodo(w http.ResponseWriter, r *http.Request) {
	todo, err := h.todoCtrl.GetAllTodo(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(getHttpStatus(err))
		return
	}
	todoJson, err := json.Marshal(todo)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(getHttpStatus(err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(todoJson)
}

func getHttpStatus(err error) int {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
