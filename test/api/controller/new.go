package controller

import (
	"context"
	"github.com/volatiletech/null/v8"
	"go-chi-example/api/models"
	"go-chi-example/api/repository"
)

type Controller interface {
	CreateTodo(ctx context.Context, title string, status null.String) (*models.Todo, error)
	GetTodoByID(ctx context.Context, id int) (*TodoDetail, error)
	GetAllTodo(ctx context.Context) ([]ModelTodo, error)
}

func New(repo repository.Repository) Controller {
	return impl{
		repository: repo,
	}
}
