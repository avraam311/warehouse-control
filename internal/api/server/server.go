package server

import (
	"net/http"

	"github.com/avraam311/warehouse-control/internal/api/handlers/auth"
	"github.com/avraam311/warehouse-control/internal/api/handlers/items"
	"github.com/avraam311/warehouse-control/internal/api/middlewares"

	"github.com/wb-go/wbf/config"
	"github.com/wb-go/wbf/ginext"
)

func NewRouter(cfg *config.Config, handlerItem *items.Handler, handlerAuth *auth.Handler) *ginext.Engine {
	e := ginext.New(cfg.GetString("server.gin_mode"))

	e.Use(middlewares.CORSMiddleware())
	e.Use(ginext.Logger())
	e.Use(ginext.Recovery())

	api := e.Group("/warehouse-control/api")
	api.Use(middlewares.RoleBasedAuthMiddleware(cfg.GetString("JWT_SECRET")))
	{
		api.POST("/auth/register", handlerAuth.Register)
		api.POST("/auth/login", handlerAuth.Login)

		api.POST("/items", handlerItem.CreateItem)
		api.GET("/items", handlerItem.GetItems)
		api.PUT("/items/:id", handlerItem.PutItem)
		api.DELETE("/items/:id", handlerItem.DeleteItem)
	}

	return e
}

func NewServer(addr string, router *ginext.Engine) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
