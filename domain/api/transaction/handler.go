package transaction

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"time"
	"transaction-api/domain/api/shared"
	transEntity "transaction-api/domain/models/transaction"
	"transaction-api/domain/service/transaction"
	"transaction-api/util"
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

// TODO:
// benerin validasinya
// kalau 0 barang
// nested struct lah
func (h *Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var req CreateTransactionRequest
	err := util.DecodeToStruct(r.Body, &req)
	if err != nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	if err := h.validate.Struct(&req); err != nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	newTrans := &transEntity.Transaction{
		UserID:           req.UserID,
		DeviceTimestamp:  time.Unix(req.DeviceTimestamp, 0),
		TransactionItems: []*transEntity.TransactionItem{},
	}

	for _, itemReq := range req.Items {
		newTrans.TransactionItems = append(newTrans.TransactionItems, &transEntity.TransactionItem{
			Title: itemReq.Title,
			Qty:   itemReq.Qty,
			Price: itemReq.Price,
		})
	}
	err = h.transSvc.CreateTransaction(h.ctx, newTrans)

	resp := CreateTransactionResponse{
		TransactionID:   newTrans.ID,
		TransactionUUID: newTrans.UUID,
	}

	shared.ResponseJson(w, resp, http.StatusCreated)

	return
}

func (h *Handler) CreateAndPayTransaction(w http.ResponseWriter, r *http.Request) {
	var req CreateAndPayTransactionRequest
	err := util.DecodeToStruct(r.Body, &req)
	if err != nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	if err := h.validate.Struct(&req); err != nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	newTrans := &transEntity.Transaction{
		UserID:           req.UserID,
		DeviceTimestamp:  time.Unix(req.DeviceTimestamp, 0),
		TransactionItems: []*transEntity.TransactionItem{},
		PaymentMethod:    req.PaymentMethod,
		PaidAmount:       req.PaidAmount,
	}

	for _, itemReq := range req.Items {
		newTrans.TransactionItems = append(newTrans.TransactionItems, &transEntity.TransactionItem{
			Title: itemReq.Title,
			Qty:   itemReq.Qty,
			Price: itemReq.Price,
		})
	}
	err = h.transSvc.CreateAndPayTransaction(h.ctx, newTrans)

	resp := CreateAndPayTransactionResponse{
		TransactionID:   newTrans.ID,
		TransactionUUID: newTrans.UUID,
		ChangeAmount:    newTrans.ChangeAmount,
	}

	shared.ResponseJson(w, resp, http.StatusCreated)

	return
}

func (h *Handler) ListUserTransaction(w http.ResponseWriter, r *http.Request) {
	pathVariables := mux.Vars(r)
	userID, err := strconv.ParseInt(pathVariables["id"], 10, 64)

	if err != nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	paging := &transEntity.Pagination{
		Page:     1,
		PageSize: 10,
	}

	q := r.URL.Query()

	if v := q.Get("page"); v != "" {
		parsed, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			shared.ResponseJson(w, shared.ErrorResponse{
				Message: err.Error(),
			}, http.StatusBadRequest)
			return
		}
		paging.Page = int32(parsed)
	}
	if v := q.Get("pageSize"); v != "" {
		parsed, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			shared.ResponseJson(w, shared.ErrorResponse{
				Message: err.Error(),
			}, http.StatusBadRequest)
			return
		}
		paging.PageSize = int32(parsed)
	}

	transactions, err := h.transSvc.ListTransactionsByID(h.ctx, userID, paging)
	if err != nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	resp := &ListUserTransactionResponse{
		Transactions: []*Transaction{},
		Pagination: &Pagination{
			Page:      paging.Page,
			PageSize:  paging.PageSize,
			TotalData: paging.TotalData,
			TotalPage: paging.TotalPage,
		},
	}

	for _, trans := range transactions {
		parsedTrans := &Transaction{
			ID:              trans.ID,
			UUID:            trans.UUID,
			TotalAmount:     trans.TotalAmount,
			DeviceTimestamp: trans.DeviceTimestamp.Unix(),
			UserID:          trans.UserID,
			ChangeAmount:    trans.ChangeAmount,
			CreatedAt:       trans.CreatedAt,
			UpdatedAt:       trans.UpdatedAt,
			Items:           []*TransactionItem{},
		}
		if !strings.EqualFold(trans.PaymentMethod, "none") {
			parsedTrans.PaymentMethod = trans.PaymentMethod
		}
		resp.Transactions = append(resp.Transactions, parsedTrans)
	}

	shared.ResponseJson(w, resp, http.StatusOK)

	return
}

func (h *Handler) GetTransactionDetail(w http.ResponseWriter, r *http.Request) {
	pathVariables := mux.Vars(r)
	orderID, err := strconv.ParseInt(pathVariables["id"], 10, 64)

	if err != nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	trans, err := h.transSvc.GetTransactionDetailByID(h.ctx, orderID)
	if err != nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	shared.ResponseJson(w, ParseToTranscationResponse(trans), http.StatusOK)
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
