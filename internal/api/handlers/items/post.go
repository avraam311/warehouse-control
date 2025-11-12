package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/avraam311/warehouse-control/internal/api/handlers"
	"github.com/avraam311/warehouse-control/internal/models"
	"github.com/avraam311/warehouse-control/internal/repository/items"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *Handler) CreateItem(c *ginext.Context) {
	var item models.ItemDTO
	if err := json.NewDecoder(c.Request.Body).Decode(&item); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to decode request body")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid request body: %s", err.Error()))
		return
	}

	if err := h.validator.Struct(item); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to validate request body")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("validation error: %s", err.Error()))
		return
	}

	id, err := h.service.CreateItem(c.Request.Context(), &item)
	if err != nil {
		if errors.Is(err, items.ErrDuplicateItemName) {
			zlog.Logger.Error().Err(err).Interface("item", item).Msg("item.name already exists")
			handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("item.name already exists"))
			return
		} else if errors.Is(err, items.ErrDuplicateItemDescription) {
			zlog.Logger.Error().Err(err).Interface("item", item).Msg("item.desctiption already exists")
			handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("item.description already exists"))
			return
		}

		zlog.Logger.Error().Err(err).Interface("item", item).Msg("failed to create item")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.Created(c.Writer, id)
}
