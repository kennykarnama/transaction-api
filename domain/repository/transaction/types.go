package transaction

import (
	"context"
	"transaction-api/domain/models/transaction"
)

type Repository interface {
	CreateTransaction(ctx context.Context, trans *transaction.Transaction) error
	GetTransactionsByUserID(ctx context.Context, userID int64, pagination *transaction.Pagination) ([]*transaction.Transaction, error)
	GetTransactionDetailByID(ctx context.Context, transID int64) (*transaction.Transaction, error)
	DeleteTransactionByID(ctx context.Context, transID int64) error
	DeleteTransactionItemByIDs(ctx context.Context, itemID int64) error
	UpdateTransaction(ctx context.Context, transID int64, data *transaction.Transaction) error
}
