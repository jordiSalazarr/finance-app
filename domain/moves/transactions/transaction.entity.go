package domainTransaction

import (
	"time"

	"finances.jordis.golang/domain"
	transactionVals "finances.jordis.golang/domain/moves/transactions/value-objects"
)

type Transaction struct {
	Pk           domain.UUID
	Description  transactionVals.Description
	Category     transactionVals.Category
	Amount       transactionVals.Amount
	Type         transactionVals.Type
	GroupId      domain.UUID
	UserId       domain.UUID
	CreatedAt    time.Time
	PayedBy      domain.UUID
	AlreadyPayed bool
}

func FromExisting(pk string, descrtiptionStr string, payedBy string, alreadyPayed bool,
	categoryStr string, amountFloat int64, typeStr string, groupId string, userID string, createdAt time.Time) (Transaction, error) {

	descr, err := transactionVals.NewDescription(descrtiptionStr)
	if err != nil {
		return Transaction{}, err
	}
	category, err := transactionVals.NewCategory(categoryStr)
	if err != nil {
		return Transaction{}, err
	}
	amount, err := transactionVals.NewAmount(int64(amountFloat))
	if err != nil {
		return Transaction{}, err
	}
	transType, err := transactionVals.NewType(typeStr)
	if err != nil {
		return Transaction{}, err
	}
	payedByID := domain.UUID{Val: payedBy}
	userId := domain.UUID{Val: userID}
	groupID := domain.UUID{Val: groupId}
	transactionID := domain.UUID{Val: pk}

	return Transaction{
		Pk:           transactionID,
		Description:  descr,
		Category:     category,
		Amount:       amount,
		Type:         transType,
		PayedBy:      payedByID,
		UserId:       userId,
		GroupId:      groupID,
		AlreadyPayed: alreadyPayed,
		CreatedAt:    createdAt,
	}, nil

}

func New(descrtiptionStr string, payedBy string, alreadyPayed bool,
	categoryStr string, amountFloat int64, typeStr string, groupId string, userID string,
) (Transaction, error) {

	descr, err := transactionVals.NewDescription(descrtiptionStr)
	if err != nil {
		return Transaction{}, err
	}
	category, err := transactionVals.NewCategory(categoryStr)
	if err != nil {
		return Transaction{}, err
	}
	amount, err := transactionVals.NewAmount(int64(amountFloat))
	if err != nil {
		return Transaction{}, err
	}
	transType, err := transactionVals.NewType(typeStr)
	if err != nil {
		return Transaction{}, err
	}
	payedByID := domain.UUID{Val: payedBy}
	userId := domain.UUID{Val: userID}
	groupID := domain.UUID{Val: groupId}
	transactionID := domain.UUID{Val: domain.NewUUID()}

	return Transaction{
		Pk:           transactionID,
		Description:  descr,
		Category:     category,
		Amount:       amount,
		Type:         transType,
		PayedBy:      payedByID,
		UserId:       userId,
		GroupId:      groupID,
		AlreadyPayed: alreadyPayed,
		CreatedAt:    time.Now(),
	}, nil

}
