package server

import (
	"net/http"

	"github.com/avraam311/warehouse-control/internal/api/handlers/items"
	"github.com/avraam311/warehouse-control/internal/api/middlewares"

	"github.com/wb-go/wbf/ginext"
)

func NewRouter(ginMode string, handlerItem *items.Handler) *ginext.Engine {
	e := ginext.New(ginMode)

	e.Use(middlewares.CORSMiddleware())
	e.Use(ginext.Logger())
	e.Use(ginext.Recovery())

	api := e.Group("/warehouse-control/api")
	{
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
