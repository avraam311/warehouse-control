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

	auth := api.Group("/auth")
	{
		auth.POST("/login", handlerAuth.Login)
		auth.POST("/register", middlewares.RoleBasedAuthMiddleware(cfg.GetString("JWT_SECRET")), handlerAuth.Register)
	}

	items := api.Group("/items")
	items.Use(middlewares.RoleBasedAuthMiddleware(cfg.GetString("JWT_SECRET")))
	{
		items.POST("/", handlerItem.CreateItem)
		items.GET("/", handlerItem.GetItems)
		items.PUT("/:id", handlerItem.PutItem)
		items.DELETE("/:id", handlerItem.DeleteItem)
	}

	return e
}

func NewServer(addr string, router *ginext.Engine) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
