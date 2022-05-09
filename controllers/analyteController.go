package controllers

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/Laranjazda/api-rest-go-with-gin/database"
	"github.com/Laranjazda/api-rest-go-with-gin/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func GetLabels(index, age int, sex, language string) {
func GetLabels(c *gin.Context) {
	index, err := strconv.Atoi(c.Params.ByName("index"))
	if err != nil {
		log.Panic(err)
	}
	language := c.Params.ByName("language")
	age, err := strconv.Atoi(c.Params.ByName("age"))
	if err != nil {
		log.Panic(err)
	}
	sex := c.Params.ByName("sex")

	l := ReadLabels(index, language)

	l.Limits = ReadLimits(index, age, language, sex)
}

func ReadLabels(index int, language string) models.Label {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var label models.Label

	matchStage := bson.D{{"$match", bson.D{{"index", index}}}}
	replaceRootStages := bson.D{{"$replaceRoot", bson.D{{"labels" + ".", language}}}}
	projectStage := bson.D{{"$project", bson.D{{"limits", 0}}}}

	pipeline := mongo.Pipeline{matchStage, replaceRootStages, projectStage}

	answer, err := database.AlunosCollection().Aggregate(ctx, pipeline, &options.AggregateOptions{})
	if err != nil {
		panic(err.Error())
	}

	// showsWithInfo := models.Label{}
	if err = answer.All(ctx, &label); err != nil {
		panic(err)
	}

	return label
}

func ReadLimits(index, age int, language, sex string) models.Limits {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var limits models.Limits

	matchStage := bson.D{{"$match", limits.Index == index}}
	replaceRootStages := bson.D{{"$replaceRoot", bson.D{{"labels" + ".", language}}}}
	unwindStage := bson.D{{"$unwind", "$use"}}
	match2Stage := bson.D{{"$match", bson.D{{"limits.sex", sex}, {"limits.min_age", age}, {"limits.max_age", age}}}}
	sortStage := bson.D{{"$sort", bson.D{{"rating", 1}, {"limits.index", 1}}}, {"$sort", bson.D{{"rating", 1}, {"limits.index", 1}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$limits"}}}}

	pipeline := mongo.Pipeline{matchStage, replaceRootStages, unwindStage, match2Stage, sortStage, groupStage}

	answer, err := database.AlunosCollection().Aggregate(ctx, pipeline, &options.AggregateOptions{})
	if err != nil {
		panic(err.Error())
	}

	// showsWithInfo := models.Limits{}
	if err = answer.All(ctx, &limits); err != nil {
		panic(err)
	}

	return limits
}
