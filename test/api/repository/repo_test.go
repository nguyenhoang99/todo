package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"go-chi-example/api/models"
	"log"
	"testing"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "nvh"
	password = ""
	dbname   = "nvh"
)

func TestImpl_CreateTodo(t *testing.T) {
	type args struct {
		ctx  context.Context
		todo *models.Todo
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "happy case",
			args: args{
				ctx: context.TODO(),
				todo: &models.Todo{
					Title:  "bac",
					Status: null.StringFrom("ninh"),
				},
			},
		},
		{
			name: " fail case",
			args: args{
				ctx: context.TODO(),
				todo: &models.Todo{
					ID:     1,
					Title:  "fail title",
					Status: null.StringFrom("fail status"),
				},
			},
			wantErr: errors.New("models: unable to insert into todos: pq: duplicate key value violates unique constraint \"todos_pkey\""),
		},
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	i := impl{
		dbConn: db,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := i.CreateTodo(tt.args.ctx, tt.args.todo)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("impl.Create todo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				row := db.QueryRow("SELECT id, title, status FROM todos ORDER BY created_at DESC LIMIT 1")
				var title string
				var status string
				var id int
				err = row.Scan(&id, &title, &status)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				require.Equal(t, tt.args.todo.Title, title)
				require.Equal(t, tt.args.todo.Status, null.StringFrom(status))
				_, err = db.Exec("DELETE FROM todos WHERE id = $1", id)
				if err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}

type Todo struct {
	ID        int
	Title     string
	Status    null.String
	CreatedAt time.Time
}

func TestImpl_GetTodoByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
		want    Todo
	}{
		{
			name: "get valid id",
			args: args{
				ctx: context.TODO(),
				id:  1,
			},
			want: Todo{
				ID:     1,
				Title:  "bac",
				Status: null.StringFrom("ninh"),
			},
		},
		{
			name: "get invalid id",
			args: args{
				ctx: context.TODO(),
				id:  100,
			},
			wantErr: sql.ErrNoRows,
		},
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	i := impl{
		dbConn: db,
	}
	ctx := context.Background()
	_, err = db.ExecContext(ctx, "INSERT INTO todos (id, title, status, created_at) VALUES ($1, $2, $3, $4)",
		1, "bac", "ninh", time.Now())
	if err != nil {
		t.Fatalf("error inserting test data: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			todoTest, err := i.GetTodoByID(tt.args.ctx, tt.args.id)
			if tt.wantErr != nil {
				require.NotNil(t, err)
				require.Equal(t, tt.wantErr.Error(), err.Error())
				return
			}
			require.Equal(t, tt.want.ID, todoTest.ID)
			require.Equal(t, tt.want.Title, todoTest.Title)
			require.Equal(t, tt.want.Status, todoTest.Status)
		})
	}
	row := db.QueryRow("SELECT id FROM todos ORDER BY created_at DESC LIMIT 1")
	var id int
	err = row.Scan(&id)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = db.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestImpl_GetAllTodo(t *testing.T) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	i := impl{
		dbConn: db,
	}
	ctx := context.Background()
	_, err = db.ExecContext(ctx, "INSERT INTO todos (id, title, status, created_at) VALUES ($1, $2, $3, $4)",
		1, "bac", "ninh", time.Now())
	if err != nil {
		t.Fatalf("error inserting test data: %v", err)
	}
	_, err = db.ExecContext(ctx, "INSERT INTO todos (id, title, status, created_at) VALUES ($1, $2, $3, $4)",
		2, "bac", "giang", time.Now())
	if err != nil {
		t.Fatalf("error inserting test data: %v", err)
	}
	todos, err := i.GetAllTodo(ctx)

	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	if len(todos) != 2 {
		t.Errorf("expected 2 todos, but got %d", len(todos))
	}

	if todos[0].ID != 1 || todos[0].Title != "bac" || todos[0].Status != null.StringFrom("ninh") {
		t.Errorf("expected todo 1 to be {ID: 1, Title: 'bac', Status: 'ninh'}, but got %v", todos[0])
	}
	if todos[1].ID != 2 || todos[1].Title != "bac" || todos[1].Status != null.StringFrom("giang") {
		t.Errorf("expected todo 2 to be {ID: 2, Title: 'bac', Status: 'giang'}, but got %v", todos[1])
	}
	_, err = db.Exec("DELETE FROM todos WHERE id IN (1, 2)")
	if err != nil {
		log.Fatal(err)
	}
}
