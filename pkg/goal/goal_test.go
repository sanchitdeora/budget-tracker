package goal

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mock_db "github.com/sanchitdeora/budget-tracker/db/mocks"
	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/pkg/exceptions"
	"github.com/stretchr/testify/assert"
)

const (
	TEST_ID 		 = "test-id"
	TEST_BUDGET_ID 	 = "test-budget-id"
	TEST_BUDGET_ID_2 = "test-budget-id-2"
	TEST_TITLE		 = "test-title"
)

var (
	ErrSomeError = errors.New("error found")
)

type ServiceMocks struct {
	DB *mock_db.MockDatabase
}

func createTestGoal(ctrl *gomock.Controller) (Service, *ServiceMocks) {
	db := mock_db.NewMockDatabase(ctrl)
	service := NewService(&Opts{DB: db})
	return service, &ServiceMocks{DB: db}
}

func TestGetGoals(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestGoal(ctrl)

	{ 	// happy path
		expGoals := &[]models.Goal{{
			GoalId: TEST_ID,
			CurrentAmount: 10,
			TargetAmount: 50,
			BudgetIdList: []string{TEST_BUDGET_ID},
		}}
		expBudget := &models.Budget{
			BudgetId: TEST_BUDGET_ID,
			GoalMap: []models.BudgetInputMap{{
				Id:   TEST_ID,
				Name: TEST_TITLE,
				CurrentAmount: 10,
				Amount: 20,
			}},
		}
		mocks.DB.EXPECT().
			GetAllGoalRecords(gomock.Any()).
			Return(expGoals, nil)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID).
			Return(expBudget, nil)

		goals, err := service.GetGoals(context.Background())

		assert.Nil(t, err)
		assert.NotNil(t, goals)
		assert.Equal(t, 1, len(*goals))
		assert.Equal(t, float32(20), (*goals)[0].CurrentAmount)
	}

	{ 	// no goals found
		mocks.DB.EXPECT().
			GetAllGoalRecords(gomock.Any()).
			Return(&[]models.Goal{}, nil)

		goals, err := service.GetGoals(context.Background())

		assert.Nil(t, err)
		assert.NotNil(t, goals)
		assert.Equal(t, 0, len(*goals))
	}
	
	{ 	// error found while getting all goals
		mocks.DB.EXPECT().
			GetAllGoalRecords(gomock.Any()).
			Return(nil, ErrSomeError)

		goals, err := service.GetGoals(context.Background())

		assert.NotNil(t, err)
		assert.Equal(t, err, ErrSomeError)
		assert.Nil(t, goals)
	}

	{ 	// error found while getting all budgets from budgetIdList
		expGoals := &[]models.Goal{{
			GoalId: TEST_ID,
			CurrentAmount: 10,
			TargetAmount: 50,
			BudgetIdList: []string{TEST_BUDGET_ID},
		}}

		mocks.DB.EXPECT().
			GetAllGoalRecords(gomock.Any()).
			Return(expGoals, nil)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID).
			Return(nil, ErrSomeError)

		goals, err := service.GetGoals(context.Background())

		assert.NotNil(t, err)
		assert.Equal(t, err, ErrSomeError)
		assert.Nil(t, goals)
	}

	{ 	// no budget found
		expGoals := &[]models.Goal{{
			GoalId: TEST_ID,
			CurrentAmount: 10,
			TargetAmount: 50,
			BudgetIdList: []string{TEST_BUDGET_ID},
		}}

		mocks.DB.EXPECT().
			GetAllGoalRecords(gomock.Any()).
			Return(expGoals, nil)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID).
			Return(nil, nil)

		goals, err := service.GetGoals(context.Background())

		assert.NotNil(t, err)
		assert.Equal(t, exceptions.ErrNoBudgetsFound, err)
		assert.Nil(t, goals)
	}
}

