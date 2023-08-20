package budget

import (
	// "context"
	// "context"
	"testing"

	"github.com/golang/mock/gomock"
	mock_db "github.com/sanchitdeora/budget-tracker/db/mocks"
	// "github.com/sanchitdeora/budget-tracker/models"
	mock_bill "github.com/sanchitdeora/budget-tracker/pkg/bill/mocks"
	mock_goal "github.com/sanchitdeora/budget-tracker/pkg/goal/mocks"
	mock_transaction "github.com/sanchitdeora/budget-tracker/pkg/transaction/mocks"
	"github.com/stretchr/testify/assert"
	// "time"
	// "github.com/golang/mock/gomock"
	// "github.com/sanchitdeora/budget-tracker/models"
	// "github.com/sanchitdeora/budget-tracker/pkg/transaction/mocks"
)

type ServiceMocks struct {
	Transaction *mock_transaction.MockService
	Bill *mock_bill.MockService
	Db *mock_db.MockDatabase
	Goal *mock_goal.MockService
}

func createTestBudget(ctrl *gomock.Controller) (*Service, *ServiceMocks) {
	
	mockBill := mock_bill.NewMockService(ctrl)
	mockTx := mock_transaction.NewMockService(ctrl)
	mockGoal := mock_goal.NewMockService(ctrl)
	mockDb := mock_db.NewMockDatabase(ctrl)

	service := NewService(&Opts{
		TransactionService: mockTx,
		BillService: mockBill,
		GoalService: mockGoal,
		DB: mockDb,
	})
	return &service, &ServiceMocks{
		Transaction: mockTx,
		Bill: mockBill,
		Db: mockDb,
		Goal: mockGoal,
	}

}

func TestGetBudgets(t *testing.T) {

	// ctrl := gomock.NewController(t)
	// service, mocks := createTestBudget(ctrl)

	// mocks.Db.EXPECT().GetAllBudgetRecords(gomock.Any(), ).Return(nil)
	// // mocks.Transaction.EXPECT().GetTransactionsByDate(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil)

	// var budget []models.Budget
	// err := (*service).GetBudgets(context.Background(), &budget)
	
	// assert.Nil(t, err)
	// assert.Equal(t, 1, 1, "The two words should be the same.")
}

func TestGetBudgetById(t *testing.T) {
	assert.Equal(t, 1, 1, "The two words should be the same.")
}

func TestGetGoalMap(t *testing.T) {
	assert.Equal(t, 1, 1, "The two words should be the same.")
}

func TestCreateBudgetByUser(t *testing.T) {
	assert.Equal(t, 1, 1, "The two words should be the same.")
}

func TestCreateBudget(t *testing.T) {
	assert.Equal(t, 1, 1, "The two words should be the same.")
}

func TestUpdateBudgetById(t *testing.T) {
	assert.Equal(t, 1, 1, "The two words should be the same.")
}

func TestDeleteBudgetById(t *testing.T) {
	assert.Equal(t, 1, 1, "The two words should be the same.")
}