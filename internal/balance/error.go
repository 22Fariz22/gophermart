package balance

import "errors"

var (
	ErrNotEnoughFunds     = errors.New("not enough funds")
	ErrInvalidOrderNumber = errors.New("invalid order number")
)
