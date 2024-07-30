package expense

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"

	"github.com/ThiagoSousaSantana/saving/cmd/db"
)

type Handler struct {
	queries *db.Queries
}

func NewHandler(queries *db.Queries) *Handler {
	return &Handler{
		queries: queries,
	}
}

func (h *Handler) RegisterRoutes(route fiber.Router) {
	route.Post("/expenses", createExpense)
	route.Get("/expenses", findExpenses)
	route.Get("/expenses/:id", findExpenseById)
	route.Delete("/expenses/:id", deleteExpense)
}

func (s *Handler) createExpense(c fiber.Ctx) error {
	var expense db.CreateExpenseParams
	err := json.NewDecoder(c.Body()).Decode(&expense)
	if err != nil {
		return err
	}

	err = s.queries.CreateExpense(c.Context(), expense)
	if err != nil {
		return err
	}

	return c.SendString("Expense created")
}

func findExpenses(c fiber.Ctx) error {
	return c.SendString("All expenses")
}

func findExpenseById(c fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString(id)
}

func deleteExpense(c fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString(id)
}
