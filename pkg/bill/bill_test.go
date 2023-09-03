package bill

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
	mock_transaction "github.com/sanchitdeora/budget-tracker/pkg/transaction/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	TEST_ID    			= "test-id"
	TEST_ID_2  			= "test-id-2"
	TEST_TRANSACTION_ID	= "test-transaction-id"
	TEST_TITLE 			= "test-title"
)

var (
	ErrSomeError = errors.New("error found")
)

type ServiceMocks struct {
	DB 			*mock_db.MockDatabase
	Transaction *mock_transaction.MockService
}

func createTestBill(ctrl *gomock.Controller) (Service, *ServiceMocks) {
	db := mock_db.NewMockDatabase(ctrl)
	tx := mock_transaction.NewMockService(ctrl)
	service := NewService(&Opts{DB: db, TransactionService: tx})
	return service, &ServiceMocks{DB: db, Transaction: tx}
}

func TestGetBills(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBill(ctrl)
	{	// happy path
		expBill := &[]models.Bill{{
			BillId: TEST_ID,
		}}
		mocks.DB.EXPECT().
			GetAllBillRecords(gomock.Any()).
			Return(expBill, nil)

		bills, err := service.GetBills(context.Background())

		assert.Nil(t, err)
		assert.NotNil(t, bills)
	}

	{	// no bills found
		mocks.DB.EXPECT().
			GetAllBillRecords(gomock.Any()).
			Return(&[]models.Bill{}, nil)

		bills, err := service.GetBills(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, 0, len(*bills))
	}

	{	// error found
		mocks.DB.EXPECT().
			GetAllBillRecords(gomock.Any()).
			Return(nil, ErrSomeError)

		bills, err := service.GetBills(context.Background())

		assert.Nil(t, bills)
		assert.Equal(t, ErrSomeError, err)
	}
}

func TestGetBillById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBill(ctrl)
	{ 	// happy path
		expBill := &models.Bill{
			BillId: TEST_ID,
		}
		mocks.DB.EXPECT().
			GetBillRecordById(gomock.Any(), gomock.Any()).
			Return(expBill, nil)

		bill, err := service.GetBillById(context.Background(), TEST_ID)

		assert.Nil(t, err)
		assert.NotNil(t, bill)
	}

	{ 	// validation error empty bill_id
		bill, err := service.GetBillById(context.Background(), "")
	
		assert.Nil(t, bill)
		assert.Equal(t, exceptions.ErrValidationError, err)
	}

	{ 	// not found error no bills found
		mocks.DB.EXPECT().
			GetBillRecordById(gomock.Any(), gomock.Any()).
			Return(nil, exceptions.ErrBillNotFound)

		bill, err := service.GetBillById(context.Background(), TEST_ID)

		assert.Nil(t, bill)
		assert.Equal(t, exceptions.ErrBillNotFound, err)
	}
}


// TestCreateRecurringBill

func TestCreateRecurringBill_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	date := time.Now()
	prevBill := &models.Bill{
		BillId: TEST_ID,
		Title: TEST_TITLE,
		AmountDue: 100,
		DueDate: date,
		Frequency: models.WEEKLY_FREQUENCY,
		SequenceStartId: "something",
		SequenceNumber: 1,
	}

	expBill := &models.Bill{
		Title: TEST_TITLE,
		Category: models.UNCATEGORIZED_CATEGORY,
		AmountDue: 100,
		DueDate: date.AddDate(0, 0, 7),
		Frequency: models.WEEKLY_FREQUENCY,
		SequenceStartId: "something",
		SequenceNumber: 2,
	}

	service, mocks := createTestBill(ctrl)

	mocks.DB.EXPECT().
		InsertBillRecord(gomock.Any(), expBill).
		Return(TEST_ID_2, nil)

	id, err := service.CreateRecurringBill(context.Background(), prevBill)

	assert.Equal(t, TEST_ID_2, id)
	assert.Nil(t, err)
}

