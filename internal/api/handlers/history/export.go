package history

import (
	"fmt"
	"net/http"

	"github.com/avraam311/warehouse-control/internal/api/handlers"
	"github.com/avraam311/warehouse-control/internal/models"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *Handler) ExportHistory(c *ginext.Context) {
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=history.csv")

	userIDAny, ok := c.Get("user_id")
	if !ok {
		zlog.Logger.Error().Msg("user_id not found in context")
		handlers.Fail(c.Writer, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}
	userID := userIDAny.(uint)

	var filter models.HistoryFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to bind query params")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid query params"))
		return
	}

	data, err := h.service.ExportHistoryCSV(c.Request.Context(), &filter, userID)
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to export history")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(data)
}
