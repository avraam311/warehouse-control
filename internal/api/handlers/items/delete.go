package items

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/avraam311/warehouse-control/internal/api/handlers"
	"github.com/avraam311/warehouse-control/internal/repository/items"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *Handler) DeleteItem(c *ginext.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to convert param id into int")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid id: %s", err.Error()))
		return
	}
	id := uint(idInt)

	userIDAny, ok := c.Get("user_id")
	if !ok {
		zlog.Logger.Error().Msg("failed to get user_id from context")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("failed to get user_id from context"))
		return
	}
	userID := userIDAny.(uint)

	err = h.service.DeleteItem(c.Request.Context(), id, userID)
	if err != nil {
		if errors.Is(err, items.ErrItemNotFound) {
			zlog.Logger.Error().Err(err).Interface("id", id).Msg("item not found")
			handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("item not found"))
			return
		}

		zlog.Logger.Error().Err(err).Msg("failed to delete item")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.OK(c.Writer, "item deleted")
}
