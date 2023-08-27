package bill

import (
	// "context"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mock_db "github.com/sanchitdeora/budget-tracker/db/mocks"
	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/pkg/exceptions"
	mock_transaction "github.com/sanchitdeora/budget-tracker/pkg/transaction/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	TEST_ID = 	 "test-id"
	TEST_TITLE = "test-title"
)

var (
	ErrSomeError = errors.New("error found")
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
	{	// happy path
		expBill := &[]models.Bill{{
			BillId: TEST_ID,
		}}
		mocks.DB.EXPECT().
			GetAllBillRecords(gomock.Any()).
			Return(expBill, nil)

		bills, err := (*service).GetBills(context.Background())

		assert.Nil(t, err)
		assert.NotNil(t, bills)
	}

	{	// no bills found
		mocks.DB.EXPECT().
			GetAllBillRecords(gomock.Any()).
			Return(&[]models.Bill{}, nil)

		bills, err := (*service).GetBills(context.Background())

		assert.Nil(t, err)
		assert.NotNil(t, bills)
		assert.Equal(t, 0, len(*bills))
	}

	{	// error found
		mocks.DB.EXPECT().
			GetAllBillRecords(gomock.Any()).
			Return(nil, ErrSomeError)

		bills, err := (*service).GetBills(context.Background())

		assert.NotNil(t, err)
		assert.Nil(t, bills)
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

		bill, err := (*service).GetBillById(context.Background(), TEST_ID)

		assert.Nil(t, err)
		assert.NotNil(t, bill)
	}

	{ 	// validation error empty bill_id
		bill, err := (*service).GetBillById(context.Background(), "")
	
		assert.Nil(t, bill)
		assert.NotNil(t, err)
		assert.Equal(t, exceptions.ErrValidationError, err)
	}

	{ 	// not found error no bills found
		mocks.DB.EXPECT().
			GetBillRecordById(gomock.Any(), gomock.Any()).
			Return(nil, exceptions.ErrBillNotFound)

		bill, err := (*service).GetBillById(context.Background(), TEST_ID)

		assert.Nil(t, bill)
		assert.NotNil(t, err)
		assert.Equal(t, exceptions.ErrBillNotFound, err)
	}
}

func TestCreateBill(t *testing.T) {
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
		id, err := (*service).CreateBill(context.Background(), bill)

		assert.Equal(t, TEST_ID, id)
		assert.Nil(t, err)
	}

	{	// validation error

		bill := &models.Bill{
			AmountDue: -100,
		}
		id, err := (*service).CreateBill(context.Background(), bill)

		assert.Equal(t, err, exceptions.ErrValidationError)
		assert.Equal(t, "", id)
	}

	{	// happy path
		mocks.DB.EXPECT().
			InsertBillRecord(gomock.Any(), gomock.Any()).
			Return("", ErrSomeError)

		bill := &models.Bill{
			Title: TEST_TITLE,
			AmountDue: 100,
		}
		id, err := (*service).CreateBill(context.Background(), bill)
	
		assert.NotNil(t, err)
		assert.Equal(t, "", id)
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
		id, err := (*service).CreateBillByUser(context.Background(), bill)

		assert.Equal(t, TEST_ID, id)
		assert.Nil(t, err)
	}

	{	// validation error
		bill := &models.Bill{
			AmountDue: -100,
		}
		id, err := (*service).CreateBillByUser(context.Background(), bill)

		assert.NotNil(t, err)
		assert.Equal(t, "", id)
	}

	{	// happy path
		mocks.DB.EXPECT().
			InsertBillRecord(gomock.Any(), gomock.Any()).
			Return("", ErrSomeError)

		bill := &models.Bill{
			Title: TEST_TITLE,
			AmountDue: 100,
		}
		id, err := (*service).CreateBillByUser(context.Background(), bill)

		assert.NotNil(t, err)
		assert.Equal(t, "", id)
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
		id, err := (*service).UpdateBillById(context.Background(), TEST_ID, bill)

		assert.Equal(t, TEST_ID, id)
		assert.Nil(t, err)
	}

	{	// validation error	when no bill_id	
		id, err := (*service).UpdateBillById(context.Background(), "", &models.Bill{})

		assert.NotNil(t, err)
		assert.Equal(t, "", id)
	}

	{	// validation error when invalid bill provided
		bill := &models.Bill{
			AmountDue: -100,
		}
		id, err := (*service).UpdateBillById(context.Background(), TEST_ID, bill)

		assert.NotNil(t, err)
		assert.Equal(t, "", id)
	}

	{	// error found
		mocks.DB.EXPECT().
			UpdateBillRecordById(gomock.Any(), TEST_ID, gomock.Any()).
			Return("", ErrSomeError)

		bill := &models.Bill{
			Title: TEST_TITLE,
			AmountDue: 100,
		}
		id, err := (*service).UpdateBillById(context.Background(), TEST_ID, bill)

		assert.NotNil(t, err)
		assert.Equal(t, "", id)
	}
}

