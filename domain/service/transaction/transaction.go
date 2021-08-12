package transaction

import (
	"context"
	transactionEntity "transaction-api/domain/models/transaction"
	"transaction-api/domain/repository/transaction"
)

type service struct {
	transRepo transaction.Repository
}

func NewService(transRepo transaction.Repository) *service {
	return &service{transRepo: transRepo}
}

func (s *service) CreateTransaction(ctx context.Context, trans *transactionEntity.Transaction) error {
	return s.transRepo.CreateTransaction(ctx, trans)
}
