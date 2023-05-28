package controller

import (
	"context"
	"github.com/volatiletech/null/v8"
	"go-chi-example/api/models"
	"go-chi-example/api/repository"
	"time"
)

type ModelTodo struct {
	ID        int         `json:"id"`
	Title     string      `json:"title"`
	Status    null.String `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
}

type TodoDetail struct {
	ID        int         `json:"id"`
	Title     string      `json:"title"`
	Status    null.String `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
}

type impl struct {
	repository repository.Repository
}

func (i impl) CreateTodo(ctx context.Context, title string, status null.String) (*models.Todo, error) {
	todo := &models.Todo{
		Title:  title,
		Status: status,
	}
	if err := i.repository.CreateTodo(ctx, todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (i impl) GetTodoByID(ctx context.Context, id int) (*TodoDetail, error) {
	todo, err := i.repository.GetTodoByID(ctx, id)
	if err != nil {
		return nil, err
	}
	todoDetail := TodoDetail{
		ID:        todo.ID,
		Title:     todo.Title,
		Status:    todo.Status,
		CreatedAt: todo.CreatedAt,
	}
	return &todoDetail, err
}

func TodoModel(t ModelTodo) ModelTodo {
	todo := ModelTodo{
		ID:        t.ID,
		Title:     t.Title,
		Status:    t.Status,
		CreatedAt: t.CreatedAt,
	}
	return todo
}

func (i impl) GetAllTodo(ctx context.Context) ([]ModelTodo, error) {
	todo, err := i.repository.GetAllTodo(ctx)
	if err != nil {
		return nil, err
	}
	todoList := make([]ModelTodo, 0, len(todo))
	for _, value := range todo {
		todoList = append(todoList, TodoModel(ModelTodo(value)))
	}
	return todoList, err
}
