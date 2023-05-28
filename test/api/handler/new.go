package handler

import "go-chi-example/api/controller"

// Handler handles admin routes
type Handler struct {
	todoCtrl controller.Controller
}

// New creates and returns a new instance of handler
func New(todoCtrl controller.Controller) Handler {
	return Handler{todoCtrl}
}
