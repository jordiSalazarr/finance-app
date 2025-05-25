package mysqlmoves

import (
	"fmt"
	"time"

	domainTransaction "finances.jordis.golang/domain/moves/transactions"
	"finances.jordis.golang/infrastructure/dbmodels"
	"gorm.io/gorm"
)

type TransactionsRepoMySQL struct {
	DB *gorm.DB
}

func NewTransactionsRepoMySQL(db *gorm.DB) *TransactionsRepoMySQL {
	return &TransactionsRepoMySQL{DB: db}
}

func (t *TransactionsRepoMySQL) GetById(id string) (domainTransaction.Transaction, error) {
	var transaction dbmodels.Transaction
	err := t.DB.Where("pk = ?", id).First(&transaction).Error
	if err != nil {
		return domainTransaction.Transaction{}, err
	}

	return mapToDomainTransaction(transaction)
}

func (t *TransactionsRepoMySQL) SaveOne(transaction domainTransaction.Transaction) error {
	dbTransaction := mapDomainToDBTransaction(transaction)
	return t.DB.Create(&dbTransaction).Error
}
func (t *TransactionsRepoMySQL) SaveMany(transaction []domainTransaction.Transaction) error {
	var dbTransactions []dbmodels.Transaction
	for _, dbTransaction := range transaction {
		dbTransactions = append(dbTransactions, mapDomainToDBTransaction(dbTransaction))
	}
	return t.DB.Create(&dbTransactions).Error
}

func (t *TransactionsRepoMySQL) MarkAsPayed(id string) error {
	return t.DB.Model(&dbmodels.Transaction{}).Where("pk = ?", id).Updates(map[string]interface{}{
		"already_payed": true,
		"updated_at":    time.Now(),
	}).Error
}

func (t *TransactionsRepoMySQL) GetUserTransactions(userID string, fromDate time.Time, toDate time.Time, category string) ([]*domainTransaction.Transaction, error) {
	var transactions []dbmodels.Transaction

	query := t.DB.Model(&dbmodels.Transaction{}).
		Where("user_pk = ?", userID).
		Where("created_at BETWEEN ? AND ?", fromDate, toDate)

	// if category != "all" {
	// 	query = query.Where("category = ?", category)
	// }

	err := query.Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return mapToDomainTransactions(transactions)
}

func (t *TransactionsRepoMySQL) GetGroupTransactions(groupID string, fromDate time.Time, toDate time.Time, category string) ([]*domainTransaction.Transaction, error) {
	var transactions []dbmodels.Transaction
	err := t.DB.Model(&dbmodels.Transaction{}).Where("group_pk = ?", groupID).Where("created_at BETWEEN ? AND ?", fromDate, toDate).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return mapToDomainTransactions(transactions)

}
func (t *TransactionsRepoMySQL) GetTransactionsPendingToRecieve(userID string) ([]*domainTransaction.Transaction, error) {
	var transactions []dbmodels.Transaction
	err := t.DB.Model(&dbmodels.Transaction{}).Where("payed_by = ?", userID).Where("already_payed = ?", false).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return mapToDomainTransactions(transactions)

}
func (t *TransactionsRepoMySQL) GetTransactionsPendingToPay(userID string) ([]*domainTransaction.Transaction, error) {
	var transactions []dbmodels.Transaction
	err := t.DB.Model(&dbmodels.Transaction{}).Where("user_pk = ?", userID).Where("already_payed = ?", false).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return mapToDomainTransactions(transactions)

}

func mapToDomainTransactions(dbTransactions []dbmodels.Transaction) ([]*domainTransaction.Transaction, error) {
	var domainTransactions []*domainTransaction.Transaction
	for _, dbTransaction := range dbTransactions {
		domainTransaction, err := mapToDomainTransaction(dbTransaction)
		if err != nil {
			fmt.Print(err.Error())
		}
		domainTransactions = append(domainTransactions, &domainTransaction)
	}
	return domainTransactions, nil
}
func mapToDomainTransaction(dbTransaction dbmodels.Transaction) (domainTransaction.Transaction, error) {
	return domainTransaction.FromExisting(dbTransaction.PK, dbTransaction.Description, dbTransaction.PayedBy,
		dbTransaction.AlreadyPayed, dbTransaction.Category, (dbTransaction.Amount), dbTransaction.Type,
		dbTransaction.GroupPK, dbTransaction.UserPK, dbTransaction.CreatedAt)

}

func mapDomainToDBTransaction(u domainTransaction.Transaction) dbmodels.Transaction {

	return dbmodels.Transaction{
		PK:           u.Pk.Val,
		Description:  u.Description.Val,
		Category:     u.Category.Val,
		Amount:       int64(u.Amount.Val),
		Type:         string(u.Type.Val),
		GroupPK:      u.GroupId.Val,
		UserPK:       u.UserId.Val,
		CreatedAt:    u.CreatedAt,
		PayedBy:      u.PayedBy.Val,
		AlreadyPayed: u.AlreadyPayed,
	}
}