func TestGetGoalById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestGoal(ctrl)

	{ 	// happy path
		expGoal := &models.Goal{
			GoalId: TEST_ID,
			CurrentAmount: 10,
			TargetAmount: 50,
			BudgetIdList: []string{TEST_BUDGET_ID},
		}
		expBudget := &models.Budget{
			BudgetId: TEST_BUDGET_ID,
			GoalMap: []models.BudgetInputMap{{
				Id:   TEST_ID,
				Name: TEST_TITLE,
				CurrentAmount: 10,
				Amount: 20,
			}},
		}
		mocks.DB.EXPECT().
			GetGoalRecordById(gomock.Any(), TEST_ID).
			Return(expGoal, nil)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID).
			Return(expBudget, nil)

		goal, err := service.GetGoalById(context.Background(), TEST_ID)

		assert.Nil(t, err)
		assert.NotNil(t, goal)
		assert.Equal(t, float32(20), goal.CurrentAmount)
	}

	{ 	// validation error empty goal_id
		goal, err := service.GetGoalById(context.Background(), "")

		assert.NotNil(t, err)
		assert.Equal(t, exceptions.ErrValidationError, err)
		assert.Nil(t, goal)
	}

	{ 	// no goals found
		mocks.DB.EXPECT().
			GetGoalRecordById(gomock.Any(), gomock.Any()).
			Return(nil, exceptions.ErrGoalNotFound)

		goal, err := service.GetGoalById(context.Background(), TEST_ID)

		assert.NotNil(t, err)
		assert.Equal(t, exceptions.ErrGoalNotFound, err)
		assert.Nil(t, goal)
	}

	{ 	// error found while getting budget
		expGoal := &models.Goal{
			GoalId: TEST_ID,
			CurrentAmount: 10,
			TargetAmount: 50,
			BudgetIdList: []string{TEST_BUDGET_ID},
		}

		mocks.DB.EXPECT().
			GetGoalRecordById(gomock.Any(), TEST_ID).
			Return(expGoal, nil)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID).
			Return(nil, ErrSomeError)

		goal, err := service.GetGoalById(context.Background(), TEST_ID)

		assert.NotNil(t, err)
		assert.Equal(t, err, ErrSomeError)
		assert.Nil(t, goal)
	}
}

func TestCreateGoal(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestGoal(ctrl)
	{	// happy path
		expBudget := &models.Budget{
			BudgetId: TEST_BUDGET_ID,
		}
		mocks.DB.EXPECT().
			InsertGoalRecord(gomock.Any(), gomock.Any()).
			Return(TEST_ID, nil)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID).
			Return(expBudget, nil)
		mocks.DB.EXPECT().
			UpdateBudgetRecordById(gomock.Any(), TEST_BUDGET_ID, gomock.Any()).
			Return(TEST_BUDGET_ID, nil)

		goal := &models.Goal{
			GoalId: TEST_ID,
			GoalName: TEST_TITLE,
			TargetAmount: 100,
			BudgetIdList: []string{TEST_BUDGET_ID},
		}
		id, err := service.CreateGoal(context.Background(), goal)

		assert.Equal(t, TEST_ID, id)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(expBudget.GoalMap))
	}

	{	// validation error

		goal := &models.Goal{
			GoalName: TEST_TITLE,
			CurrentAmount: -50,
			TargetAmount: -100,
		}
		id, err := service.CreateGoal(context.Background(), goal)

		assert.NotNil(t, err)
		assert.Equal(t, err, exceptions.ErrValidationError)
		assert.Equal(t, "", id)
	}

	{	// error found
		mocks.DB.EXPECT().
			InsertGoalRecord(gomock.Any(), gomock.Any()).
			Return("", ErrSomeError)

		goal := &models.Goal{
			GoalName: TEST_TITLE,
			TargetAmount: 100,
		}
		id, err := service.CreateGoal(context.Background(), goal)

		assert.NotNil(t, err)
		assert.Equal(t, err, ErrSomeError)
		assert.Equal(t, "", id)
	}

	{	// error found while getting budget by id
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID).
			Return(nil, ErrSomeError)

		goal := &models.Goal{
			GoalName: TEST_TITLE,
			TargetAmount: 100,
			BudgetIdList: []string{TEST_BUDGET_ID},
		}
		id, err := service.CreateGoal(context.Background(), goal)

		assert.NotNil(t, err)
		assert.Equal(t, err, ErrSomeError)
		assert.Equal(t, "", id)
	}

	{	// error found while updating goal map in budget
		expBudget := &models.Budget{
			BudgetId: TEST_BUDGET_ID,
		}
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID).
			Return(expBudget, nil)
		mocks.DB.EXPECT().
			UpdateBudgetRecordById(gomock.Any(), TEST_BUDGET_ID, gomock.Any()).
			Return("", ErrSomeError)

		goal := &models.Goal{
			GoalName: TEST_TITLE,
			TargetAmount: 100,
			BudgetIdList: []string{TEST_BUDGET_ID},
		}
		id, err := service.CreateGoal(context.Background(), goal)

		assert.NotNil(t, err)
		assert.Equal(t, err, ErrSomeError)
		assert.Equal(t, "", id)
	}
}

