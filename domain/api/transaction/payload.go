package transaction

type CreateTransactionRequest struct {
	UserID          int64 `json:"userID"`
	DeviceTimestamp int64 `json:"deviceTimestamp"`
	Items           []*TransactionItem
}

type TransactionItem struct {
	Title string `json:"title"`
	Qty   int64  `json:"qty"`
	Price int64  `json:"price"`
}

type CreateTransactionResponse struct {
	TransactionID   int64  `json:"transactionID"`
	TransactionUUID string `json:"transactionUUID"`
}
