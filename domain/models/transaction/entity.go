package transaction

import "time"

type Transaction struct {
	ID              int64
	UUID            string
	UserID          int64
	DeviceTimeStamp time.Time
	TotalAmount     int64
	PaidAmount      int64
	ChangeAmount    int64
	PaymentMethod   string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
	Items           []*TransactionItem
}

type TransactionItem struct {
	ID            int64
	UUID          string
	TransactionID int64
	Title         string
	Qty           int64
	Price         int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}
