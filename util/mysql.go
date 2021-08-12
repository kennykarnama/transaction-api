package util

import (
	"errors"
	"github.com/go-sql-driver/mysql"
)

const (
	// DuplicateEntryCode represents MySQL error code 1062
	DuplicateEntryCode uint16 = 0x426
)

// GetErrorNumber infers error number from mysql error
func GetErrorNumber(err error) (uint16, error) {
	mysqlError, ok := err.(*mysql.MySQLError)

	if !ok {
		return 0, errors.New("errors.parsing invalid MySQLError")
	}

	return mysqlError.Number, nil
}

// IsDuplicatedEntryError checks if given error is mysql duplicated entry error
// by checking its error number
func IsDuplicatedEntryError(err error) bool {
	errorNumber, err := GetErrorNumber(err)

	if err != nil {
		return false
	}

	if errorNumber == DuplicateEntryCode {
		return true
	}

	return false
}
