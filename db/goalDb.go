package db

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/pkg/exceptions"
	"github.com/sanchitdeora/budget-tracker/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *DatabaseImpl) GetAllGoalRecords(ctx context.Context) (*[]models.Goal, error) {
	cur, err := goalCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Println("error while fetching all goals, error:", err, ctx)
		return nil, err
	}

	var results []bson.M
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Println("error while fetching all goals, error:", err, ctx)
			return nil, err
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Println("error while fetching all goals, error:", err, ctx)
		return nil, err
	}
	cur.Close(ctx)

	var goals []models.Goal
	err = utils.ConvertBsonToStruct(results, &goals)
	if err != nil {
		log.Println("error while converting bson to struct, error:", err, ctx)
		return nil, err
	}

	fmt.Printf("Get All goal. Count of elements: %v\n", len(results))
	return &goals, nil
	
}

func (db *DatabaseImpl) GetGoalRecordById(ctx context.Context, key string) (*models.Goal, error) {
	var result bson.M
	var goal models.Goal
	filter := bson.M{GOAL_ID_KEY: key}
	err := goalCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Println("error while fetching goals by id:", key, "error:", err, ctx)
		return nil, err
	}
	if len(result) == 0 {
		log.Println("goal not found for id:", key, ctx)
		return nil, exceptions.ErrGoalNotFound
	}

	if err := utils.ConvertBsonToStruct(result, &goal); err != nil {
		log.Println("error while converting bson to struct, error:", err, ctx)
		return nil, err
	}
	

	return &goal, nil
	
}

func (db *DatabaseImpl) InsertGoalRecord(ctx context.Context, goal *models.Goal) (string, error) {
	if goal.GoalId == "" {
		goal.GoalId = GOAL_PREFIX + uuid.NewString()
	}	
	data := bson.D{
		{Key: GOAL_ID_KEY, Value: goal.GoalId},
		{Key: GOAL_NAME_KEY, Value: goal.GoalName},
		{Key: CURRENT_AMOUNT_KEY, Value: goal.CurrentAmount},
		{Key: TARGET_AMOUNT_KEY, Value: goal.TargetAmount},
		{Key: TARGET_DATE_KEY, Value: goal.TargetDate},
		{Key: BUDGET_ID_LIST_KEY, Value: goal.BudgetIdList},
	}

	result, err := goalCollection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created goal. ResultId: %v GoalId: %v\n", result.InsertedID, goal.GoalId)
	return goal.GoalId, err
}

func (db *DatabaseImpl) UpdateGoalRecordById(ctx context.Context, id string, goal *models.Goal) (string, error) {
	data := bson.D{{Key: "$set", 
		Value: bson.D{
			{Key: GOAL_NAME_KEY, Value: goal.GoalName},
			{Key: CURRENT_AMOUNT_KEY, Value: goal.CurrentAmount},
			{Key: TARGET_AMOUNT_KEY, Value: goal.TargetAmount},
			{Key: TARGET_DATE_KEY, Value: goal.TargetDate},
			{Key: BUDGET_ID_LIST_KEY, Value: goal.BudgetIdList},
		}},
	}
	filter := bson.D{{Key: GOAL_ID_KEY, Value: id}}

	result, err := goalCollection.UpdateOne(ctx, filter, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated goal. ModifiedCount: %v GoalId: %v\n", result.ModifiedCount, id)
	return id, err
}

func (db *DatabaseImpl) DeleteGoalRecordById(ctx context.Context, id string) (string, error) {
	filter := bson.D{{Key: GOAL_ID_KEY, Value: id}}

	result, err := goalCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted goal. DeletedCount: %v GoalId: %v\n", result.DeletedCount, id)
	return id, err
}