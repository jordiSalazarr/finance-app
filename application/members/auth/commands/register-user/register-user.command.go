package auth_useCases

type RegisterUserCommand struct {
	Name           string
	Mail           string
	Password       string
	CurrentBalance int64
	MonthlyIncome  int64
}
