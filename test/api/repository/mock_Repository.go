// Code generated by mockery v2.20.0. DO NOT EDIT.

package repository

import (
	context "context"
	models "go-chi-example/api/models"

	mock "github.com/stretchr/testify/mock"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// CreateTodo provides a mock function with given fields: ctx, t
func (_m *MockRepository) CreateTodo(ctx context.Context, t *models.Todo) error {
	ret := _m.Called(ctx, t)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Todo) error); ok {
		r0 = rf(ctx, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllTodo provides a mock function with given fields: ctx
func (_m *MockRepository) GetAllTodo(ctx context.Context) ([]ModelTodo, error) {
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
func (_m *MockRepository) GetTodoByID(ctx context.Context, id int) (ModelTodo, error) {
	ret := _m.Called(ctx, id)

	var r0 ModelTodo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (ModelTodo, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) ModelTodo); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(ModelTodo)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRepository(t mockConstructorTestingTNewMockRepository) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
