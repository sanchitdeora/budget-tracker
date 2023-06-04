package transaction

import (
	// "context"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mock_db "github.com/sanchitdeora/budget-tracker/db/mocks"
	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/stretchr/testify/assert"
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
	mocks.DB.EXPECT().GetAllTransactionRecords(gomock.Any(), gomock.Any()).Return(nil)
	
	var transaction []models.Transaction
	err := (*service).GetTransactions(context.Background(), &transaction)
	
	assert.NoError(t, err)
	assert.Equal(t, 1, 1, "The two words should be the same.")
}

func TestGetTransactionById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestTransaction(ctrl)
	mocks.DB.EXPECT().GetTransactionRecordById(gomock.Any(), "test-id", gomock.Any()).Return(nil)
	
	var transaction models.Transaction
	err := (*service).GetTransactionById(context.Background(), "test-id", &transaction)
	
	assert.NoError(t, err)
	assert.Equal(t, 1, 1, "The two words should be the same.")
}

func TestCreateTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestTransaction(ctrl)
	mocks.DB.EXPECT().InsertTransactionRecord(gomock.Any(), gomock.Any()).Return("test-id", nil)
	
	var transaction models.Transaction
	id, err := (*service).CreateTransaction(context.Background(), transaction)
	
	assert.Equal(t, "test-id", id)
	assert.NoError(t, err)
	assert.Equal(t, 1, 1, "The two words should be the same.")
}

func TestUpdateTransactionById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestTransaction(ctrl)
	mocks.DB.EXPECT().UpdateTransactionRecordById(gomock.Any(), "test-id", gomock.Any()).Return("test-id", nil)
	
	var transaction models.Transaction
	id, err := (*service).UpdateTransactionById(context.Background(), "test-id", transaction)
	
	assert.Equal(t, "test-id", id)
	assert.NoError(t, err)
	assert.Equal(t, 1, 1, "The two words should be the same.")
}

func TestDeleteTransactionById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestTransaction(ctrl)
	mocks.DB.EXPECT().DeleteTransactionRecordById(gomock.Any(), gomock.Any()).Return("test-id", nil)
	
	id, err := (*service).DeleteTransactionById(context.Background(), gomock.Any().String())
	
	assert.Equal(t, "test-id", id)
	assert.NoError(t, err)
	assert.Equal(t, 1, 1, "The two words should be the same.")
}