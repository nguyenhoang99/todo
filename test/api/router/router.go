package router

import (
	"context"
	"github.com/go-chi/chi/v5"
	"go-chi-example/api/handler"
)

type Router struct {
	Ctx         context.Context
	TodoHandler handler.Handler
}

func (rtr Router) HandlerTodo() chi.Router {
	r := chi.NewRouter()
	r.Post("/create-todo", rtr.TodoHandler.CreateHandlerTodo)
	r.Get("/{id}", rtr.TodoHandler.GetHandlerTodoByID)
	r.Get("/", rtr.TodoHandler.GetAllTodo)
	return r
}
