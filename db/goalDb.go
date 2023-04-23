package db

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/sanchitdeora/budget-tracker/models"
	"github.com/sanchitdeora/budget-tracker/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllGoals(ctx context.Context, goals *[]models.Goal) error {
	cur, err := goalCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return err
	}

	var results []bson.M
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(ctx)

	err = utils.ConvertBsonToStruct(results, goals)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Get All goal. Count of elements: %v\n", len(results))
	return nil
	
}

func GetGoalRecordById(ctx context.Context, key string) (*models.Goal, error) {
	var result bson.M
	var goal models.Goal
	filter := bson.M{GOAL_ID_KEY: key}
	err := goalCollection.FindOne(ctx, filter).Decode(&result)
	if len(result) == 0 {
		return nil, nil
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := utils.ConvertBsonToStruct(result, &goal); err != nil {
		log.Println(err)
		return nil, err
	}

	return &goal, nil
	
}

func InsertGoalRecord(ctx context.Context, goal models.Goal) (string, error) {
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

func UpdateGoalRecordById(ctx context.Context, id string, goal models.Goal) (string, error) {
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

func DeleteGoalRecordById(ctx context.Context, id string) (string, error) {
	filter := bson.D{{Key: GOAL_ID_KEY, Value: id}}

	result, err := goalCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted goal. DeletedCount: %v GoalId: %v\n", result.DeletedCount, id)
	return id, err
}