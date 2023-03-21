package handlers

import (
	"github.com/GOsling-Inc/GOsling/services"
)

type Handler struct {
	service *services.Service
}

func New(s *services.Service) *Handler {
	return &Handler{
		service: s,
	}
}

