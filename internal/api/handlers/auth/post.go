package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/avraam311/warehouse-control/internal/api/handlers"
	"github.com/avraam311/warehouse-control/internal/models"

	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func (h *Handler) Register(c *ginext.Context) {
	var user models.UserWithRoleDTO
	if err := json.NewDecoder(c.Request.Body).Decode(&user); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to decode request body")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid request body: %s", err.Error()))
		return
	}

	// Валидация роли
	validRoles := map[string]struct{}{
		"admin":   {},
		"manager": {},
		"viewer":  {},
	}
	if _, ok := validRoles[user.Role]; !ok {
		zlog.Logger.Error().Msgf("invalid role: %s", user.Role)
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid role: %s", user.Role))
		return
	}

	if err := h.validator.Struct(user); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to validate request body")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("validation error: %s", err.Error()))
		return
	}

	id, err := h.service.Register(c.Request.Context(), &user)
	if err != nil {
		zlog.Logger.Error().Err(err).Interface("user", user).Msg("failed to register user")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.Created(c.Writer, id)
}

func (h *Handler) Login(c *ginext.Context) {
	var user models.UserDTO
	if err := json.NewDecoder(c.Request.Body).Decode(&user); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to decode request body")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid request body: %s", err.Error()))
		return
	}

	if err := h.validator.Struct(user); err != nil {
		zlog.Logger.Error().Err(err).Msg("failed to validate request body")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("validation error: %s", err.Error()))
		return
	}

	token, err := h.service.Login(c.Request.Context(), &user)
	if err != nil {
		zlog.Logger.Error().Err(err).Interface("token", user).Msg("failed to login")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.Created(c.Writer, token)
}
