package budget

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mock_db "github.com/sanchitdeora/budget-tracker/db/mocks"
	"github.com/sanchitdeora/budget-tracker/models"
	mock_bill "github.com/sanchitdeora/budget-tracker/pkg/bill/mocks"
	"github.com/sanchitdeora/budget-tracker/pkg/exceptions"
	mock_goal "github.com/sanchitdeora/budget-tracker/pkg/goal/mocks"
	mock_transaction "github.com/sanchitdeora/budget-tracker/pkg/transaction/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	TEST_ID 		 	  = "test-id"
	TEST_ID_2 		 	  = "test-id-2"
	TEST_NAME		 	  = "test-name"
	TEST_INCOME_MAP_ID_1  = "test-income-map-id-1"
	TEST_EXPENSE_MAP_ID_1 = "test-expense-map-id-1"
	TEST_GOAL_MAP_ID_1    = "test-goal-map-id-1"
	TEST_GOAL_MAP_ID_2    = "test-goal-map-id-2"
	TEST_AMOUNT_500 	  = float32(500)
	
	TEST_TRANSACTION_ID_DEBIT_1  = "test-transaction-id-debit-1"
	TEST_TRANSACTION_ID_DEBIT_2  = "test-transaction-id-debit-2"
	TEST_TRANSACTION_ID_CREDIT_1 = "test-transaction-id-credit-1"
	TEST_TRANSACTION_ID_CREDIT_2 = "test-transaction-id-credit-2"
	TEST_TRANSACTION_TYPE_CREDIT = true
	TEST_TRANSACTION_TYPE_DEBIT  = false
)

var (
	TEST_CREATION_TIME 	 = time.Unix(1600000000, 0)
	TEST_EXPIRATION_TIME = time.Unix(1700000000, 0)
	ErrSomeError 		 = errors.New("error found")
	
		TEST_BUDGET_HAPPY_PATH = models.Budget{
			BudgetId: TEST_ID,
			BudgetName: TEST_NAME,
			CreationTime: TEST_CREATION_TIME,
			ExpirationTime: TEST_EXPIRATION_TIME,
			IncomeMap: []models.BudgetInputMap{{
				Id: models.INCOME_CATEGORY,
				Name:TEST_NAME,
			}},
			ExpenseMap: []models.BudgetInputMap{
				{
					Id: models.BILLS_AND_UTILITIES_CATEGORY,
					Name:TEST_NAME,
				},
				{
					Id: models.UNCATEGORIZED_CATEGORY,
					Name:TEST_NAME,
				},
			},
			GoalMap: []models.BudgetInputMap{{
				Id: TEST_GOAL_MAP_ID_1,
				Name: TEST_NAME,
			}},
		}

	TEST_BUDGETS_HAPPY_PATH = []models.Budget{TEST_BUDGET_HAPPY_PATH}

	TEST_TRANSACTIONS_HAPPY_PATH = []models.Transaction{
		{
			TransactionId: TEST_TRANSACTION_ID_DEBIT_1,
			Category: models.BILLS_AND_UTILITIES_CATEGORY,
			Type: TEST_TRANSACTION_TYPE_DEBIT,
			Amount: TEST_AMOUNT_500,
		},
		{
			TransactionId: TEST_TRANSACTION_ID_DEBIT_2,
			Category: models.TAXES_CATEGORY,
			Type: TEST_TRANSACTION_TYPE_DEBIT,
			Amount: TEST_AMOUNT_500,
		},
		{
			TransactionId: TEST_TRANSACTION_ID_CREDIT_1,
			Category: models.INCOME_CATEGORY,
			Type: TEST_TRANSACTION_TYPE_CREDIT,
			Amount: TEST_AMOUNT_500,
		},
		{
			TransactionId: TEST_TRANSACTION_ID_CREDIT_2,
			Category: models.INVESTMENTS_CATEGORY,
			Type: TEST_TRANSACTION_TYPE_CREDIT,
			Amount: TEST_AMOUNT_500,
		},
	}
)

