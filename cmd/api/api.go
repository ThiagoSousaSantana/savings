package api

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"

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

	expenseHandler := expense.NewHandler()
	expenseHandler.RegisterRoutes(app)

	return app.Listen(s.addr)
}
