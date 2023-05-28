package handler

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"go-chi-example/api/controller"
	"go-chi-example/api/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

type todoResponse struct {
	ID     int         `json:"id"`
	Title  string      `json:"title"`
	Status null.String `json:"status"`
}

type CreateTodoRequestTest struct {
	Title  string      `json:"title"`
	Status null.String `json:"status"`
}

func TestMockHandler_CreateHandlerTodo(t *testing.T) {
	mockTodoController := new(controller.MockController)
	validRqData := &CreateTodoRequestTest{
		Title:  "bac",
		Status: null.StringFrom("ninh"),
	}
	validRqDataBlob, _ := json.Marshal(validRqData)
	validRq := httptest.NewRequest("GET", "/todos/", bytes.NewReader(validRqDataBlob))
	errData := &CreateTodoRequestTest{
		Title:  "bac",
		Status: null.StringFrom("giang"),
	}
	errRqDataBlob, _ := json.Marshal(errData)
	errRq := httptest.NewRequest("GET", "/todos/", bytes.NewReader(errRqDataBlob))
	type fields struct {
		todoCtrl controller.Controller
	}
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantErr      error
		wantResponse *todoResponse
		setup        func()
	}{
		{
			name: "happy case",
			fields: fields{
				todoCtrl: mockTodoController,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: validRq,
			},
			setup: func() {
				mockTodoController.On("CreateTodo", validRq.Context(), validRqData.Title, validRqData.Status).Return(&models.Todo{ID: 1, Title: "bac", Status: null.StringFrom("ninh")}, nil)
			},
			wantResponse: &todoResponse{ID: 1, Title: "bac", Status: null.StringFrom("ninh")},
		},
		{
			name: "fail case",
			fields: fields{
				todoCtrl: mockTodoController,
			},
			args: args{
				w: httptest.NewRecorder(),
				r: errRq,
			},
			setup: func() {
				mockTodoController.On("CreateTodo", errRq.Context(), errData.Title, errData.Status).Return(nil, sql.ErrNoRows)
			},
			wantErr: errors.New(" error from controller "),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			h := &Handler{
				todoCtrl: tt.fields.todoCtrl,
			}
			h.CreateHandlerTodo(tt.args.w, tt.args.r)
			if tt.wantResponse != nil {
				res := &todoResponse{}
				err := json.Unmarshal(tt.args.w.Body.Bytes(), res)
				if err != nil {
					fmt.Println(err)
					t.Error("Invalid response format")
				} else if res.Title != tt.wantResponse.Title || res.Status != tt.wantResponse.Status {
					t.Error("Invalid created todo")
				}
			}
		})
	}
}

func TestHandler_GetHandlerTodoByID(t *testing.T) {
	t.Parallel()
	mockTodoController := new(controller.MockController)
	h := Handler{
		todoCtrl: mockTodoController,
	}
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody *models.Todo
		wantErr  error
		setup    func() *chi.Context
	}{
		{
			name: "Get valid id",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/todos/1", nil),
			},
			wantCode: http.StatusOK,
			wantBody: &models.Todo{
				ID:     1,
				Title:  "test",
				Status: null.StringFrom("todo"),
			},
			setup: func() *chi.Context {
				rctx := chi.NewRouteContext()
				rctx.URLParams.Add("id", "1")
				mockTodoController.On("GetTodoByID", mock.Anything, 1).Return(&controller.TodoDetail{ID: 1, Title: "test", Status: null.StringFrom("todo")}, nil)
				return rctx
			},
		},
		{
			name: "not found todo",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/todos/100", nil),
			},
			wantCode: http.StatusNotFound,
			wantErr:  errors.New("error from controller"),
			setup: func() *chi.Context {
				rctx := chi.NewRouteContext()
				rctx.URLParams.Add("id", "100")
				mockTodoController.On("GetTodoByID", mock.Anything, 100).Return(nil, sql.ErrNoRows)
				return rctx
			},
		},
		{
			name: "invalid todo id",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/todos/abc", nil),
			},
			wantCode: http.StatusBadRequest,
			wantErr:  errors.New("error from controller"),
			setup: func() *chi.Context {
				rctx := chi.NewRouteContext()
				rctx.URLParams.Add("id", "abc")
				return rctx
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				rctx := tt.setup()
				tt.args.r = tt.args.r.WithContext(context.WithValue(tt.args.r.Context(), chi.RouteCtxKey, rctx))
			}
			h.GetHandlerTodoByID(tt.args.w, tt.args.r)
			require.Equal(t, tt.args.w.Code, tt.wantCode)
			if tt.args.w.Code == http.StatusOK {
				var gotTodo *models.Todo
				err := json.Unmarshal(tt.args.w.Body.Bytes(), &gotTodo)
				assert.NoError(t, err)
				require.Equal(t, tt.wantBody, gotTodo)
			}

		})
	}
}

func TestHandler_GetAllTodo(t *testing.T) {
	mockTodoController := new(controller.MockController)
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantErr  error
		setup    func()
	}{
		{
			name: "happy case",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/todos/", nil),
			},
			wantCode: http.StatusOK,
			setup: func() {
				mockTodoController.On("GetAllTodo", mock.Anything).Return([]controller.ModelTodo{}, nil)
			},
		},
		{
			name: "json parse error",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/todos/", nil),
			},
			wantCode: http.StatusBadRequest,
			wantErr:  sql.ErrNoRows,
		},
		{
			name: "sql error",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/todos/", nil),
			},
			wantCode: http.StatusNotFound,
			wantErr:  sql.ErrNoRows,
			setup: func() {
				mockTodoController.On("GetAllTodo", mock.Anything).Return(nil, sql.ErrNoRows)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			h := Handler{
				todoCtrl: mockTodoController,
			}
			h.GetAllTodo(tt.args.w, tt.args.r)
			require.Equal(t, tt.args.w.Code, tt.wantCode)
		})
	}
}
