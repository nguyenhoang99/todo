package controller

import (
	"context"
	"github.com/friendsofgo/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"go-chi-example/api/models"
	"go-chi-example/api/repository"
	"testing"
	"time"
)

func TestMockController_CreateTodo(t *testing.T) {
	mockTodoRepo := new(repository.MockRepository)
	type args struct {
		ctx  context.Context
		todo *models.Todo
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
		setup   func()
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
			setup: func() {
				mockTodoRepo.On("CreateTodo", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "fail case",
			args: args{
				ctx: context.TODO(),
				todo: &models.Todo{
					ID:     1,
					Title:  "bac",
					Status: null.StringFrom("giang"),
				},
			},
			wantErr: errors.New("error from repo"),
			setup: func() {
				mockTodoRepo.On("CreateTodo", mock.Anything, mock.Anything).Return(errors.New("error from repo"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := impl{mockTodoRepo}
			if tt.setup != nil {
				tt.setup()
			}
			todo, err := i.CreateTodo(tt.args.ctx, tt.args.todo.Title, tt.args.todo.Status)
			if err != nil && err != tt.wantErr {
				t.Errorf("impl.CreateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				require.Equal(t, tt.args.todo.Title, todo.Title)
				require.Equal(t, tt.args.todo.Status, todo.Status)
			}
		})
	}
}

func TestMockController_GetTodoByID(t *testing.T) {
	mockTodoRepo := new(repository.MockRepository)
	mockTodoDetail := repository.ModelTodo{
		ID:        1,
		Title:     "bac",
		Status:    null.StringFrom("ninh"),
		CreatedAt: time.Now(),
	}
	type args struct {
		ctx context.Context
		id  int
	}
	test := []struct {
		name    string
		args    args
		wantErr error
		want    TodoDetail
		setup   func()
	}{
		{
			name: "get valid id",
			args: args{
				ctx: context.TODO(),
				id:  1,
			},
			want: TodoDetail{
				ID:        1,
				Title:     "bac",
				Status:    null.StringFrom("ninh"),
				CreatedAt: time.Now(),
			},
			setup: func() {
				mockTodoRepo.On("GetTodoByID", mock.Anything, 1).Return(mockTodoDetail, nil)
			},
		},
		{
			name: "get invalid id",
			args: args{
				ctx: context.TODO(),
				id:  100,
			},
			want: TodoDetail{
				ID:        100,
				Title:     "bac",
				Status:    null.StringFrom("giang"),
				CreatedAt: time.Now(),
			},
			setup: func() {
				mockTodoRepo.On("GetTodoByID", mock.Anything, 100).Return(repository.ModelTodo{}, errors.New("failed to get todo"))
			},
			wantErr: errors.New("failed to get todo"),
		},
	}
	i := impl{mockTodoRepo}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			todo, err := i.GetTodoByID(tt.args.ctx, tt.args.id)
			if err != nil && tt.wantErr != nil {
				require.NotNil(t, err)
				require.Equal(t, tt.wantErr.Error(), err.Error())
				return
			}
			require.Equal(t, tt.want.ID, todo.ID)
			require.Equal(t, tt.want.Title, todo.Title)
			require.Equal(t, tt.want.Status, todo.Status)
		})
	}
}

func TestImpl_GetAllTodo(t *testing.T) {
	mockTodoRepo := new(repository.MockRepository)
	mockTodoList := []repository.ModelTodo{
		{ID: 1,
			Title:     "bac",
			Status:    null.StringFrom("ninh"),
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			Title:     "bac",
			Status:    null.StringFrom("giang"),
			CreatedAt: time.Now(),
		},
	}
	type args struct {
		ctx context.Context
	}
	test := []struct {
		name    string
		args    args
		want    []ModelTodo
		wantErr error
		setup   func()
	}{
		{
			name: "happy case",
			args: args{
				ctx: context.TODO(),
			},
			want: []ModelTodo{
				{ID: 1,
					Title:     "bac",
					Status:    null.StringFrom("ninh"),
					CreatedAt: time.Now(),
				},
				{
					ID:        2,
					Title:     "bac",
					Status:    null.StringFrom("giang"),
					CreatedAt: time.Now(),
				},
			},
			setup: func() {
				mockTodoRepo.On("GetAllTodo", mock.Anything).Return(mockTodoList, nil)
			},
		},
		{
			name: "fail case",
			args: args{
				ctx: context.TODO(),
			},
			want: []ModelTodo{
				{ID: 1,
					Title:     "bac",
					Status:    null.StringFrom("ninh"),
					CreatedAt: time.Now(),
				},
				{
					ID:        4,
					Title:     "bac",
					Status:    null.StringFrom("giang"),
					CreatedAt: time.Now(),
				},
			},
			setup: func() {
				mockTodoRepo.On("GetAllTodo", mock.Anything).Return(nil, errors.New("failed to get todo"))
			},
			wantErr: errors.New("failed to get todo"),
		},
	}

	i := impl{mockTodoRepo}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.setup != nil {
				tt.setup()
			}
			todoList, err := i.GetAllTodo(tt.args.ctx)
			if err != nil && tt.wantErr != nil {
				require.NotNil(t, err)
				require.Equal(t, tt.wantErr.Error(), err.Error())
				return
			}
			require.Equal(t, tt.want[0].ID, todoList[0].ID)
			require.Equal(t, tt.want[0].Title, todoList[0].Title)
			require.Equal(t, tt.want[0].Status, todoList[0].Status)
			require.Equal(t, tt.want[1].ID, todoList[1].ID)
			require.Equal(t, tt.want[1].Title, todoList[1].Title)
			require.Equal(t, tt.want[1].Status, todoList[1].Status)
		})
	}
}
