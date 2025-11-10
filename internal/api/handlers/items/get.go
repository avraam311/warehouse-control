package items

import (
	"fmt"
	"net/http"

	"github.com/avraam311/warehouse-control/internal/api/handlers"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *Handler) GetItems(c *ginext.Context) {
	items, err := h.service.GetItems(c.Request.Context())
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to get items")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.OK(c.Writer, items)
}
