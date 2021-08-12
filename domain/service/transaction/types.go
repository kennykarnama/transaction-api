package transaction

import (
	"context"
	"transaction-api/domain/models/transaction"
)

type Service interface {
	CreateTransaction(ctx context.Context, trans *transaction.Transaction) error
	CreateAndPayTransaction(ctx context.Context, trans *transaction.Transaction) error
	ListTransactionsByID(ctx context.Context, userID int64, paging *transaction.Pagination) ([]*transaction.Transaction, error)
	GetTransactionDetailByID(ctx context.Context, transID int64) (*transaction.Transaction, error)
	DeleteTransactionByID(ctx context.Context, transID int64) error
}