func TestCreateRecurringBill_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBill(ctrl)
	{	// happy path
		mocks.DB.EXPECT().
			InsertBillRecord(gomock.Any(), gomock.Any()).
			Return(TEST_ID, nil)

		bill := &models.Bill{
			Title: TEST_TITLE,
			AmountDue: 100,
		}
		id, err := service.CreateRecurringBill(context.Background(), bill)

		assert.Equal(t, TEST_ID, id)
		assert.Nil(t, err)
	}
	{	// validation error

		bill := &models.Bill{
			AmountDue: -100,
		}
		id, err := service.CreateRecurringBill(context.Background(), bill)

		assert.Equal(t, err, exceptions.ErrValidationError)
		assert.Equal(t, "", id)
	}
	{	// error while inserting bill
		mocks.DB.EXPECT().
			InsertBillRecord(gomock.Any(), gomock.Any()).
			Return("", ErrSomeError)

		bill := &models.Bill{
			Title: TEST_TITLE,
			AmountDue: 100,
		}
		id, err := service.CreateRecurringBill(context.Background(), bill)
	
		assert.Equal(t, "", id)
		assert.Equal(t, ErrSomeError, err)
	}
}

func TestCreateBillByUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBill(ctrl)
	{	// happy path
		mocks.DB.EXPECT().
			InsertBillRecord(gomock.Any(), gomock.Any()).
			Return(TEST_ID, nil)

		bill := &models.Bill{
			Title: TEST_TITLE,
			AmountDue: 100,
		}
		id, err := service.CreateBillByUser(context.Background(), bill)

		assert.Equal(t, TEST_ID, id)
		assert.Nil(t, err)
	}

	{	// validation error
		bill := &models.Bill{
			AmountDue: -100,
		}
		id, err := service.CreateBillByUser(context.Background(), bill)

		assert.Equal(t, "", id)
		assert.Equal(t, exceptions.ErrValidationError, err)
	}

	{	// error found while inserting bill
		mocks.DB.EXPECT().
			InsertBillRecord(gomock.Any(), gomock.Any()).
			Return("", ErrSomeError)

		bill := &models.Bill{
			Title: TEST_TITLE,
			AmountDue: 100,
		}
		id, err := service.CreateBillByUser(context.Background(), bill)

		assert.Equal(t, "", id)
		assert.Equal(t, ErrSomeError, err)
	}
}

func TestUpdateBillById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBill(ctrl)
	{	// happy path
		mocks.DB.EXPECT().
			UpdateBillRecordById(gomock.Any(), TEST_ID, gomock.Any()).
			Return(TEST_ID, nil)

		bill := &models.Bill{
			Title: TEST_TITLE,
			AmountDue: 100,
		}
		id, err := service.UpdateBillById(context.Background(), TEST_ID, bill)

		assert.Equal(t, TEST_ID, id)
		assert.Nil(t, err)
	}

	{	// validation error	when no bill_id	
		id, err := service.UpdateBillById(context.Background(), "", &models.Bill{})
		
		assert.Equal(t, "", id)
		assert.Equal(t, exceptions.ErrValidationError, err)
	}

	{	// validation error when invalid bill provided
		bill := &models.Bill{
			AmountDue: -100,
		}
		id, err := service.UpdateBillById(context.Background(), TEST_ID, bill)

		assert.Equal(t, "", id)
		assert.Equal(t, exceptions.ErrValidationError, err)
	}

	{	// error found
		mocks.DB.EXPECT().
			UpdateBillRecordById(gomock.Any(), TEST_ID, gomock.Any()).
			Return("", ErrSomeError)

		bill := &models.Bill{
			Title: TEST_TITLE,
			AmountDue: 100,
		}
		id, err := service.UpdateBillById(context.Background(), TEST_ID, bill)

		assert.Equal(t, "", id)
		assert.Equal(t, ErrSomeError, err)
	}
}


// TestUpdateBillIsPaid

