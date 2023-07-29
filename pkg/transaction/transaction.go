package transaction

import (
	"context"
	"log"
	"time"

	"github.com/sanchitdeora/budget-tracker/db"
	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/pkg/exceptions"
)

//go:generate mockgen -destination=./mocks/mock_transaction.go -package=mock_transaction github.com/sanchitdeora/budget-tracker/pkg/transaction Service
type Service interface {
	GetTransactions(ctx context.Context) (*[]models.Transaction, error)
	GetTransactionById(ctx context.Context, id string) (*models.Transaction, error)
	GetTransactionsByDate(ctx context.Context, startDate time.Time, endDate time.Time) (*[]models.Transaction, error)
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

func (s *serviceImpl) GetTransactions(ctx context.Context) (*[]models.Transaction, error) {
	return s.DB.GetAllTransactionRecords(ctx)
}

func (s *serviceImpl) GetTransactionById(ctx context.Context, id string) (*models.Transaction, error) {
	if id == "" {
		log.Println("Missing Transaction Id")
		return nil, exceptions.ErrValidationError
	}
	return s.DB.GetTransactionRecordById(ctx, id)
}

func (s *serviceImpl) GetTransactionsByDate(ctx context.Context, startDate time.Time, endDate time.Time) (*[]models.Transaction, error) {
	if endDate.Before(startDate) {
		log.Println("end date cannot be before start date")
		return nil, exceptions.ErrValidationError
	}
	return s.DB.GetAllTransactionRecordsByDateRange(ctx, startDate, endDate)
}

func (s *serviceImpl) CreateTransaction(ctx context.Context, transaction models.Transaction) (string, error) {
	if !transaction.IsValid() {
		return "", exceptions.ErrValidationError
	}

	transaction.SetCategory()
	return s.DB.InsertTransactionRecord(ctx, transaction)
}

func (s *serviceImpl) UpdateTransactionById(ctx context.Context, id string, transaction models.Transaction) (string, error) {
	if id == "" {
		log.Println("Missing Transaction Id")
		return "", exceptions.ErrValidationError
	}

	if !transaction.IsValid() {
		return "", exceptions.ErrValidationError
	}
	transaction.SetCategory()
	return s.DB.UpdateTransactionRecordById(ctx, id, transaction)
}

func (s *serviceImpl) DeleteTransactionById(ctx context.Context, id string) (string, error) {
	if id == "" {
		log.Println("Missing Transaction Id")
		return "", exceptions.ErrValidationError
	}

	return s.DB.DeleteTransactionRecordById(ctx, id)
}
