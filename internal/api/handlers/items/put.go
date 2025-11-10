package items

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/avraam311/warehouse-control/internal/api/handlers"
	"github.com/avraam311/warehouse-control/internal/models"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *Handler) PutItem(c *ginext.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to convert param id into int")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid id: %s", err.Error()))
		return
	}
	id := uint(idInt)

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

	err = h.service.ReplaceItem(c.Request.Context(), id, &item)
	if err != nil {
		zlog.Logger.Error().Err(err).Interface("item", item).Msg("failed to replace item")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.Created(c.Writer, "item replaced")
}