func TestUpdateBillIsPaid_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBill(ctrl)

	{	// bill is unpaid && transaction id does not exist
		bill := &models.Bill{
			BillId: TEST_ID,
			Title: TEST_TITLE,
			AmountDue: 100,
			IsPaid: false,
			Frequency: models.MONTHLY_FREQUENCY,
		}

		mocks.DB.EXPECT().
			GetBillRecordById(gomock.Any(), TEST_ID).
			Return(bill, nil)
		mocks.Transaction.EXPECT().
			CreateTransaction(gomock.Any(), gomock.Any()).
			Return(TEST_TRANSACTION_ID, nil)
		mocks.DB.EXPECT().
			InsertBillRecord(gomock.Any(), gomock.Any()).
			Return(TEST_ID_2, nil)
		mocks.DB.EXPECT().
			UpdateBillRecordById(gomock.Any(), TEST_ID, gomock.Any()).
			Return(TEST_ID, nil)

		id, err := service.UpdateBillIsPaid(context.Background(), TEST_ID)
		assert.Nil(t, err)
		assert.Equal(t, TEST_ID, id)
	}

	{	// bill is unpaid && transaction id exists
		bill := &models.Bill{
			BillId: TEST_ID,
			Title: TEST_TITLE,
			AmountDue: 100,
			IsPaid: false,
			Frequency: models.MONTHLY_FREQUENCY,
			TransactionId: TEST_TRANSACTION_ID,
		}

		expTx := &models.Transaction{
			TransactionId: TEST_TRANSACTION_ID,
		}

		mocks.DB.EXPECT().
			GetBillRecordById(gomock.Any(), TEST_ID).
			Return(bill, nil)
		mocks.DB.EXPECT().
			InsertBillRecord(gomock.Any(), gomock.Any()).
			Return(TEST_ID_2, nil)
		mocks.Transaction.EXPECT().
			GetTransactionById(gomock.Any(), gomock.Any()).
			Return(expTx, nil)
		mocks.Transaction.EXPECT().
			UpdateTransactionById(gomock.Any(), TEST_TRANSACTION_ID, gomock.Any()).
			Return(TEST_TRANSACTION_ID, nil)
		mocks.DB.EXPECT().
			UpdateBillRecordById(gomock.Any(), TEST_ID, gomock.Any()).
			Return(TEST_ID, nil)

		id, err := service.UpdateBillIsPaid(context.Background(), TEST_ID)
		assert.Nil(t, err)
		assert.Equal(t, TEST_ID, id)
	}

}

func TestUpdateBillIsPaid_DoNothing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBill(ctrl)
	{	// bill is already paid && transaction created, do nothing
		expBill := &models.Bill{
			BillId: TEST_ID,
			Title: TEST_TITLE,
			AmountDue: 100,
			IsPaid: true,
			Frequency: models.ONCE_FREQUENCY,
			TransactionId: TEST_TRANSACTION_ID,
		}

		mocks.DB.EXPECT().
			GetBillRecordById(gomock.Any(), TEST_ID).
			Return(expBill, nil)

		id, err := service.UpdateBillIsPaid(context.Background(), TEST_ID)

		assert.Nil(t, err)
		assert.Equal(t, TEST_ID, id)
	}

	{	// bill is unpaid && frequency is one, do nothing
		expBill := &models.Bill{
			BillId: TEST_ID,
			Title: TEST_TITLE,
			AmountDue: 100,
			IsPaid: false,
			Frequency: models.ONCE_FREQUENCY,
			NextSequenceId: "",
		}

		mocks.DB.EXPECT().
			GetBillRecordById(gomock.Any(), TEST_ID).
			Return(expBill, nil)
		mocks.Transaction.EXPECT().
			CreateTransaction(gomock.Any(), gomock.Any()).
			Return(TEST_TRANSACTION_ID, nil)
		mocks.DB.EXPECT().
			UpdateBillRecordById(gomock.Any(), TEST_ID, gomock.Any()).
			Return(TEST_ID, nil)

		id, err := service.UpdateBillIsPaid(context.Background(), TEST_ID)
		assert.Nil(t, err)
		assert.Equal(t, TEST_ID, id)
	}

	{	// bill is unpaid && Next Sequence Id already exists
		expBill := &models.Bill{
			BillId: TEST_ID,
			Title: TEST_TITLE,
			AmountDue: 100,
			IsPaid: false,
			Frequency: models.MONTHLY_FREQUENCY,
			NextSequenceId: TEST_ID_2,
		}

		mocks.DB.EXPECT().
			GetBillRecordById(gomock.Any(), TEST_ID).
			Return(expBill, nil)
		mocks.Transaction.EXPECT().
			CreateTransaction(gomock.Any(), gomock.Any()).
			Return(TEST_TRANSACTION_ID, nil)
		mocks.DB.EXPECT().
			UpdateBillRecordById(gomock.Any(), TEST_ID, gomock.Any()).
			Return(TEST_ID, nil)

		id, err := service.UpdateBillIsPaid(context.Background(), TEST_ID)
		assert.Nil(t, err)
		assert.Equal(t, TEST_ID, id)
	}

}