type ServiceMocks struct {
	Transaction *mock_transaction.MockService
	Bill *mock_bill.MockService
	DB *mock_db.MockDatabase
	Goal *mock_goal.MockService
}

func createTestBudget(ctrl *gomock.Controller) (Service, *ServiceMocks) {
	
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
	return service, &ServiceMocks{
		Transaction: mockTx,
		Bill: mockBill,
		DB: mockDb,
		Goal: mockGoal,
	}

}


// TestGetBudgets

func TestGetBudgets_HappyPath(t *testing.T) {

	ctrl := gomock.NewController(t)
	service, mocks := createTestBudget(ctrl)

	{	// happy path
		mocks.DB.EXPECT().
			GetAllBudgetRecords(gomock.Any()).
			Return(&TEST_BUDGETS_HAPPY_PATH, nil)

		mocks.Transaction.EXPECT().
			GetTransactionsByDate(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&TEST_TRANSACTIONS_HAPPY_PATH, nil)

		budgets, err := service.GetBudgets(context.Background())

		assert.NotNil(t, budgets)
		assert.Equal(t, models.INCOME_CATEGORY, (*budgets)[0].IncomeMap[0].Id)
		assert.Equal(t, TEST_AMOUNT_500, (*budgets)[0].IncomeMap[0].CurrentAmount)
		assert.Equal(t, models.UNCATEGORIZED_CATEGORY, (*budgets)[0].IncomeMap[1].Id)
		assert.Equal(t, TEST_AMOUNT_500, (*budgets)[0].IncomeMap[1].CurrentAmount)

		assert.Equal(t, models.BILLS_AND_UTILITIES_CATEGORY, (*budgets)[0].ExpenseMap[0].Id)
		assert.Equal(t, TEST_AMOUNT_500, (*budgets)[0].ExpenseMap[0].CurrentAmount)
		assert.Equal(t, models.UNCATEGORIZED_CATEGORY, (*budgets)[0].ExpenseMap[1].Id)
		assert.Equal(t, TEST_AMOUNT_500, (*budgets)[0].ExpenseMap[1].CurrentAmount)
		
		assert.Nil(t, err)
	}

	{	// happy path with no transactions and empty income map
		expBudgets := TEST_BUDGETS_HAPPY_PATH
		expBudgets[0].IncomeMap = make([]models.BudgetInputMap, 0)

		mocks.DB.EXPECT().
			GetAllBudgetRecords(gomock.Any()).
			Return(&expBudgets, nil)

		mocks.Transaction.EXPECT().
			GetTransactionsByDate(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&[]models.Transaction{}, nil)

		budgets, err := service.GetBudgets(context.Background())

		assert.NotNil(t, budgets)
		assert.Equal(t, 0, len((*budgets)[0].IncomeMap))

		assert.Equal(t, models.BILLS_AND_UTILITIES_CATEGORY, (*budgets)[0].ExpenseMap[0].Id)
		assert.Equal(t, float32(0), (*budgets)[0].ExpenseMap[0].CurrentAmount)
		assert.Equal(t, models.UNCATEGORIZED_CATEGORY, (*budgets)[0].ExpenseMap[1].Id)
		assert.Equal(t, float32(0), (*budgets)[0].ExpenseMap[1].CurrentAmount)
		
		assert.Nil(t, err)
	}
}

func TestGetBudgets_ReturnsError(t *testing.T) {

	ctrl := gomock.NewController(t)
	service, mocks := createTestBudget(ctrl)

	{	// error found while getting all budgets from db
		mocks.DB.EXPECT().
			GetAllBudgetRecords(gomock.Any()).
			Return(nil, ErrSomeError)

		budgets, err := service.GetBudgets(context.Background())

		assert.Nil(t, budgets)
		assert.Equal(t, ErrSomeError, err)
	}

	{	// error found while getting transactions by date
		budgets := &[]models.Budget{{
			BudgetId: TEST_ID,
			BudgetName: TEST_NAME,
			CreationTime: TEST_CREATION_TIME,
			ExpirationTime: TEST_EXPIRATION_TIME,
		}}
		mocks.DB.EXPECT().
			GetAllBudgetRecords(gomock.Any()).
			Return(budgets, nil)

		mocks.Transaction.EXPECT().
			GetTransactionsByDate(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil, ErrSomeError)

		budgets, err := service.GetBudgets(context.Background())

		assert.Nil(t, budgets)
		assert.Equal(t, ErrSomeError, err)
	}

	{	// no budgets found
		budgets := &[]models.Budget{}
		mocks.DB.EXPECT().
			GetAllBudgetRecords(gomock.Any()).
			Return(budgets, nil)

		budgets, err := service.GetBudgets(context.Background())

		assert.NotNil(t, budgets)
		assert.Nil(t, err)
	}

}

func TestGetBudgetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBudget(ctrl)
	{ 	// happy path
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), gomock.Any()).
			Return(&TEST_BUDGET_HAPPY_PATH, nil)
		mocks.Transaction.EXPECT().
			GetTransactionsByDate(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&TEST_TRANSACTIONS_HAPPY_PATH, nil)

		budget, err := service.GetBudgetById(context.Background(), TEST_ID)

		assert.Nil(t, err)
		assert.NotNil(t, budget)
		assert.Equal(t, models.INCOME_CATEGORY, budget.IncomeMap[0].Id)
		assert.Equal(t, TEST_AMOUNT_500, budget.IncomeMap[0].CurrentAmount)
		assert.Equal(t, models.UNCATEGORIZED_CATEGORY, budget.IncomeMap[1].Id)
		assert.Equal(t, TEST_AMOUNT_500, budget.IncomeMap[1].CurrentAmount)

		assert.Equal(t, models.BILLS_AND_UTILITIES_CATEGORY, budget.ExpenseMap[0].Id)
		assert.Equal(t, TEST_AMOUNT_500, budget.ExpenseMap[0].CurrentAmount)
		assert.Equal(t, models.UNCATEGORIZED_CATEGORY, budget.ExpenseMap[1].Id)
		assert.Equal(t, TEST_AMOUNT_500, budget.ExpenseMap[1].CurrentAmount)
		
	}

	{ 	// validation error empty budget id
		budget, err := service.GetBudgetById(context.Background(), "")

		assert.Nil(t, budget)
		assert.NotNil(t, err)
		assert.Equal(t, exceptions.ErrValidationError, err)
	}

	{ 	// not found error no budget
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), gomock.Any()).
			Return(nil, exceptions.ErrBudgetNotFound)

		budget, err := service.GetBudgetById(context.Background(), TEST_ID)

		assert.Nil(t, budget)
		assert.NotNil(t, err)
		assert.Equal(t, exceptions.ErrBudgetNotFound, err)
	}

	{	// error while updating current amounts
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), gomock.Any()).
			Return(&TEST_BUDGET_HAPPY_PATH, nil)
		mocks.Transaction.EXPECT().
			GetTransactionsByDate(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil, ErrSomeError)

		budget, err := service.GetBudgetById(context.Background(), TEST_ID)

		assert.Nil(t, budget)
		assert.Equal(t, ErrSomeError, err)

	}
}


// TestCreateBudgetByUser
func TestCreateBudgetByUser_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBudget(ctrl)

	budget := TEST_BUDGET_HAPPY_PATH
	budget.BudgetId = ""

	mocks.DB.EXPECT().
		InsertBudgetRecord(gomock.Any(), &budget).
		Return(TEST_ID, nil)
	mocks.Goal.EXPECT().
		UpdateBudgetIdsList(gomock.Any(), TEST_GOAL_MAP_ID_1, TEST_ID).
		Return("", nil)

	id, err := service.CreateBudgetByUser(context.Background(), &budget)

	assert.Equal(t, TEST_ID, id)
	assert.Nil(t, err)
}