func TestUpdateGoalById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestGoal(ctrl)
	{	// happy path
		expBudget1 := &models.Budget{
			BudgetId: TEST_BUDGET_ID,
			GoalMap: []models.BudgetInputMap{{
				Id:   TEST_ID,
				Name: TEST_TITLE,
				CurrentAmount: 10,
				Amount: 20,
			}},
		}
		expBudget2 := &models.Budget{
			BudgetId: TEST_BUDGET_ID_2,
			GoalMap: []models.BudgetInputMap{},
		}
		goal := &models.Goal{
			GoalId: TEST_ID,
			GoalName: TEST_TITLE,
			CurrentAmount: 50,
			TargetAmount: 100,
			BudgetIdList: []string{TEST_BUDGET_ID_2},
			
		}
		expGoal := &models.Goal{
			GoalId: TEST_ID,
			CurrentAmount: 10,
			TargetAmount: 50,
			BudgetIdList: []string{TEST_BUDGET_ID},
		}
		mocks.DB.EXPECT().
			GetGoalRecordById(gomock.Any(), TEST_ID).
			Return(expGoal, nil)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID).
			Return(expBudget1, nil).
			Times(2)
		mocks.DB.EXPECT().
			UpdateBudgetRecordById(gomock.Any(), TEST_BUDGET_ID, gomock.Any()).
			Return(TEST_BUDGET_ID, nil)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID_2).
			Return(expBudget2, nil)
		mocks.DB.EXPECT().
			UpdateBudgetRecordById(gomock.Any(), TEST_BUDGET_ID_2, gomock.Any()).
			Return(TEST_BUDGET_ID_2, nil)
		mocks.DB.EXPECT().
			UpdateGoalRecordById(gomock.Any(), TEST_ID, gomock.Any()).
			Return(TEST_ID, nil)

		id, err := service.UpdateGoalById(context.Background(), TEST_ID, goal)

		assert.Nil(t, err)
		assert.Equal(t, TEST_ID, id)
	}

	{	// validation error when goal Id is empty

		goal := &models.Goal{
			GoalName: TEST_TITLE,
			CurrentAmount: -50,
			TargetAmount: -100,
		}
		id, err := service.UpdateGoalById(context.Background(), "", goal)

		assert.NotNil(t, err)
		assert.Equal(t, exceptions.ErrValidationError, err)
		assert.Equal(t, "", id)
	}

	{	// validation error with body

		goal := &models.Goal{
			GoalId: TEST_ID,
			GoalName: TEST_TITLE,
			CurrentAmount: -50,
			TargetAmount: -100,
		}
		id, err := service.UpdateGoalById(context.Background(), TEST_ID, goal)

		assert.NotNil(t, err)
		assert.Equal(t, exceptions.ErrValidationError, err)
		assert.Equal(t, "", id)
	}

	{	// error found while getting goal
		goal := &models.Goal{
			GoalId: TEST_ID,
			GoalName: TEST_TITLE,
			CurrentAmount: 50,
			TargetAmount: 100,
			BudgetIdList: []string{TEST_BUDGET_ID},
		}
		mocks.DB.EXPECT().
			GetGoalRecordById(gomock.Any(), TEST_ID).
			Return(nil, ErrSomeError)

		id, err := service.UpdateGoalById(context.Background(), TEST_ID, goal)

		assert.NotNil(t, err)
		assert.Equal(t, err, ErrSomeError)
		assert.Equal(t, "", id)
	}

	{	// error found while removing goal map from budget
		expBudget := &models.Budget{
			BudgetId: TEST_BUDGET_ID,
			GoalMap: []models.BudgetInputMap{{
				Id:   TEST_ID,
				Name: TEST_TITLE,
				CurrentAmount: 10,
				Amount: 20,
			}},
		}
		goal := &models.Goal{
			GoalId: TEST_ID,
			GoalName: TEST_TITLE,
			CurrentAmount: 50,
			TargetAmount: 100,
			BudgetIdList: []string{TEST_BUDGET_ID_2},
			
		}
		expGoal := &models.Goal{
			GoalId: TEST_ID,
			CurrentAmount: 10,
			TargetAmount: 50,
			BudgetIdList: []string{TEST_BUDGET_ID},
		}
		mocks.DB.EXPECT().
			GetGoalRecordById(gomock.Any(), TEST_ID).
			Return(expGoal, nil)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID).
			Return(expBudget, nil).
			Times(1)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID).
			Return(nil, ErrSomeError).
			Times(1)

		id, err := service.UpdateGoalById(context.Background(), TEST_ID, goal)

		assert.NotNil(t, err)
		assert.Equal(t, err, ErrSomeError)
		assert.Equal(t, "", id)
	}

	{	// error found while updating goal map in budget
		expBudget := &models.Budget{
			BudgetId: TEST_BUDGET_ID,
			GoalMap: []models.BudgetInputMap{{
				Id:   TEST_ID,
				Name: TEST_TITLE,
				CurrentAmount: 10,
				Amount: 20,
			}},
		}
		goal := &models.Goal{
			GoalId: TEST_ID,
			GoalName: TEST_TITLE,
			CurrentAmount: 50,
			TargetAmount: 100,
			BudgetIdList: []string{TEST_BUDGET_ID_2},
			
		}
		expGoal := &models.Goal{
			GoalId: TEST_ID,
			CurrentAmount: 10,
			TargetAmount: 50,
			BudgetIdList: []string{TEST_BUDGET_ID},
		}
		mocks.DB.EXPECT().
			GetGoalRecordById(gomock.Any(), TEST_ID).
			Return(expGoal, nil)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID).
			Return(expBudget, nil).
			Times(2)
		mocks.DB.EXPECT().
			UpdateBudgetRecordById(gomock.Any(), TEST_BUDGET_ID, gomock.Any()).
			Return(TEST_BUDGET_ID, nil)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID_2).
			Return(nil, ErrSomeError)

		id, err := service.UpdateGoalById(context.Background(), TEST_ID, goal)

		assert.NotNil(t, err)
		assert.Equal(t, err, ErrSomeError)
		assert.Equal(t, "", id)
	}

	{	// error found while updating goal
		expBudget1 := &models.Budget{
			BudgetId: TEST_BUDGET_ID,
			GoalMap: []models.BudgetInputMap{{
				Id:   TEST_ID,
				Name: TEST_TITLE,
				CurrentAmount: 10,
				Amount: 20,
			}},
		}
		expBudget2 := &models.Budget{
			BudgetId: TEST_BUDGET_ID_2,
			GoalMap: []models.BudgetInputMap{},
		}
		goal := &models.Goal{
			GoalId: TEST_ID,
			GoalName: TEST_TITLE,
			CurrentAmount: 50,
			TargetAmount: 100,
			BudgetIdList: []string{TEST_BUDGET_ID_2},
			
		}
		expGoal := &models.Goal{
			GoalId: TEST_ID,
			CurrentAmount: 10,
			TargetAmount: 50,
			BudgetIdList: []string{TEST_BUDGET_ID},
		}
		mocks.DB.EXPECT().
			GetGoalRecordById(gomock.Any(), TEST_ID).
			Return(expGoal, nil)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID).
			Return(expBudget1, nil).
			Times(2)
		mocks.DB.EXPECT().
			UpdateBudgetRecordById(gomock.Any(), TEST_BUDGET_ID, gomock.Any()).
			Return(TEST_BUDGET_ID, nil)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID_2).
			Return(expBudget2, nil)
		mocks.DB.EXPECT().
			UpdateBudgetRecordById(gomock.Any(), TEST_BUDGET_ID_2, gomock.Any()).
			Return(TEST_BUDGET_ID_2, nil)
		mocks.DB.EXPECT().
			UpdateGoalRecordById(gomock.Any(), TEST_ID, gomock.Any()).
			Return("", ErrSomeError)

		id, err := service.UpdateGoalById(context.Background(), TEST_ID, goal)

		assert.NotNil(t, err)
		assert.Equal(t, err, ErrSomeError)
		assert.Equal(t, "", id)
	}
}

