package domainUsers

type UserRepository interface {
	Save(user User) error
	Exists(mail string) bool
	GetVerifiedUser(mail string) (User, error)
	VerificateUser(mail string) error
	GetUser(mail string) (User, error)
	GetById(id string) (User, error)
	UpdateCurrentBalance(id string, val int64) error
	//THis has to be a db transaction
	UpdateActorsCurrentBalance(debtor_id, payed_by string, val int64) error
	GetAll() ([]User, error)
}
