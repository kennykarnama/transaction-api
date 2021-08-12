package transaction

import "time"

type CreateTransactionRequest struct {
	Transaction *Transaction `json:"transaction"`
}

type Transaction struct {
	ID              int64              `json:"id,omitempty"`
	UUID            string             `json:"uuid,omitempty"`
	TotalAmount     int64              `json:"totalAmount,omitempty"`
	DeviceTimestamp int64              `json:"deviceTimestamp,omitempty"`
	UserID          int64              `json:"userID,omitempty"`
	ChangeAmount    int64              `json:"changeAmount,omitempty"`
	PaymentMethod   string             `json:"paymentMethod,omitempty"`
	PaidAmount      int64              `json:"paidAmount,omitempty"`
	CreatedAt       time.Time          `json:"createdAt,omitempty"`
	UpdatedAt       time.Time          `json:"updatedAt,omitempty"`
	Items           []*TransactionItem `json:"items,omitempty"`
}

type TransactionItem struct {
	ID            int64  `json:"id"`
	UUID          string `json:"uuid"`
	Title         string `json:"title"`
	Qty           int64  `json:"qty"`
	Price         int64  `json:"price"`
	TransactionID int64  `json:"transactionID"`
}

type CreateTransactionResponse struct {
	Transaction *Transaction `json:"transaction"`
}

type CreateAndPayTransactionRequest struct {
	Transaction *Transaction `json:"transaction"`
}

type CreateAndPayTransactionResponse struct {
	TransactionID   int64  `json:"transactionID"`
	TransactionUUID string `json:"transactionUUID"`
	ChangeAmount    int64  `json:"changeAmount"`
}

type ListUserTransactionQueryParam struct {
	Page     int32
	PageSize int32
}

type Pagination struct {
	Page      int32 `json:"page"`
	PageSize  int32 `json:"pageSize"`
	TotalData int64 `json:"totalData"`
	TotalPage int   `json:"totalPage"`
}

type ListUserTransactionResponse struct {
	Transactions []*Transaction `json:"transactions"`
	Pagination   *Pagination    `json:"pagination"`
}

type GetTransactionDetailResponse struct {
	Transaction *Transaction `json:"transaction"`
}

type UpdateTransactionRequest struct {
	Transaction *Transaction `json:"transaction"`
}
