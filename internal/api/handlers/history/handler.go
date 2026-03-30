package history

import (
	"github.com/avraam311/warehouse-control/internal/service/items"
)

type Handler struct {
	service *items.Service
}

func NewHandler(service *items.Service) *Handler {
	return &Handler{
		service: service,
	}
}