func TestUpdateBudgetIdsList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestGoal(ctrl)
	
	{	// happy path
		goal := &models.Goal{
			GoalId: TEST_ID,
		}

		mocks.DB.EXPECT().
			GetGoalRecordById(gomock.Any(), TEST_ID).
			Return(goal, nil)
		mocks.DB.EXPECT().
			UpdateGoalRecordById(gomock.Any(), TEST_ID, gomock.Any()).
			Return(TEST_ID, nil)

		goalId, err := service.UpdateBudgetIdsList(context.Background(), TEST_ID, TEST_BUDGET_ID)

		assert.Equal(t, TEST_ID, goalId)
		assert.Nil(t, err)
	}

	{	// validation error with goal id empty
		goalId, err := service.UpdateBudgetIdsList(context.Background(), "", "")

		assert.NotNil(t, err)
		assert.Equal(t, exceptions.ErrValidationError, err)
		assert.Equal(t, "", goalId)
	}

	{	// validation error with budget id empty
		goalId, err := service.UpdateBudgetIdsList(context.Background(), TEST_ID, "")

		assert.NotNil(t, err)
		assert.Equal(t, exceptions.ErrValidationError, err)
		assert.Equal(t, "", goalId)
	}

	{	// error found while getting goal by id
		mocks.DB.EXPECT().
			GetGoalRecordById(gomock.Any(), TEST_ID).
			Return(nil, ErrSomeError)
		
		goalId, err := service.UpdateBudgetIdsList(context.Background(), TEST_ID, TEST_BUDGET_ID)

		assert.NotNil(t, err)
		assert.Equal(t, ErrSomeError, err)
		assert.Equal(t, "", goalId)
	}

	{	// budget id already exists, no need to update
		goal := &models.Goal{
			GoalId: TEST_ID,
			BudgetIdList: []string{TEST_BUDGET_ID},
		}
		budget := &models.Budget{
			BudgetId: TEST_BUDGET_ID,
			GoalMap: []models.BudgetInputMap{{
				Id:   TEST_ID,
				Name: TEST_TITLE,
				CurrentAmount: 10,
				Amount: 20,
			}},
		}
		
		mocks.DB.EXPECT().
			GetGoalRecordById(gomock.Any(), TEST_ID).
			Return(goal, nil)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID).
			Return(budget, nil)

		goalId, err := service.UpdateBudgetIdsList(context.Background(), TEST_ID, TEST_BUDGET_ID)

		assert.Equal(t, TEST_ID, goalId)
		assert.Nil(t, err)
	}

	{	// error found while updating goal by id
		goal := &models.Goal{
			GoalId: TEST_ID,
		}

		mocks.DB.EXPECT().
			GetGoalRecordById(gomock.Any(), TEST_ID).
			Return(goal, nil)
		mocks.DB.EXPECT().
			UpdateGoalRecordById(gomock.Any(), TEST_ID, gomock.Any()).
			Return("", ErrSomeError)

		goalId, err := service.UpdateBudgetIdsList(context.Background(), TEST_ID, TEST_BUDGET_ID)

		assert.NotNil(t, err)
		assert.Equal(t, ErrSomeError, err)
		assert.Equal(t, "", goalId)
	}
}

