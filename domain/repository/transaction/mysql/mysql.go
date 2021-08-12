package mysql

import (
	"context"
	"fmt"
	"github.com/kennykarnama/gorm-paginator/pagination"
	"gorm.io/gorm"
	"transaction-api/domain/models/transaction"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateTransaction(ctx context.Context, trans *transaction.Transaction) error {
	db := r.db.Session(&gorm.Session{FullSaveAssociations: true})
	err := db.Save(trans).Error
	if err != nil {
		return fmt.Errorf("action=repo.createTransaction err=%v", err)
	}
	return nil
}

func (r *repository) GetTransactionsByUserID(ctx context.Context, userID int64, paging *transaction.Pagination) ([]*transaction.Transaction, error) {
	var result []*transaction.Transaction
	q := r.db.Model(&transaction.Transaction{})
	q = q.Where("user_id = ?", userID)
	paginator := pagination.Paging(&pagination.Param{
		DB:      q,
		Page:    int(paging.Page),
		Limit:   int(paging.PageSize),
		ShowSQL: true,
	}, &result)
	if q.Error != nil {
		return nil, q.Error
	}
	paging.TotalData = paginator.TotalRecord
	paging.TotalPage = paginator.TotalPage
	return result, nil
}

func (r *repository) GetTransactionDetailByID(ctx context.Context, transID int64) (*transaction.Transaction, error) {
	var trans transaction.Transaction
	err := r.db.Preload("TransactionItems").First(&trans, "id = ?", transID).Error
	if err != nil {
		return nil, err
	}
	return &trans, nil
}
