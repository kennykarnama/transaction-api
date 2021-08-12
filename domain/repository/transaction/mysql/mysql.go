package mysql

import (
	"context"
	"fmt"
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