func TestCreateBudgetByUser_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBudget(ctrl)
	{ 	// validation error
		budget := &models.Budget{IsClosed: true}

		id, err := service.CreateBudgetByUser(context.Background(), budget)

		assert.Equal(t, exceptions.ErrValidationError, err)
		assert.Equal(t, "", id)
	}

	{ 	// error while inserting budget
		budget := &models.Budget{BudgetName: TEST_NAME}
		mocks.DB.EXPECT().
			InsertBudgetRecord(gomock.Any(), budget).
			Return("", ErrSomeError)

		id, err := service.CreateBudgetByUser(context.Background(), budget)

		assert.Equal(t, ErrSomeError, err)
		assert.Equal(t, "", id)
	}

	{ 	// error while updating budget ids list in goal
		budget := TEST_BUDGET_HAPPY_PATH
		budget.BudgetId = ""
		mocks.DB.EXPECT().
			InsertBudgetRecord(gomock.Any(), gomock.Any()).
			Return(TEST_ID, nil)
		mocks.Goal.EXPECT().
			UpdateBudgetIdsList(gomock.Any(), TEST_GOAL_MAP_ID_1, TEST_ID).
			Return("", ErrSomeError)

		id, err := service.CreateBudgetByUser(context.Background(), &budget)

		assert.Equal(t, ErrSomeError, err)
		assert.Equal(t, "", id)
	}

}


// TestCreateRecurringBudget

func TestCreateRecurringBudget_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBudget(ctrl)
	
	timeNow := time.Now()
	prevBudget := &models.Budget{
		BudgetId: TEST_ID, 
		BudgetName: TEST_NAME,
		ExpenseMap: []models.BudgetInputMap{{
			Id: models.BILLS_AND_UTILITIES_CATEGORY,
			Amount: 100,
			Name: TEST_NAME,
		}},
		ExpirationTime: timeNow,
		Frequency: models.MONTHLY_FREQUENCY,
	}

	expBudget := &models.Budget{
		BudgetName: TEST_NAME,
		ExpenseMap: []models.BudgetInputMap{{
			Id: models.BILLS_AND_UTILITIES_CATEGORY,
			Amount: 100,
			Name: TEST_NAME,
		}},
		Savings: -100,
		Frequency: models.MONTHLY_FREQUENCY,
		CreationTime: timeNow,
		ExpirationTime: timeNow.AddDate(0, 1, 0),
		SequenceNumber: 1,
		SequenceStartId: TEST_ID,
	}

	mocks.DB.EXPECT().
		InsertBudgetRecord(gomock.Any(), expBudget).
		Return(TEST_ID_2, nil)

	id, err := service.CreateRecurringBudget(context.Background(), prevBudget)

	assert.Nil(t, err)
	assert.Equal(t, TEST_ID_2, id)
}

func TestCreateRecurringBudget_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBudget(ctrl)
	{ 	// validation error for current budget
		id, err := service.CreateRecurringBudget(context.Background(), nil)

		assert.Equal(t, exceptions.ErrValidationError, err)
		assert.Equal(t, "", id)
	}

	{ 	// validation error for prev budget when no budget id
		id, err := service.CreateRecurringBudget(context.Background(), &models.Budget{})

		assert.Equal(t, exceptions.ErrValidationError, err)
		assert.Equal(t, "", id)
	}

	{	// validation error for prev budget when invalid transaction provided
		prevBudget := &models.Budget{BudgetId: "id", BudgetName: TEST_NAME, IsClosed: true}
		id, err := service.CreateRecurringBudget(context.Background(), prevBudget)

		assert.Equal(t, exceptions.ErrValidationError, err)
		assert.Equal(t, "", id)
	}

	{ 	// error while inserting budget
		prevBudget := &models.Budget{BudgetId: "id", BudgetName: TEST_NAME}
		mocks.DB.EXPECT().
			InsertBudgetRecord(gomock.Any(), gomock.Any()).
			Return("", ErrSomeError)

		id, err := service.CreateRecurringBudget(context.Background(), prevBudget)

		assert.Equal(t, ErrSomeError, err)
		assert.Equal(t, "", id)
	}
}

func TestUpdateBudgetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBudget(ctrl)
	{	// validation error	when no budget id	
		id, err := service.UpdateBudgetById(context.Background(), "", &models.Budget{})

		assert.Equal(t, exceptions.ErrValidationError, err)
		assert.Equal(t, "", id)
	}

	{	// validation error when invalid budget provided
		budget := &models.Budget{IsClosed: true}
		id, err := service.UpdateBudgetById(context.Background(), TEST_ID, budget)

		assert.Equal(t, exceptions.ErrValidationError, err)
		assert.Equal(t, "", id)
	}

	{	// error found while getting budget id
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_ID).
			Return(nil, ErrSomeError)

		id, err := service.UpdateBudgetById(context.Background(), TEST_ID, &models.Budget{BudgetName: TEST_NAME})

		assert.Equal(t, ErrSomeError, err)
		assert.Equal(t, "", id)
	}

	
	{	// error found while removing budget id from goal
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_ID).
			Return(&TEST_BUDGET_HAPPY_PATH, nil)
		mocks.Transaction.EXPECT().
			GetTransactionsByDate(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&TEST_TRANSACTIONS_HAPPY_PATH, nil)
		mocks.Goal.EXPECT().
			RemoveBudgetIdFromGoal(gomock.Any(), TEST_GOAL_MAP_ID_1, TEST_ID).
			Return("", ErrSomeError)

		id, err := service.UpdateBudgetById(context.Background(), TEST_ID, &models.Budget{BudgetName: TEST_NAME})

		assert.Equal(t, ErrSomeError, err)
		assert.Equal(t, "", id)
	}

	{	// error found while updating budget ids list from goal
		budget := &models.Budget{
			BudgetName: TEST_NAME,
			GoalMap: []models.BudgetInputMap{{
				Id: TEST_GOAL_MAP_ID_2,
			}},
		}
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_ID).
			Return(&TEST_BUDGET_HAPPY_PATH, nil)
		mocks.Transaction.EXPECT().
			GetTransactionsByDate(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&TEST_TRANSACTIONS_HAPPY_PATH, nil)
		mocks.Goal.EXPECT().
			RemoveBudgetIdFromGoal(gomock.Any(), TEST_GOAL_MAP_ID_1, TEST_ID).
			Return("", nil)
		mocks.Goal.EXPECT().
			UpdateBudgetIdsList(gomock.Any(), TEST_GOAL_MAP_ID_2, TEST_ID).
			Return("", ErrSomeError)

		id, err := service.UpdateBudgetById(context.Background(), TEST_ID, budget)

		assert.Equal(t, ErrSomeError, err)
		assert.Equal(t, "", id)
	}

	{	// error found while updating budget
		budget := &models.Budget{
			BudgetName: TEST_NAME,
			GoalMap: []models.BudgetInputMap{{
				Id: TEST_GOAL_MAP_ID_2,
			}},
		}
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_ID).
			Return(&TEST_BUDGET_HAPPY_PATH, nil)
		mocks.Transaction.EXPECT().
			GetTransactionsByDate(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&TEST_TRANSACTIONS_HAPPY_PATH, nil)
		mocks.Goal.EXPECT().
			RemoveBudgetIdFromGoal(gomock.Any(), TEST_GOAL_MAP_ID_1, TEST_ID).
			Return("", nil)
		mocks.Goal.EXPECT().
			UpdateBudgetIdsList(gomock.Any(), TEST_GOAL_MAP_ID_2, TEST_ID).
			Return("", nil)
		mocks.DB.EXPECT().
			UpdateBudgetRecordById(gomock.Any(), TEST_ID, gomock.Any()).
			Return("", ErrSomeError)

		id, err := service.UpdateBudgetById(context.Background(), TEST_ID, budget)

		assert.Equal(t, ErrSomeError, err)
		assert.Equal(t, "", id)
	}

	{	// happy path
		budget := &models.Budget{
			BudgetName: TEST_NAME,
			GoalMap: []models.BudgetInputMap{{
				Id: TEST_GOAL_MAP_ID_2,
			}},
		}
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_ID).
			Return(&TEST_BUDGET_HAPPY_PATH, nil)
		mocks.Transaction.EXPECT().
			GetTransactionsByDate(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&TEST_TRANSACTIONS_HAPPY_PATH, nil)
		mocks.Goal.EXPECT().
			RemoveBudgetIdFromGoal(gomock.Any(), TEST_GOAL_MAP_ID_1, TEST_ID).
			Return("", nil)
		mocks.Goal.EXPECT().
			UpdateBudgetIdsList(gomock.Any(), TEST_GOAL_MAP_ID_2, TEST_ID).
			Return("", nil)
		mocks.DB.EXPECT().
			UpdateBudgetRecordById(gomock.Any(), TEST_ID, gomock.Any()).
			Return(TEST_ID, nil)

		id, err := service.UpdateBudgetById(context.Background(), TEST_ID, budget)

		assert.Equal(t, TEST_ID, id)
		assert.Nil(t, err)
	}
}


