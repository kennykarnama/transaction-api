package transaction

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

	if req.Transaction == nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: "transcation payload not found",
		}, http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(&req); err != nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	newTrans := ParseToTransaction(req.Transaction)

	err = h.transSvc.CreateTransaction(h.ctx, newTrans)

	shared.ResponseJson(w, &CreateTransactionResponse{Transaction: ParseToTranscationResponse(newTrans)}, http.StatusCreated)

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

	if req.Transaction == nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: "transcation payload not found",
		}, http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(&req); err != nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	newTrans := ParseToTransaction(req.Transaction)

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
		resp.Transactions = append(resp.Transactions, ParseToTranscationResponse(trans))
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
		if err == transaction.ErrTransactionNotFound {
			shared.ResponseJson(w, shared.ErrorResponse{
				Message: err.Error(),
			}, http.StatusNotFound)
			return
		}
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	shared.ResponseJson(w, ParseToTranscationResponse(trans), http.StatusOK)
	return
}

func (h *Handler) DeleteTransactionByID(w http.ResponseWriter, r *http.Request) {
	pathVariables := mux.Vars(r)
	transID, err := strconv.ParseInt(pathVariables["id"], 10, 64)

	if err != nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	err = h.transSvc.DeleteTransactionByID(h.ctx, transID)
	if err != nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	shared.ResponseJson(w, shared.Empty{}, http.StatusOK)
	return
}

func (h *Handler) DeleteTransactionItem(w http.ResponseWriter, r *http.Request) {
	pathVariables := mux.Vars(r)
	orderID, err := strconv.ParseInt(pathVariables["orderID"], 10, 64)
	if err != nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	itemID, err := strconv.ParseInt(pathVariables["itemID"], 10, 64)
	if err != nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	err = h.transSvc.DeleteTransactionItemByIDs(h.ctx, orderID, itemID)
	if err != nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	shared.ResponseJson(w, shared.Empty{}, http.StatusOK)
	return
}

func (h *Handler) UpdateTransactionByID(w http.ResponseWriter, r *http.Request) {
	pathVariables := mux.Vars(r)
	orderID, err := strconv.ParseInt(pathVariables["id"], 10, 64)
	if err != nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	var req UpdateTransactionRequest
	err = util.DecodeToStruct(r.Body, &req)
	if err != nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	newTrans := ParseToTransaction(req.Transaction)

	err = h.transSvc.UpdateTranscationByID(h.ctx, orderID, newTrans)
	if err != nil {
		shared.ResponseJson(w, shared.ErrorResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	shared.ResponseJson(w, ParseToTranscationResponse(newTrans), http.StatusOK)
	return
}
