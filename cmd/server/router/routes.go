package router

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/server/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Router struct {
	server *gin.Engine
	list   []domain.Ticket
}

func NewRouter(server *gin.Engine, list []domain.Ticket) *Router {
	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	return &Router{
		server,
		list,
	}
}

func (r *Router) MapRoutes() {
	repository := tickets.NewRepository(r.list)
	ticketService := tickets.NewService(repository)
	handler := handler.NewService(ticketService)

	group := r.server.Group("/ticket")
	{
		group.GET("/getByCountry/:dest", handler.GetTicketsByCountry())
		group.GET("/getAverage/:dest", handler.AverageDestination())
	}
}
