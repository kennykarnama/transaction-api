package transaction

import (
	"context"
	"transaction-api/domain/models/transaction"
)

type Service interface {
	CreateTransaction(ctx context.Context, trans *transaction.Transaction) error
}
