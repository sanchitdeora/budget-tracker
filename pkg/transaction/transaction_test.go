package transaction

import (
	// "context"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mock_db "github.com/sanchitdeora/budget-tracker/db/mocks"
	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/pkg/exceptions"
	"github.com/stretchr/testify/assert"
)

const (
	TEST_ID = 	 "test-id"
	TEST_TITLE = "test-title"
)

type ServiceMocks struct {
	DB *mock_db.MockDatabase
}

func createTestTransaction(ctrl *gomock.Controller) (*Service, *ServiceMocks) {
	db := mock_db.NewMockDatabase(ctrl)
	service := NewService(&Opts{DB: db})
	return &service, &ServiceMocks{DB: db}
}

func TestGetTransactions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestTransaction(ctrl)
	
	{ 	// happy path
		expTx := &[]models.Transaction{{
			TransactionId: TEST_ID,
		}}
		mocks.DB.EXPECT().
			GetAllTransactionRecords(gomock.Any()).
			Return(expTx, nil)

		transactions, err := (*service).GetTransactions(context.Background())

		assert.Nil(t, err)
		assert.NotNil(t, transactions)
		assert.Equal(t, 1, len(*transactions))
	}

	{ 	// no transactions found
		mocks.DB.EXPECT().
			GetAllTransactionRecords(gomock.Any()).
			Return(&[]models.Transaction{}, nil)

		transactions, err := (*service).GetTransactions(context.Background())

		assert.Nil(t, err)
		assert.NotNil(t, transactions)
		assert.Equal(t, 0, len(*transactions))
	}
	
	{ 	// error found
		mocks.DB.EXPECT().
			GetAllTransactionRecords(gomock.Any()).
			Return(nil, errors.New("error found"))
		
		transactions, err := (*service).GetTransactions(context.Background())
		
		assert.NotNil(t, err)
		assert.Nil(t, transactions)
	}
}

func TestGetTransactionById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestTransaction(ctrl)
	{ 	// happy path
		expTx := &models.Transaction{
			TransactionId: TEST_ID,
		}
		mocks.DB.EXPECT().
			GetTransactionRecordById(gomock.Any(), gomock.Any()).
			Return(expTx, nil)

		transaction, err := (*service).GetTransactionById(context.Background(), TEST_ID)
		
		assert.Nil(t, err)
		assert.NotNil(t, transaction)
	}

	{ 	// validation error empty transactions id
		transaction, err := (*service).GetTransactionById(context.Background(), "")
		
		assert.Nil(t, transaction)
		assert.NotNil(t, err)
		assert.Equal(t, exceptions.ErrValidationError, err)
	}

	{ 	// not found error no transactions found
		mocks.DB.EXPECT().
			GetTransactionRecordById(gomock.Any(), gomock.Any()).
			Return(nil, exceptions.ErrTransactionNotFound)
		
		transaction, err := (*service).GetTransactionById(context.Background(), TEST_ID)
		
		assert.Nil(t, transaction)
		assert.NotNil(t, err)
		assert.Equal(t, exceptions.ErrTransactionNotFound, err)
	}
}

func TestGetTransactionByDate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestTransaction(ctrl)
	{ 	// happy path
		expTx := &[]models.Transaction{{
			TransactionId: TEST_ID,
		}}
		mocks.DB.EXPECT().
			GetAllTransactionRecordsByDateRange(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(expTx, nil)

		transactions, err := (*service).GetTransactionsByDate(context.Background(), time.UnixMilli(1688245563949), time.UnixMilli(1690923963949))
		
		assert.Nil(t, err)
		assert.NotNil(t, transactions)
		assert.Equal(t, 1, len(*transactions))
	}

	{ 	// validation error empty transactions id
		transactions, err := (*service).GetTransactionsByDate(context.Background(), time.UnixMilli(1688245563949), time.UnixMilli(1688245563948))
		
		assert.Nil(t, transactions)
		assert.NotNil(t, err)
		assert.Equal(t, exceptions.ErrValidationError, err)
	}
	
	{ 	// no transactions found
		mocks.DB.EXPECT().
			GetAllTransactionRecordsByDateRange(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&[]models.Transaction{}, nil)
		
		transactions, err := (*service).GetTransactionsByDate(context.Background(), time.UnixMilli(1688245563949), time.UnixMilli(1690923963949))
		
		assert.Nil(t, err)
		assert.NotNil(t, transactions)
		assert.Equal(t, 0, len(*transactions))
	}
}

func TestCreateTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	service, mocks := createTestTransaction(ctrl)
	{	// happy path
		mocks.DB.EXPECT().InsertTransactionRecord(gomock.Any(), gomock.Any()).Return(TEST_ID, nil)
		
		var transaction models.Transaction
		transaction.Title = TEST_TITLE
		transaction.Amount = 100
		id, err := (*service).CreateTransaction(context.Background(), transaction)
		
		assert.Equal(t, TEST_ID, id)
		assert.Nil(t, err)
	}

	{	// validation error
		
		var transaction models.Transaction
		transaction.Amount = -100
		id, err := (*service).CreateTransaction(context.Background(), transaction)
		
		assert.NotNil(t, err)
		assert.Equal(t, "", id)
	}

	{	// error found
		mocks.DB.EXPECT().InsertTransactionRecord(gomock.Any(), gomock.Any()).Return("", errors.New("error found"))
		
		var transaction models.Transaction
		transaction.Title = TEST_TITLE
		transaction.Amount = 100
		id, err := (*service).CreateTransaction(context.Background(), transaction)
		
		assert.NotNil(t, err)
		assert.Equal(t, "", id)
	}
}

func TestUpdateTransactionById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestTransaction(ctrl)
	{	// happy path
		mocks.DB.EXPECT().UpdateTransactionRecordById(gomock.Any(), TEST_ID, gomock.Any()).Return(TEST_ID, nil)
		
		var transaction models.Transaction
		transaction.Title = TEST_TITLE
		transaction.Amount = 100
		id, err := (*service).UpdateTransactionById(context.Background(), TEST_ID, transaction)
		
		assert.Equal(t, TEST_ID, id)
		assert.Nil(t, err)
	}

	{	// validation error	when no transaction_id	
		var transaction models.Transaction
		id, err := (*service).UpdateTransactionById(context.Background(), "", transaction)
		
		assert.NotNil(t, err)
		assert.Equal(t, "", id)
	}

	{	// validation error when invalid transaction provided
		var transaction models.Transaction
		transaction.Amount = -100
		id, err := (*service).UpdateTransactionById(context.Background(), TEST_ID, transaction)
		
		assert.NotNil(t, err)
		assert.Equal(t, "", id)
	}

	{	// error found
		mocks.DB.EXPECT().InsertTransactionRecord(gomock.Any(), gomock.Any()).Return("", errors.New("error found"))
		
		var transaction models.Transaction
		transaction.Title = TEST_TITLE
		transaction.Amount = 100
		id, err := (*service).CreateTransaction(context.Background(), transaction)
		
		assert.NotNil(t, err)
		assert.Equal(t, "", id)
	}
}

func TestDeleteTransactionById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestTransaction(ctrl)
	{	// happy path
		mocks.DB.EXPECT().DeleteTransactionRecordById(gomock.Any(), gomock.Any()).Return(TEST_ID, nil)
	
		id, err := (*service).DeleteTransactionById(context.Background(), TEST_ID)
		
		assert.Equal(t, TEST_ID, id)
		assert.Nil(t, err)
	}

	{	// validation error	
		id, err := (*service).DeleteTransactionById(context.Background(), "")

		assert.NotNil(t, err)
		assert.Equal(t, "", id)
	}

	{	// error found
		mocks.DB.EXPECT().DeleteTransactionRecordById(gomock.Any(), gomock.Any()).Return("", errors.New("error found"))
		
		id, err := (*service).DeleteTransactionById(context.Background(), TEST_ID)
		
		assert.NotNil(t, err)
		assert.Equal(t, "", id)
	}
}