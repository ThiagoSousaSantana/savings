package routes

import (
	"net/http"

	"go.uber.org/zap"
)

type IncomeHandler struct {
	log *zap.Logger
}

func (*IncomeHandler) Pattern() string {
	return "/incomes"
}

func NewIncomeHandler(log *zap.Logger) *IncomeHandler {
	return &IncomeHandler{
		log: log,
	}
}

func (h *IncomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Request received", zap.String("path", r.URL.Path))
	w.Write([]byte("Income handler hit"))
}
