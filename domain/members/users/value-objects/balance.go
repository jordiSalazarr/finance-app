package userVals

type Balance struct {
	Val int64
}

func NewUserBalance(val int64) Balance {
	return Balance{Val: val}
}