func TestUpdateBillIsPaid_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBill(ctrl)

	{	// validation error	when no bill_id	
		id, err := service.UpdateBillIsPaid(context.Background(), "")

		assert.Equal(t, "", id)
		assert.Equal(t, exceptions.ErrValidationError, err)
	}

	{	// error found while fetching bill by id
		mocks.DB.EXPECT().
			GetBillRecordById(context.Background(), TEST_ID).
			Return(nil, ErrSomeError)

		id, err := service.UpdateBillIsPaid(context.Background(), TEST_ID)

		assert.Equal(t, "", id)
		assert.Equal(t, ErrSomeError, err)
	}

	{	// error found while creating transaction	
		expBill := &models.Bill{
			BillId: TEST_ID,
			Title: TEST_TITLE,
			AmountDue: 100,
			IsPaid: false,
			Frequency: models.ONCE_FREQUENCY,
		}

		mocks.DB.EXPECT().
			GetBillRecordById(gomock.Any(), TEST_ID).
			Return(expBill, nil)
		mocks.Transaction.EXPECT().
			CreateTransaction(gomock.Any(), gomock.Any()).
			Return("", ErrSomeError)

		id, err := service.UpdateBillIsPaid(context.Background(), TEST_ID)

		assert.Equal(t, "", id)
		assert.Equal(t, ErrSomeError, err)
	}

	{	// error found while creating new bill	
		expBill := &models.Bill{
			BillId: TEST_ID,
			Title: TEST_TITLE,
			AmountDue: 100,
			IsPaid: false,
			Frequency: models.MONTHLY_FREQUENCY,
		}

		mocks.DB.EXPECT().
			GetBillRecordById(gomock.Any(), TEST_ID).
			Return(expBill, nil)
		mocks.DB.EXPECT().
			InsertBillRecord(gomock.Any(), gomock.Any()).
			Return("", ErrSomeError)

		id, err := service.UpdateBillIsPaid(context.Background(), TEST_ID)

		assert.Equal(t, "", id)
		assert.Equal(t, ErrSomeError, err)
	}

	{	// error found while fetching transaction	
		bill := &models.Bill{
			BillId: TEST_ID,
			Title: TEST_TITLE,
			AmountDue: 100,
			IsPaid: false,
			Frequency: models.MONTHLY_FREQUENCY,
			TransactionId: TEST_TRANSACTION_ID,
		}

		mocks.DB.EXPECT().
			GetBillRecordById(gomock.Any(), TEST_ID).
			Return(bill, nil)
		mocks.DB.EXPECT().
			InsertBillRecord(gomock.Any(), gomock.Any()).
			Return(TEST_ID_2, nil)
		mocks.Transaction.EXPECT().
			GetTransactionById(gomock.Any(), gomock.Any()).
			Return(nil, ErrSomeError)

		id, err := service.UpdateBillIsPaid(context.Background(), TEST_ID)
		assert.Equal(t, "", id)
		assert.Equal(t, ErrSomeError, err)
	}

	{	// error found while updating transaction	
		bill := &models.Bill{
			BillId: TEST_ID,
			Title: TEST_TITLE,
			AmountDue: 100,
			IsPaid: false,
			Frequency: models.MONTHLY_FREQUENCY,
			TransactionId: TEST_TRANSACTION_ID,
		}

		tx := &models.Transaction{
			TransactionId: TEST_TRANSACTION_ID,
		}

		mocks.DB.EXPECT().
			GetBillRecordById(gomock.Any(), TEST_ID).
			Return(bill, nil)
		mocks.DB.EXPECT().
			InsertBillRecord(gomock.Any(), gomock.Any()).
			Return(TEST_ID_2, nil)
		mocks.Transaction.EXPECT().
			GetTransactionById(gomock.Any(), TEST_TRANSACTION_ID).
			Return(tx, nil)
		mocks.Transaction.EXPECT().
			UpdateTransactionById(gomock.Any(), TEST_TRANSACTION_ID, gomock.Any()).
			Return("", ErrSomeError)

		id, err := service.UpdateBillIsPaid(context.Background(), TEST_ID)
		assert.Equal(t, "", id)
		assert.Equal(t, ErrSomeError, err)
	}
}


