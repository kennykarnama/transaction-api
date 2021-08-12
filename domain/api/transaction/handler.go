package transaction

import (
	"context"
	"github.com/go-playground/validator/v10"
	"net/http"
	"transaction-api/domain/service/transaction"
)

type Handler struct {
	ctx      context.Context
	validate *validator.Validate
	transSvc transaction.Service
}

func NewHandler(ctx context.Context, validate *validator.Validate, transSvc transaction.Service) *Handler {
	return &Handler{
		ctx:      ctx,
		transSvc: transSvc,
		validate: validate,
	}
}

func (h *Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) CreateAndPayTransaction(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) ListTransaction(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) GetTransactionDetail(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) DeleteTransactionByID(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) DeleteTransactionItems(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *Handler) UpdateTransactionByID(w http.ResponseWriter, r *http.Request) {
	return
}
