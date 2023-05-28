package repository

import (
	"context"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go-chi-example/api/models"
	"time"
)

type ModelTodo struct {
	ID        int
	Title     string
	Status    null.String
	CreatedAt time.Time
}

func (i impl) CreateTodo(ctx context.Context, t *models.Todo) error {
	if err := t.Insert(ctx, i.dbConn, boil.Infer()); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (i impl) GetTodoByID(ctx context.Context, id int) (ModelTodo, error) {
	t, err := models.Todos(models.TodoWhere.ID.EQ(id)).One(ctx, i.dbConn)
	if err != nil {
		return ModelTodo{}, err
	}
	return TodoModel(t), nil
}

func TodoModel(t *models.Todo) ModelTodo {
	todo := ModelTodo{
		ID:        t.ID,
		Title:     t.Title,
		Status:    t.Status,
		CreatedAt: t.CreatedAt,
	}
	return todo
}

func (i impl) GetAllTodo(ctx context.Context) ([]ModelTodo, error) {
	todo, err := models.Todos().All(ctx, i.dbConn)
	fmt.Println(todo)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	result := make([]ModelTodo, 0, len(todo))
	for _, a := range todo {
		result = append(result, TodoModel(a))
	}
	fmt.Print(result)
	return result, nil
}
