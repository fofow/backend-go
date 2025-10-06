package handler

import (
	"gitlab.com/msstoci/popow-api/internal/service"
)

type Handler struct {
	svc service.Service
}

func New(svc service.Service) *Handler {
	return &Handler{svc}
}
