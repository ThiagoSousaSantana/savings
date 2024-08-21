package routes

import (
	"net/http"

	"go.uber.org/zap"
)

type ExpenseHandler struct {
	log *zap.Logger
}

func (*ExpenseHandler) Pattern() string {
	return "/expenses"
}

func NewExpenseHandler(log *zap.Logger) *ExpenseHandler {
	return &ExpenseHandler{
		log: log,
	}
}

func (h *ExpenseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Request received", zap.String("path", r.URL.Path))
	w.Write([]byte("Expense handler hit"))
}