func TestUpdateBillIsPaid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBill(ctrl)

	{	// validation error	when no bill_id	
		id, err := (*service).UpdateBillIsPaid(context.Background(), "")

		assert.NotNil(t, err)
		assert.Equal(t, "", id)
	}

	{	// error found while fetching bill by id
		mocks.DB.EXPECT().
			GetBillRecordById(context.Background(), TEST_ID).
			Return(nil, ErrSomeError)

		id, err := (*service).UpdateBillIsPaid(context.Background(), TEST_ID)

		assert.NotNil(t, err)
		assert.Equal(t, "", id)
	}

	{	// bill is already paid && frequency is one, do nothing
		expBill := &models.Bill{
			BillId: TEST_ID,
			Title: TEST_TITLE,
			AmountDue: 100,
			IsPaid: true,
			Frequency: models.ONCE_FREQUENCY,
		}

		mocks.DB.EXPECT().
			UpdateBillRecordIsPaid(gomock.Any(), "test-id", gomock.Any()).
			Return("test-id", nil)
		mocks.DB.EXPECT().
			GetBillRecordById(gomock.Any(), "test-id").
			Return(expBill, nil)

		id, err := (*service).UpdateBillIsPaid(context.Background(), "test-id")

		assert.Nil(t, err)
		assert.Equal(t, "test-id", id)
	}

	{	// bill is unpaid && frequency is one, do nothing
		expBill := &models.Bill{
			BillId: TEST_ID,
			Title: TEST_TITLE,
			AmountDue: 100,
			IsPaid: false,
			Frequency: models.ONCE_FREQUENCY,
		}

		mocks.DB.EXPECT().
			GetBillRecordById(gomock.Any(), "test-id").
			Return(expBill, nil)
		mocks.Transaction.EXPECT().
			CreateTransaction(gomock.Any(), gomock.Any()).
			Return("test-id", nil)
		mocks.DB.EXPECT().
			UpdateBillRecordIsPaid(gomock.Any(), "test-id", gomock.Any()).
			Return("test-id", nil)

		id, err := (*service).UpdateBillIsPaid(context.Background(), "test-id")
		assert.Nil(t, err)
		assert.Equal(t, "test-id", id)
	}

	{	// bill is unpaid && frequency is more, do nothing
		expBill := &models.Bill{
			BillId: TEST_ID,
			Title: TEST_TITLE,
			AmountDue: 100,
			IsPaid: false,
			Frequency: models.MONTHLY_FREQUENCY,
		}

		mocks.DB.EXPECT().
			GetBillRecordById(gomock.Any(), "test-id").
			Return(expBill, nil)
		mocks.Transaction.EXPECT().
			CreateTransaction(gomock.Any(), gomock.Any()).
			Return("test-id", nil)
		mocks.DB.EXPECT().
			InsertBillRecord(gomock.Any(), gomock.Any()).
			Return("test-id", nil)
		mocks.DB.EXPECT().
			UpdateBillRecordIsPaid(gomock.Any(), "test-id", gomock.Any()).
			Return("test-id", nil)

		id, err := (*service).UpdateBillIsPaid(context.Background(), "test-id")
		assert.NoError(t, err)
		assert.Equal(t, "test-id", id)
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
			GetBillRecordById(gomock.Any(), "test-id").
			Return(expBill, nil)
		mocks.Transaction.EXPECT().
			CreateTransaction(gomock.Any(), gomock.Any()).
			Return("", ErrSomeError)

		id, err := (*service).UpdateBillIsPaid(context.Background(), "test-id")
		assert.Equal(t, "", id)
		assert.NotNil(t, err)
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
			GetBillRecordById(gomock.Any(), "test-id").
			Return(expBill, nil)
		mocks.Transaction.EXPECT().
			CreateTransaction(gomock.Any(), gomock.Any()).
			Return("test-id", nil)
		mocks.DB.EXPECT().
			InsertBillRecord(gomock.Any(), gomock.Any()).
			Return("", ErrSomeError)

		id, err := (*service).UpdateBillIsPaid(context.Background(), "test-id")
		assert.Equal(t, "", id)
		assert.NotNil(t, err)
	}
}

func TestUpdateBillIsUnpaid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBill(ctrl)
	{	// happy path
		mocks.DB.EXPECT().
			UpdateBillRecordIsUnpaid(gomock.Any(), TEST_ID).
			Return(TEST_ID, nil)

		id, err := (*service).UpdateBillIsUnpaid(context.Background(), TEST_ID)

		assert.Equal(t, TEST_ID, id)
		assert.Nil(t, err)
	}

	{	// validation error	when no bill_id	
		id, err := (*service).UpdateBillIsUnpaid(context.Background(), "")

		assert.NotNil(t, err)
		assert.Equal(t, "", id)
	}

	{	// error found
		mocks.DB.EXPECT().
			UpdateBillRecordIsUnpaid(gomock.Any(), gomock.Any()).
			Return("", ErrSomeError)

		id, err := (*service).UpdateBillIsUnpaid(context.Background(), TEST_ID)

		assert.NotNil(t, err)
		assert.Equal(t, "", id)
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
	
		id, err := (*service).DeleteBillById(context.Background(), TEST_ID)

		assert.Equal(t, TEST_ID, id)
		assert.Nil(t, err)
	}

	{	// validation error	
		id, err := (*service).DeleteBillById(context.Background(), "")

		assert.NotNil(t, err)
		assert.Equal(t, "", id)
	}

	{	// error found
		mocks.DB.EXPECT().
			DeleteBillRecordById(gomock.Any(), gomock.Any()).
			Return("", ErrSomeError)

		id, err := (*service).DeleteBillById(context.Background(), TEST_ID)

		assert.NotNil(t, err)
		assert.Equal(t, "", id)
	}
}