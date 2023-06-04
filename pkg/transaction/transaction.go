package transaction

import (
	"context"

	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/models"
)

//go:generate mockgen -destination=../mocks/mock_transaction.go -package=mocks github.com/sanchitdeora/budget-tracker/pkg/transaction Service
type Service interface {
	GetTransactions(ctx context.Context, transactions *[]models.Transaction) (error)
	GetTransactionById(ctx context.Context, id string, transaction *models.Transaction) (error)
	CreateTransaction(ctx context.Context, transaction models.Transaction) (string, error)
	UpdateTransactionById(ctx context.Context, id string, transaction models.Transaction) (string, error)
	DeleteTransactionById(ctx context.Context, id string) (string, error)
}

type Opts struct {
	DB db.Database
}

type serviceImpl struct {
	*Opts
}

func NewService(opts *Opts) Service {
	return &serviceImpl{Opts: opts}
}

func (s *serviceImpl) GetTransactions(ctx context.Context, transactions *[]models.Transaction) (error) {
	// TODO: input validation
	return s.DB.GetAllTransactionRecords(ctx, transactions)
}

func (s *serviceImpl) GetTransactionById(ctx context.Context, id string, transaction *models.Transaction) (error) {
	// TODO: input validation
	return s.DB.GetTransactionRecordById(ctx, id, transaction)
}

func (s *serviceImpl) CreateTransaction(ctx context.Context, transaction models.Transaction) (string, error) {
	// TODO: input validation
	transaction.SetCategory()
	return s.DB.InsertTransactionRecord(ctx, transaction)
}

func (s *serviceImpl) UpdateTransactionById(ctx context.Context, id string, transaction models.Transaction) (string, error) {
	// TODO: input validation
	transaction.SetCategory()
	return s.DB.UpdateTransactionRecordById(ctx, id, transaction)
}

func (s *serviceImpl) DeleteTransactionById(ctx context.Context, id string) (string, error) {
	// TODO: input validation
	return s.DB.DeleteTransactionRecordById(ctx, id)
}
