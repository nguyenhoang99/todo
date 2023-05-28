package repository

import (
	"context"
	"database/sql"
	"go-chi-example/api/models"
)

type Repository interface {
	CreateTodo(ctx context.Context, t *models.Todo) error
	GetTodoByID(ctx context.Context, id int) (ModelTodo, error)
	GetAllTodo(ctx context.Context) ([]ModelTodo, error)
}

type impl struct {
	dbConn *sql.DB
}

func New(dbConn *sql.DB) Repository {
	return impl{
		dbConn: dbConn,
	}
}
