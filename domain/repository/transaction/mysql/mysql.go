package mysql

import (
	"context"
	"fmt"
	"github.com/kennykarnama/gorm-paginator/pagination"
	"gorm.io/gorm"
	"transaction-api/domain/models/transaction"
	transRepo "transaction-api/domain/repository/transaction"
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
	q = q.Where("user_id = ? AND deleted_at IS NULL", userID)
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
		if err == gorm.ErrRecordNotFound {
			return nil, transRepo.ErrRecordNotFound
		}
		return nil, err
	}
	return &trans, nil
}

func (r *repository) DeleteTransactionByID(ctx context.Context, transID int64) error {
	err := r.db.Delete(&transaction.Transaction{}, "id = ?", transID).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteTransactionItemByIDs(ctx context.Context, itemID int64) error {
	r.db.Transaction(func(tx *gorm.DB) error {
		// get detail itemID
		var item transaction.TransactionItem
		err := r.db.Model(&transaction.TransactionItem{}).First(&item, "id = ?", itemID).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return transRepo.ErrRecordNotFound
			}
			return err
		}
		// update transaction amount
		err = r.db.Model(&transaction.Transaction{}).Where("id = ?",
			item.TransactionID).Update("total_amount", gorm.Expr("total_amount - ?", item.Qty*item.Price)).Error
		if err != nil {
			return err
		}

		err = r.db.Delete(&transaction.TransactionItem{}, "id = ?", itemID).Error
		if err != nil {
			return err
		}

		return nil
	})
	return nil
}

func (r *repository) UpdateTransaction(ctx context.Context, transID int64, data *transaction.Transaction) error {
	db := r.db.Session(&gorm.Session{FullSaveAssociations: true})
	err := db.Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}
