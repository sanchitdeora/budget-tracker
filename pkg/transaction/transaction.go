package transaction

import (
	"context"

	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/models"
)

type Service interface {
	GetTransactions(ctx context.Context, transactions *[]models.Transaction) (error)
	GetTransactionById(ctx context.Context, id string, transaction *models.Transaction) (error)
	CreateTransaction(ctx context.Context, transaction models.Transaction) (string, error)
	UpdateTransactionById(ctx context.Context, id string, transaction models.Transaction) (string, error)
	DeleteTransactionById(ctx context.Context, id string) (string, error)
}

type Opts struct {}

type serviceImpl struct {
	*Opts
}

func NewService(opts *Opts) Service {
	return &serviceImpl{Opts: opts}
}


func (s *serviceImpl) GetTransactions(ctx context.Context, transactions *[]models.Transaction) (error) {
	// TODO: input validation
	return db.GetAllTransactions(ctx, transactions)
}

func (s *serviceImpl) GetTransactionById(ctx context.Context, id string, transaction *models.Transaction) (error) {
	// TODO: input validation
	return db.GetTransactionRecordById(ctx, id, transaction)
}

func (s *serviceImpl) CreateTransaction(ctx context.Context, transaction models.Transaction) (string, error) {
	// TODO: input validation
	transaction.SetCategory()
	return db.InsertTransactionRecord(ctx, transaction)
}

func (s *serviceImpl) UpdateTransactionById(ctx context.Context, id string, transaction models.Transaction) (string, error) {
	// TODO: input validation
	transaction.SetCategory()
	return db.UpdateTransactionRecordById(ctx, id, transaction)
}

func (s *serviceImpl) DeleteTransactionById(ctx context.Context, id string) (string, error) {
	// TODO: input validation
	return db.DeleteTransactionRecordById(ctx, id)
}
