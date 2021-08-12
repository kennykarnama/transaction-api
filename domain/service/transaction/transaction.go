package transaction

import (
	"context"
	uuid "github.com/satori/go.uuid"
	transactionEntity "transaction-api/domain/models/transaction"
	"transaction-api/domain/repository/transaction"
)

type ListTransactionsByUserIDResult struct {
}

type service struct {
	transRepo transaction.Repository
}

func NewService(transRepo transaction.Repository) *service {
	return &service{transRepo: transRepo}
}

func (s *service) CreateTransaction(ctx context.Context, trans *transactionEntity.Transaction) error {
	trans.UUID = uuid.NewV4().String()
	for _, item := range trans.TransactionItems {
		item.UUID = uuid.NewV4().String()
	}
	trans.SetTotalAmount()
	// TODO: maybe introduced enum ?
	if trans.PaymentMethod == "" {
		trans.PaymentMethod = "none"
	}
	return s.transRepo.CreateTransaction(ctx, trans)
}

func (s *service) CreateAndPayTransaction(ctx context.Context, trans *transactionEntity.Transaction) error {
	trans.SetTotalAmount()
	amountOfChange := trans.PaidAmount - trans.TotalAmount
	if amountOfChange < 0 {
		return ErrNotSufficientBalance
	}
	trans.ChangeAmount = amountOfChange
	return s.CreateTransaction(ctx, trans)
}

func (s *service) ListTransactionsByID(ctx context.Context, userID int64, paging *transactionEntity.Pagination) ([]*transactionEntity.Transaction, error) {
	return s.transRepo.GetTransactionsByUserID(ctx, userID, paging)
}

func (s *service) GetTransactionDetailByID(ctx context.Context, transID int64) (*transactionEntity.Transaction, error) {
	return s.transRepo.GetTransactionDetailByID(ctx, transID)
}