// func UpdateGoalAmount(t *testing.T) {
// 	assert.Equal(t, 1, 1, "The two words should be the same.")
// }

func TestRemoveBudgetIdFromGoal(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestGoal(ctrl)

	{	// happy path
		goal := &models.Goal{
			GoalId: TEST_ID,
			BudgetIdList: []string{TEST_BUDGET_ID},
		}
		budget := &models.Budget{
			BudgetId: TEST_BUDGET_ID,
			GoalMap: []models.BudgetInputMap{{
				Id:   TEST_ID,
				Name: TEST_TITLE,
				CurrentAmount: 10,
				Amount: 20,
			}},
		}
		
		mocks.DB.EXPECT().
			GetGoalRecordById(gomock.Any(), TEST_ID).
			Return(goal, nil)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID).
			Return(budget, nil)
		mocks.DB.EXPECT().
			UpdateGoalRecordById(gomock.Any(), TEST_ID, gomock.Any()).
			Return(TEST_ID, nil)

		goalId, err := service.RemoveBudgetIdFromGoal(context.Background(), TEST_ID, TEST_BUDGET_ID)

		assert.Equal(t, TEST_ID, goalId)
		assert.Nil(t, err)
	}

	{	// validation error with goal id empty
		goalId, err := service.RemoveBudgetIdFromGoal(context.Background(), "", "")

		assert.NotNil(t, err)
		assert.Equal(t, exceptions.ErrValidationError, err)
		assert.Equal(t, "", goalId)
	}

	{	// validation error with budget id empty
		goalId, err := service.RemoveBudgetIdFromGoal(context.Background(), TEST_ID, "")

		assert.NotNil(t, err)
		assert.Equal(t, exceptions.ErrValidationError, err)
		assert.Equal(t, "", goalId)
	}

	{	// error found while getting goal by id
		mocks.DB.EXPECT().
			GetGoalRecordById(gomock.Any(), TEST_ID).
			Return(nil, ErrSomeError)
		
		goalId, err := service.RemoveBudgetIdFromGoal(context.Background(), TEST_ID, TEST_BUDGET_ID)

		assert.NotNil(t, err)
		assert.Equal(t, ErrSomeError, err)
		assert.Equal(t, "", goalId)
	}

	{	// budget id does not exist, no need to update
		goal := &models.Goal{
			GoalId: TEST_ID,
		}
		
		mocks.DB.EXPECT().
			GetGoalRecordById(gomock.Any(), TEST_ID).
			Return(goal, nil)

		goalId, err := service.RemoveBudgetIdFromGoal(context.Background(), TEST_ID, TEST_BUDGET_ID)

		assert.Equal(t, TEST_ID, goalId)
		assert.Nil(t, err)
	}

	{	// error found while updating goal
		goal := &models.Goal{
			GoalId: TEST_ID,
			BudgetIdList: []string{TEST_BUDGET_ID},
		}
		budget := &models.Budget{
			BudgetId: TEST_BUDGET_ID,
			GoalMap: []models.BudgetInputMap{{
				Id:   TEST_ID,
				Name: TEST_TITLE,
				CurrentAmount: 10,
				Amount: 20,
			}},
		}
		
		mocks.DB.EXPECT().
			GetGoalRecordById(gomock.Any(), TEST_ID).
			Return(goal, nil)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID).
			Return(budget, nil)
		mocks.DB.EXPECT().
			UpdateGoalRecordById(gomock.Any(), TEST_ID, gomock.Any()).
			Return("", ErrSomeError)

		goalId, err := service.RemoveBudgetIdFromGoal(context.Background(), TEST_ID, TEST_BUDGET_ID)

		assert.NotNil(t, err)
		assert.Equal(t, ErrSomeError, err)
		assert.Equal(t, "", goalId)
	}
}

