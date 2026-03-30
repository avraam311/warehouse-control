package history

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/avraam311/warehouse-control/internal/api/handlers"
	"github.com/avraam311/warehouse-control/internal/models"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

type CompareRequest struct {
	ItemID   uint `json:"item_id" validate:"required"`
	Version1 uint `json:"version1" validate:"required"`
	Version2 uint `json:"version2" validate:"required"`
}

func (h *Handler) CompareVersions(c *ginext.Context) {
	userIDAny, ok := c.Get("user_id")
	if !ok {
		zlog.Logger.Error().Msg("user_id not found")
		handlers.Fail(c.Writer, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}
	userID := userIDAny.(uint)

	var req CompareRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to decode compare request")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid request"))
		return
	}

	filter1 := &models.HistoryFilter{ItemID: &req.ItemID, Limit: func() *int { l := 1; return &l }()}
	history1, err := h.service.GetHistory(c.Request.Context(), filter1, userID)
	if err != nil || len(history1) == 0 {
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("version1 not found"))
		return
	}

	filter2 := &models.HistoryFilter{ItemID: &req.ItemID, Limit: func() *int { l := 1; return &l }()}
	history2, err := h.service.GetHistory(c.Request.Context(), filter2, userID)
	if err != nil || len(history2) == 0 {
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("version2 not found"))
		return
	}

	v1 := history1[0]
	v2 := history2[0]

	diff := &models.VersionDiff{
		ItemID:   req.ItemID,
		Version1: req.Version1,
		Version2: req.Version2,
	}

	if v1.Name != v2.Name {
		diff.NameDiff = &v1.Name 
	}
	if v1.Description != v2.Description {
		diff.DescDiff = &v1.Description
	}
	if v1.Price != v2.Price {
		diff.PriceDiff = &v1.Price
	}

	handlers.OK(c.Writer, diff)
}
