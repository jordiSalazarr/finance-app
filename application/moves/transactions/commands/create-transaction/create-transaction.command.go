package createtransaction

type CreateTransactionsCommand struct {
	Pk          string
	Description string
	Amount      int64
	Type        string
	GroupID     string
	Category    string
	PayedBY     string
	UserID      string
}
