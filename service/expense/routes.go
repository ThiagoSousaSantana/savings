package expense

import (
	"github.com/gofiber/fiber/v3"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(route fiber.Router) {
	route.Post("/expenses", createExpense)
	route.Get("/expenses", findExpenses)
	route.Get("/expenses/:id", findExpenseById)
	route.Delete("/expenses/:id", deleteExpense)
}

func createExpense(c fiber.Ctx) error {
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