// TestUpdateBudgetIsClosed

func TestUpdateBudgetIsClosed_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBudget(ctrl)

	{	// happy path
		budget := TEST_BUDGET_HAPPY_PATH

		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_ID).
			Return(&budget, nil)
		mocks.Transaction.EXPECT().
			GetTransactionsByDate(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&TEST_TRANSACTIONS_HAPPY_PATH, nil)

		budget.IsClosed = true

		mocks.DB.EXPECT().
			UpdateBudgetRecordById(gomock.Any(), TEST_ID, gomock.Any()).
			Return(TEST_ID, nil)

		id, err := service.UpdateBudgetIsClosed(context.Background(), TEST_ID, true)
		
		assert.Equal(t, TEST_ID, id)
		assert.Nil(t, err)
	}
}

func TestUpdateBudgetIsClosed_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBudget(ctrl)

	{	// validation error	when no budget id	
		id, err := service.UpdateBudgetIsClosed(context.Background(), "", true)
		
		assert.Equal(t, "", id)
		assert.Equal(t, exceptions.ErrValidationError, err)
	}

	{	// error found while fetching budget by id
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_ID).
			Return(nil, ErrSomeError)

		id, err := service.UpdateBudgetIsClosed(context.Background(), TEST_ID, true)
		
		assert.Equal(t, "", id)
		assert.Equal(t, ErrSomeError, err)
	}

	{	// error found while updating budget by id
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_ID).
			Return(&TEST_BUDGET_HAPPY_PATH, nil)
		mocks.Transaction.EXPECT().
			GetTransactionsByDate(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&TEST_TRANSACTIONS_HAPPY_PATH, nil)

		mocks.DB.EXPECT().
			UpdateBudgetRecordById(gomock.Any(), TEST_ID, gomock.Any()).
			Return("", ErrSomeError)

		id, err := service.UpdateBudgetIsClosed(context.Background(), TEST_ID, true)

		assert.Equal(t, "", id)
		assert.Equal(t, ErrSomeError, err)
	}
}

func TestDeleteBudgetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBudget(ctrl)
	{	// happy path
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_ID).
			Return(&TEST_BUDGET_HAPPY_PATH, nil)
		mocks.Transaction.EXPECT().
			GetTransactionsByDate(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&TEST_TRANSACTIONS_HAPPY_PATH, nil)
		mocks.Goal.EXPECT().
			RemoveBudgetIdFromGoal(gomock.Any(), TEST_GOAL_MAP_ID_1, TEST_ID).
			Return("", nil)
		mocks.DB.EXPECT().
			DeleteBudgetRecordById(gomock.Any(),TEST_ID).
			Return(TEST_ID, nil)

		id, err := service.DeleteBudgetById(context.Background(), TEST_ID)

		assert.Equal(t, TEST_ID, id)
		assert.Nil(t, err)
	}

	{	// validation error	
		id, err := service.DeleteBudgetById(context.Background(), "")

		assert.NotNil(t, err)
		assert.Equal(t, "", id)
	}

	{	// error found while getting budget by id
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_ID).
			Return(nil, exceptions.ErrBudgetNotFound)

		id, err := service.DeleteBudgetById(context.Background(), TEST_ID)

		assert.Equal(t, exceptions.ErrBudgetNotFound, err)
		assert.Equal(t, "", id)
	}

	{	// error found while calling removing budget ids from goal
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_ID).
			Return(&TEST_BUDGET_HAPPY_PATH, nil)
		mocks.Transaction.EXPECT().
			GetTransactionsByDate(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&TEST_TRANSACTIONS_HAPPY_PATH, nil)
		mocks.Goal.EXPECT().
			RemoveBudgetIdFromGoal(gomock.Any(), TEST_GOAL_MAP_ID_1, TEST_ID).
			Return("", ErrSomeError)

		id, err := service.DeleteBudgetById(context.Background(), TEST_ID)

		assert.Equal(t, ErrSomeError, err)
		assert.Equal(t, "", id)
	}

	{	// error found while deleting budget by id
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_ID).
			Return(&TEST_BUDGET_HAPPY_PATH, nil)
		mocks.Transaction.EXPECT().
			GetTransactionsByDate(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(&TEST_TRANSACTIONS_HAPPY_PATH, nil)
		mocks.Goal.EXPECT().
			RemoveBudgetIdFromGoal(gomock.Any(), TEST_GOAL_MAP_ID_1, TEST_ID).
			Return("", nil)
		mocks.DB.EXPECT().
			DeleteBudgetRecordById(gomock.Any(), TEST_ID).
			Return("", ErrSomeError)

		id, err := service.DeleteBudgetById(context.Background(), TEST_ID)

		assert.Equal(t, ErrSomeError, err)
		assert.Equal(t, "", id)
	}
}


// TestBudgetMaintainer

func TestBudgetMaintainer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestBudget(ctrl)

	timeNow := time.Now()
	{	// maintainer lifecycle
		budgets := &[]models.Budget{
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
				ExpirationTime: timeNow.AddDate(0, 0, 1),
			},
			{
				BudgetId: TEST_ID+"-fail",
				BudgetName: TEST_NAME,
				Frequency: models.MONTHLY_FREQUENCY,
				ExpirationTime: timeNow.AddDate(0, 0, -1),
			},
			{
				BudgetId: TEST_ID,
				BudgetName: TEST_NAME,
				Frequency: models.MONTHLY_FREQUENCY,
				ExpirationTime: timeNow.AddDate(0, 0, -1),
			},
			{
				BudgetId: TEST_ID,
				BudgetName: TEST_NAME,
				Frequency: models.MONTHLY_FREQUENCY,
				ExpirationTime: timeNow.AddDate(0, 0, -1),
			},
		}

		expBudget := &models.Budget{
			BudgetName: TEST_NAME,
			Frequency: models.MONTHLY_FREQUENCY,
			CreationTime: timeNow.AddDate(0, 0, -1),
			ExpirationTime: timeNow.AddDate(0, 1, -1),
			SequenceStartId: TEST_ID,
			SequenceNumber: 1,
		}

		mocks.DB.EXPECT().
			GetAllBudgetRecords(gomock.Any()).
			Return(budgets, nil)
		mocks.Transaction.EXPECT().
			GetTransactionsByDate(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil, nil).
			AnyTimes()
		mocks.DB.EXPECT().
			InsertBudgetRecord(gomock.Any(), gomock.Any()).
			Return("", ErrSomeError)
		mocks.DB.EXPECT().
			InsertBudgetRecord(gomock.Any(), expBudget).
			Return(TEST_ID_2, nil).
			Times(2)
		mocks.DB.EXPECT().
			UpdateBudgetRecordById(gomock.Any(), gomock.Any(), gomock.Any()).
			Return("", ErrSomeError)
		mocks.DB.EXPECT().
			UpdateBudgetRecordById(gomock.Any(), TEST_ID, gomock.Any()).
			Return(TEST_ID, nil)

		service.BudgetMaintainer(context.Background())
	}

	{	// error found while fetching budget, close maintainer
		mocks.DB.EXPECT().
			GetAllBudgetRecords(gomock.Any()).
			Return(nil, ErrSomeError)

		service.BudgetMaintainer(context.Background())
	}
}
