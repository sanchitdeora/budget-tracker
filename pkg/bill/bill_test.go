package bill

import (
	// "context"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mock_db "github.com/sanchitdeora/budget-tracker/db/mocks"
	"github.com/sanchitdeora/budget-tracker/models"
	mock_transaction "github.com/sanchitdeora/budget-tracker/pkg/transaction/mocks"
	"github.com/stretchr/testify/assert"
)

type ServiceMocks struct {
	DB 			*mock_db.MockDatabase
	Transaction *mock_transaction.MockService
}

func createTestBill(ctrl *gomock.Controller) (*Service, *ServiceMocks) {
	db := mock_db.NewMockDatabase(ctrl)
	tx := mock_transaction.NewMockService(ctrl)
	service := NewService(&Opts{DB: db, TransactionService: tx})
	return &service, &ServiceMocks{DB: db, Transaction: tx}
}

func TestGetBills(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBill(ctrl)
	mocks.DB.EXPECT().GetAllBillRecords(gomock.Any(), gomock.Any()).Return(nil)
	
	var bill []models.Bill
	err := (*service).GetBills(context.Background(), &bill)
	
	assert.NoError(t, err)
	assert.Equal(t, 1, 1, "The two words should be the same.")
}

func TestGetBillById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBill(ctrl)
	mocks.DB.EXPECT().GetBillRecordById(gomock.Any(), "test-id", gomock.Any()).Return(nil)
	
	var bill models.Bill
	err := (*service).GetBillById(context.Background(), "test-id", &bill)
	
	assert.NoError(t, err)
	assert.Equal(t, 1, 1, "The two words should be the same.")
}

func TestCreateBill(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBill(ctrl)
	mocks.DB.EXPECT().InsertBillRecord(gomock.Any(), gomock.Any()).Return("test-id", nil)
	
	var bill models.Bill
	id, err := (*service).CreateBill(context.Background(), bill)
	
	assert.Equal(t, "test-id", id)
	assert.NoError(t, err)
	assert.Equal(t, 1, 1, "The two words should be the same.")
}

func TestUpdateBillById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBill(ctrl)
	mocks.DB.EXPECT().UpdateBillRecordById(gomock.Any(), "test-id", gomock.Any()).Return("test-id", nil)
	
	var bill models.Bill
	id, err := (*service).UpdateBillById(context.Background(), "test-id", bill)
	
	assert.Equal(t, "test-id", id)
	assert.NoError(t, err)
	assert.Equal(t, 1, 1, "The two words should be the same.")
}

func TestUpdateBillIsPaid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBill(ctrl)

	// mocks.Transaction.EXPECT().CreateTransaction(gomock.Any(), gomock.Any()).Return("test-id", nil)
	// mocks.DB.EXPECT().GetBillRecordById(gomock.Any(), "test-id", gomock.Any())
	// mocks.DB.EXPECT().InsertBillRecord(gomock.Any(), gomock.Any()).Return("test-id", nil)
	// mocks.DB.EXPECT().UpdateBillRecordIsPaid(gomock.Any(), "test-id", gomock.Any()).Return("test-id", nil)
	
	// assert.Equal(t, "test-id", id)

	{
		// bill is already paid && frequency is one, do nothing
		
		var bill models.Bill
		// mocks.Transaction.EXPECT().CreateTransaction(gomock.Any(), gomock.Any()).Return("test-id", nil)
		// mocks.DB.EXPECT().InsertBillRecord(gomock.Any(), gomock.Any()).Return("test-id", nil)
		mocks.DB.EXPECT().UpdateBillRecordIsPaid(gomock.Any(), "test-id", gomock.Any()).Return("test-id", nil)
		
		mocks.DB.EXPECT().GetBillRecordById(gomock.Any(), "test-id", &bill).DoAndReturn(func(ctx context.Context, id string, bill *models.Bill) (*models.Bill, error) {
			*bill = models.Bill{
				BillId: "test-id",
				Title: "test-title",
				AmountDue: 100,
				IsPaid: true,
				Frequency: "once",
			}
			return bill, nil
		})
		id, err := (*service).UpdateBillIsPaid(context.Background(), "test-id")
		assert.NoError(t, err)
		assert.Equal(t, "test-id", id)
	}
	{
		// bill is unpaid && frequency is one, do nothing
		
		var bill models.Bill
		mocks.Transaction.EXPECT().CreateTransaction(gomock.Any(), gomock.Any()).Return("test-id", nil)
		// mocks.DB.EXPECT().InsertBillRecord(gomock.Any(), gomock.Any()).Return("test-id", nil)
		mocks.DB.EXPECT().UpdateBillRecordIsPaid(gomock.Any(), "test-id", gomock.Any()).Return("test-id", nil)
		
		mocks.DB.EXPECT().GetBillRecordById(gomock.Any(), "test-id", &bill).DoAndReturn(func(ctx context.Context, id string, bill *models.Bill) (*models.Bill, error) {
			*bill = models.Bill{
				BillId: "test-id",
				Title: "test-title",
				AmountDue: 100,
				IsPaid: false,
				Frequency: "once",
			}
			return bill, nil
		})
		id, err := (*service).UpdateBillIsPaid(context.Background(), "test-id")
		assert.NoError(t, err)
		assert.Equal(t, "test-id", id)
	}
	{
		// bill is unpaid && frequency is more, do nothing
		
		var bill models.Bill
		mocks.Transaction.EXPECT().CreateTransaction(gomock.Any(), gomock.Any()).Return("test-id", nil)
		mocks.DB.EXPECT().InsertBillRecord(gomock.Any(), gomock.Any()).Return("test-id", nil)
		mocks.DB.EXPECT().UpdateBillRecordIsPaid(gomock.Any(), "test-id", gomock.Any()).Return("test-id", nil)
		
		mocks.DB.EXPECT().GetBillRecordById(gomock.Any(), "test-id", &bill).DoAndReturn(func(ctx context.Context, id string, bill *models.Bill) (*models.Bill, error) {
			*bill = models.Bill{
				BillId: "test-id",
				Title: "test-title",
				AmountDue: 100,
				IsPaid: false,
				Frequency: "monthly",
			}
			return bill, nil
		})
		id, err := (*service).UpdateBillIsPaid(context.Background(), "test-id")
		assert.NoError(t, err)
		assert.Equal(t, "test-id", id)
	}

	assert.Equal(t, 1, 1, "The two words should be the same.")
}

func TestUpdateBillIsUnpaid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBill(ctrl)
	
	mocks.DB.EXPECT().UpdateBillRecordIsUnpaid(gomock.Any(), "test-id").Return("test-id", nil)
	
	id, err := (*service).UpdateBillIsUnpaid(context.Background(), "test-id")
	
	assert.Equal(t, "test-id", id)
	assert.NoError(t, err)
	assert.Equal(t, 1, 1, "The two words should be the same.")
}

func TestDeleteBillById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBill(ctrl)
	mocks.DB.EXPECT().DeleteBillRecordById(gomock.Any(), gomock.Any()).Return("test-id", nil)
	
	id, err := (*service).DeleteBillById(context.Background(), gomock.Any().String())
	
	assert.Equal(t, "test-id", id)
	assert.NoError(t, err)
	assert.Equal(t, 1, 1, "The two words should be the same.")
}