// TestUpdateBillIsUnpaid

func TestUpdateBillIsUnpaid_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBill(ctrl)

	bill := &models.Bill{
		BillId: TEST_ID,
		AmountDue: 100,
		IsPaid: true,
		DatePaid: time.Now(),
		TransactionId: TEST_TRANSACTION_ID,
	}

	expBill := &models.Bill{
		BillId: TEST_ID,
		AmountDue: 100,
		IsPaid: false,
	}

	mocks.DB.EXPECT().
		GetBillRecordById(gomock.Any(), TEST_ID).
		Return(bill, nil)
	mocks.Transaction.EXPECT().
		DeleteTransactionById(gomock.Any(), TEST_TRANSACTION_ID).
		Return(TEST_TRANSACTION_ID, nil)
	mocks.DB.EXPECT().
		UpdateBillRecordById(gomock.Any(), TEST_ID, expBill).
		Return(TEST_ID, nil)

	id, err := service.UpdateBillIsUnpaid(context.Background(), TEST_ID)

	assert.Equal(t, TEST_ID, id)
	assert.Nil(t, err)
}

func TestUpdateBillIsUnpaid_ReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBill(ctrl)

	{	// validation error	when no bill_id	
		id, err := service.UpdateBillIsUnpaid(context.Background(), "")

		assert.Equal(t, "", id)
		assert.Equal(t, exceptions.ErrValidationError, err)
	}

	{	// error found while fetching bill
		mocks.DB.EXPECT().
			GetBillRecordById(gomock.Any(), TEST_ID).
			Return(nil, ErrSomeError)

		id, err := service.UpdateBillIsUnpaid(context.Background(), TEST_ID)

		assert.Equal(t, "", id)
		assert.Equal(t, ErrSomeError, err)
	}

	{	// error found while deleting transaction by id

		bill := &models.Bill{
			BillId: TEST_ID,
			AmountDue: 100,
			IsPaid: true,
			DatePaid: time.Now(),
			TransactionId: TEST_TRANSACTION_ID,
		}

		mocks.DB.EXPECT().
			GetBillRecordById(gomock.Any(), TEST_ID).
			Return(bill, nil)
		mocks.Transaction.EXPECT().
			DeleteTransactionById(gomock.Any(), TEST_TRANSACTION_ID).
			Return("", ErrSomeError)

		id, err := service.UpdateBillIsUnpaid(context.Background(), TEST_ID)

		assert.Equal(t, "", id)
		assert.Equal(t, ErrSomeError, err)
	}

	{	// error found while updating bill

		bill := &models.Bill{
			BillId: TEST_ID,
			AmountDue: 100,
			IsPaid: true,
			DatePaid: time.Now(),
		}

		mocks.DB.EXPECT().
			GetBillRecordById(gomock.Any(), TEST_ID).
			Return(bill, nil)
		mocks.DB.EXPECT().
			UpdateBillRecordById(gomock.Any(), TEST_ID, gomock.Any()).
			Return("", ErrSomeError)

		id, err := service.UpdateBillIsUnpaid(context.Background(), TEST_ID)

		assert.Equal(t, "", id)
		assert.Equal(t, ErrSomeError, err)
	}
}

