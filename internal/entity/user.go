package entity

type User struct {
	ID            uint32
	Login         string
	Password      string
	BalanceTotal  int
	WithdrawTotal int
}
