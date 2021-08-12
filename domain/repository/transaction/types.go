package transaction

import (
	"context"
	"transaction-api/domain/models/transaction"
)

type Repository interface {
	CreateTransaction(ctx context.Context, trans *transaction.Transaction) error
}