func TestDeleteBillById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBill(ctrl)
	{	// happy path
		mocks.DB.EXPECT().
			DeleteBillRecordById(gomock.Any(), gomock.Any()).
			Return(TEST_ID, nil)
	
		id, err := service.DeleteBillById(context.Background(), TEST_ID)

		assert.Equal(t, TEST_ID, id)
		assert.Nil(t, err)
	}

	{	// validation error	
		id, err := service.DeleteBillById(context.Background(), "")

		assert.Equal(t, "", id)
		assert.Equal(t, exceptions.ErrValidationError, err)
	}

	{	// error found
		mocks.DB.EXPECT().
			DeleteBillRecordById(gomock.Any(), gomock.Any()).
			Return("", ErrSomeError)

		id, err := service.DeleteBillById(context.Background(), TEST_ID)

		assert.Equal(t, "", id)
		assert.Equal(t, ErrSomeError, err)
	}
}

// TestBillMaintainer

func TestBillMaintainer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBill(ctrl)

	timeNow := time.Now()
	{	// maintainer lifecycle
		bills := &[]models.Bill{
			{
				Frequency: models.ONCE_FREQUENCY,
			},
			{
				Frequency: models.MONTHLY_FREQUENCY,
				NextSequenceId: "test-next-sequence-id",
			},
			{
				Frequency: models.MONTHLY_FREQUENCY,
			},
			{
				Frequency: models.MONTHLY_FREQUENCY,
				DueDate: timeNow.AddDate(0, 0, 1),
			},
			{
				BillId: TEST_ID+"-fail",
				Title: TEST_TITLE,
				Frequency: models.MONTHLY_FREQUENCY,
				DueDate: timeNow.AddDate(0, 0, -1),
			},
			{
				BillId: TEST_ID,
				Title: TEST_TITLE,
				Frequency: models.MONTHLY_FREQUENCY,
				DueDate: timeNow.AddDate(0, 0, -1),
			},
		}

		expBill := &models.Bill{
			Title: TEST_TITLE,
			Category: models.UNCATEGORIZED_CATEGORY,
			Frequency: models.MONTHLY_FREQUENCY,
			DueDate: timeNow.AddDate(0, 1, -1),
			SequenceStartId: TEST_ID,
			SequenceNumber: 1,
		}

		mocks.DB.EXPECT().
			GetAllBillRecords(gomock.Any()).
			Return(bills, nil)
		mocks.DB.EXPECT().
			InsertBillRecord(gomock.Any(), gomock.Any()).
			Return("", ErrSomeError).
			Times(1)
		mocks.DB.EXPECT().
			InsertBillRecord(gomock.Any(), expBill).
			Return(TEST_ID_2, nil)
		mocks.DB.EXPECT().
			UpdateBillRecordById(gomock.Any(), TEST_ID, gomock.Any()).
			Return(TEST_ID, nil)

		service.BillMaintainer(context.Background())
	}

	{	// error found while fetching bill, close maintainer
		mocks.DB.EXPECT().
			GetAllBillRecords(gomock.Any()).
			Return(nil, ErrSomeError)
		
		service.BillMaintainer(context.Background())
	}
}
