package entity

type User struct {
	ID             uint32
	Login          string
	Password       string
	Balance_total  int
	Withdraw_total int
}
