package transaction

import (
	"strings"
	transEntity "transaction-api/domain/models/transaction"
)

func ParseToTranscationResponse(transModel *transEntity.Transaction) *Transaction {
	if transModel == nil {
		return &Transaction{}
	}
	result := &Transaction{
		ID:              transModel.ID,
		UUID:            transModel.UUID,
		TotalAmount:     transModel.TotalAmount,
		DeviceTimestamp: transModel.DeviceTimestamp.Unix(),
		UserID:          transModel.UserID,
		ChangeAmount:    transModel.ChangeAmount,
		CreatedAt:       transModel.CreatedAt,
		UpdatedAt:       transModel.UpdatedAt,
		Items:           []*TransactionItem{},
	}
	if !strings.EqualFold(transModel.PaymentMethod, "none") {
		result.PaymentMethod = transModel.PaymentMethod
	}
	for _, itemReq := range transModel.TransactionItems {
		result.Items = append(result.Items, &TransactionItem{
			Title: itemReq.Title,
			Qty:   itemReq.Qty,
			Price: itemReq.Price,
		})
	}

	return result
}
