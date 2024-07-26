package api

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"

	"github.com/ThiagoSousaSantana/saving/cmd/db"
	"github.com/ThiagoSousaSantana/saving/service/expense"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	app := fiber.New()
	app.Use(logger.New())
	v1 := app.Group("/api/v1")

	queries := db.New(s.db)

	expenseHandler := expense.NewHandler(queries)
	expenseHandler.RegisterRoutes(v1)

	return app.Listen(s.addr)
}
