// Code generated by mockery v2.20.0. DO NOT EDIT.

package controller

import (
	context "context"
	models "go-chi-example/api/models"

	mock "github.com/stretchr/testify/mock"

	null "github.com/volatiletech/null/v8"
)

// MockController is an autogenerated mock type for the Controller type
type MockController struct {
	mock.Mock
}

// CreateTodo provides a mock function with given fields: ctx, title, status
func (_m *MockController) CreateTodo(ctx context.Context, title string, status null.String) (*models.Todo, error) {
	ret := _m.Called(ctx, title, status)

	var r0 *models.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, null.String) (*models.Todo, error)); ok {
		return rf(ctx, title, status)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, null.String) *models.Todo); ok {
		r0 = rf(ctx, title, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Todo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, null.String) error); ok {
		r1 = rf(ctx, title, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllTodo provides a mock function with given fields: ctx
func (_m *MockController) GetAllTodo(ctx context.Context) ([]ModelTodo, error) {
	ret := _m.Called(ctx)

	var r0 []ModelTodo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]ModelTodo, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []ModelTodo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ModelTodo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTodoByID provides a mock function with given fields: ctx, id
func (_m *MockController) GetTodoByID(ctx context.Context, id int) (*TodoDetail, error) {
	ret := _m.Called(ctx, id)

	var r0 *TodoDetail
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*TodoDetail, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *TodoDetail); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TodoDetail)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockController interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockController creates a new instance of MockController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockController(t mockConstructorTestingTNewMockController) *MockController {
	mock := &MockController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}