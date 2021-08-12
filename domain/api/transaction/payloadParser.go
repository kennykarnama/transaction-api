package transaction

import (
	"time"
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
	result.PaymentMethod = transModel.PaymentMethod
	for _, itemReq := range transModel.TransactionItems {
		result.Items = append(result.Items, &TransactionItem{
			ID:            itemReq.ID,
			UUID:          itemReq.UUID,
			Title:         itemReq.Title,
			Qty:           itemReq.Qty,
			Price:         itemReq.Price,
			TransactionID: itemReq.TransactionID,
		})
	}

	return result
}

func ParseToTransaction(transaction *Transaction) *transEntity.Transaction {
	trans := &transEntity.Transaction{
		ID:               transaction.ID,
		UUID:             transaction.UUID,
		UserID:           transaction.UserID,
		DeviceTimestamp:  time.Unix(transaction.DeviceTimestamp, 0),
		TotalAmount:      transaction.TotalAmount,
		PaidAmount:       transaction.PaidAmount,
		ChangeAmount:     transaction.ChangeAmount,
		TransactionItems: []transEntity.TransactionItem{},
	}
	v := transEntity.StringToPaymentMethod(transaction.PaymentMethod)
	trans.PaymentMethod = v.String()
	for _, item := range transaction.Items {
		trans.TransactionItems = append(trans.TransactionItems, transEntity.TransactionItem{
			ID:            item.ID,
			UUID:          item.UUID,
			TransactionID: item.TransactionID,
			Title:         item.Title,
			Qty:           item.Qty,
			Price:         item.Price,
		})
	}

	return trans
}
