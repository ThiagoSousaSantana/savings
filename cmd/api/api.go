package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"

	"github.com/ThiagoSousaSantana/saving/cmd/db"
	"github.com/ThiagoSousaSantana/saving/service/expense"
)

type APIServer struct {
	addr    string
	queries *db.Queries
}

func NewAPIServer(addr string, queries *db.Queries) *APIServer {
	return &APIServer{
		addr:    addr,
		queries: queries,
	}
}

func (s *APIServer) Run() error {
	app := fiber.New()
	app.Use(logger.New())
	v1 := app.Group("/api/v1")

	expenseHandler := expense.NewHandler()
	expenseHandler.RegisterRoutes(v1)

	return app.Listen(s.addr)
}