func TestDeleteGoalById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service, mocks := createTestGoal(ctrl)
	{	// happy path
		expGoal := &models.Goal{
			GoalId: TEST_ID,
			CurrentAmount: 10,
			TargetAmount: 50,
			BudgetIdList: []string{TEST_BUDGET_ID},
		}
		expBudget := &models.Budget{
			BudgetId: TEST_BUDGET_ID,
			GoalMap: []models.BudgetInputMap{{
				Id:   TEST_ID,
				Name: TEST_TITLE,
				CurrentAmount: 10,
				Amount: 20,
			}},
		}
		mocks.DB.EXPECT().
			GetGoalRecordById(gomock.Any(), TEST_ID).
			Return(expGoal, nil)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID).
			Return(expBudget, nil).
			Times(2)
		mocks.DB.EXPECT().
			UpdateBudgetRecordById(gomock.Any(), TEST_BUDGET_ID, gomock.Any()).
			Return(TEST_BUDGET_ID, nil)
		mocks.DB.EXPECT().
			DeleteGoalRecordById(gomock.Any(), gomock.Any()).
			Return(TEST_ID, nil)

		id, err := service.DeleteGoalById(context.Background(), TEST_ID)

		assert.Equal(t, TEST_ID, id)
		assert.Nil(t, err)
	}

	{	// validation error	
		id, err := service.DeleteGoalById(context.Background(), "")

		assert.NotNil(t, err)
		assert.Equal(t, exceptions.ErrValidationError, err)
		assert.Equal(t, "", id)
	}

	{	// error from get goal record
		mocks.DB.EXPECT().
			GetGoalRecordById(gomock.Any(), TEST_ID).
			Return(nil, ErrSomeError)

		id, err := service.DeleteGoalById(context.Background(), TEST_ID)

		assert.NotNil(t, err)
		assert.Equal(t, err, ErrSomeError)
		assert.Equal(t, "", id)
	}

	{	// error while getting goals in remove goal map in budget
		expGoal := &models.Goal{
			GoalId: TEST_ID,
			CurrentAmount: 10,
			TargetAmount: 50,
			BudgetIdList: []string{TEST_BUDGET_ID},
		}
		expBudget := &models.Budget{
			BudgetId: TEST_BUDGET_ID,
			GoalMap: []models.BudgetInputMap{{
				Id:   TEST_ID,
				Name: TEST_TITLE,
				CurrentAmount: 10,
				Amount: 20,
			}},
		}
		mocks.DB.EXPECT().
			GetGoalRecordById(gomock.Any(), TEST_ID).
			Return(expGoal, nil)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID).
			Return(expBudget, nil)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID).
			Return(nil, ErrSomeError)

		id, err := service.DeleteGoalById(context.Background(), TEST_ID)

		assert.NotNil(t, err)
		assert.Equal(t, err, ErrSomeError)
		assert.Equal(t, "", id)
	}

	{	// error while updating budgets in remove goal map in budget
		expGoal := &models.Goal{
			GoalId: TEST_ID,
			CurrentAmount: 10,
			TargetAmount: 50,
			BudgetIdList: []string{TEST_BUDGET_ID},
		}
		expBudget := &models.Budget{
			BudgetId: TEST_BUDGET_ID,
			GoalMap: []models.BudgetInputMap{{
				Id:   TEST_ID,
				Name: TEST_TITLE,
				CurrentAmount: 10,
				Amount: 20,
			}},
		}
		mocks.DB.EXPECT().
			GetGoalRecordById(gomock.Any(), TEST_ID).
			Return(expGoal, nil)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID).
			Return(expBudget, nil).
			Times(2)
		mocks.DB.EXPECT().
			UpdateBudgetRecordById(gomock.Any(), TEST_BUDGET_ID, gomock.Any()).
			Return("", ErrSomeError)

		id, err := service.DeleteGoalById(context.Background(), TEST_ID)

		assert.NotNil(t, err)
		assert.Equal(t, err, ErrSomeError)
		assert.Equal(t, "", id)
	}

	{	// error found
		expGoal := &models.Goal{
			GoalId: TEST_ID,
			CurrentAmount: 10,
			TargetAmount: 50,
			BudgetIdList: []string{TEST_BUDGET_ID},
		}
		expBudget := &models.Budget{
			BudgetId: TEST_BUDGET_ID,
			GoalMap: []models.BudgetInputMap{{
				Id:   TEST_ID,
				Name: TEST_TITLE,
				CurrentAmount: 10,
				Amount: 20,
			}},
		}

		mocks.DB.EXPECT().
			GetGoalRecordById(gomock.Any(), TEST_ID).
			Return(expGoal, nil)
		mocks.DB.EXPECT().
			GetBudgetRecordById(gomock.Any(), TEST_BUDGET_ID).
			Return(expBudget, nil).
			Times(2)
		mocks.DB.EXPECT().
			UpdateBudgetRecordById(gomock.Any(), TEST_BUDGET_ID, gomock.Any()).
			Return(TEST_BUDGET_ID, nil)
		mocks.DB.EXPECT().
			DeleteGoalRecordById(gomock.Any(), gomock.Any()).
			Return("", ErrSomeError)

		id, err := service.DeleteGoalById(context.Background(), TEST_ID)

		assert.NotNil(t, err)
		assert.Equal(t, err, ErrSomeError)
		assert.Equal(t, "", id)
	}
}