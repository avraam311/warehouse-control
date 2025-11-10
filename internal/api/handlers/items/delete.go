package items

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/avraam311/warehouse-control/internal/api/handlers"

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

	err = h.service.DeleteItem(c.Request.Context(), id)
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to delete item")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.OK(c.Writer, "item deleted")
}
