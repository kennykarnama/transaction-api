package transaction

import (
	"gorm.io/gorm"
	"time"
)

type PaymentMethod int

const (
	NonePaymentMethod PaymentMethod = iota
	Cash
	Card
)

func (pm PaymentMethod) String() string {
	switch pm {
	case Cash:
		return "cash"
	case Card:
		return "card"
	default:
		return "none"
	}
}

func (pm PaymentMethod) MaskingNone() string {
	s := pm.String()
	if s == "none" {
		return ""
	}
	return s
}

func StringToPaymentMethod(s string) PaymentMethod {
	switch s {
	case "cash":
		return Cash
	case "card":
		return Card
	default:
		return NonePaymentMethod
	}
}

type Transaction struct {
	ID               int64
	UUID             string
	UserID           int64
	DeviceTimestamp  time.Time
	TotalAmount      int64
	PaidAmount       int64
	ChangeAmount     int64
	PaymentMethod    string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
	TransactionItems []TransactionItem
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
	DeletedAt     gorm.DeletedAt
}

func (t *Transaction) SetTotalAmount() int64 {
	var total int64
	for _, item := range t.TransactionItems {
		total += item.Qty * item.Price
	}
	t.TotalAmount = total
	return total
}

type Pagination struct {
	Page      int32
	PageSize  int32
	TotalData int64
	TotalPage int
}

func (t *Transaction) IsItemExist(itemID int64) bool {
	for _, item := range t.TransactionItems {
		if item.ID == itemID {
			return true
		}
	}
	return false
}
