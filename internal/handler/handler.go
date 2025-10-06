package handler

import (
	"github.com/fofow/backend-go/internal/service"
)

type Handler struct {
	svc service.Service
}

func New(svc service.Service) *Handler {
	return &Handler{svc}
}
