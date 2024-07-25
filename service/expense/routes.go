package expense

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /expenses", createExpense)
	mux.HandleFunc("GET /expenses", findExpenses)
	mux.HandleFunc("GET /expenses/{id}", findExpenseById)
	mux.HandleFunc("DELETE /expenses/{id}", deleteExpense)
}

func createExpense(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expense created")
}

func findExpenses(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All expenses")
}

func findExpenseById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, id)
}

func deleteExpense(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "Expense "+id+" deleted")
}
