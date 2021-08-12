package transaction

import "errors"

var (
	ErrNotSufficientBalance = errors.New("balance is not sufficient")
	ErrTransactionNotFound  = errors.New("transaction not found")
)
