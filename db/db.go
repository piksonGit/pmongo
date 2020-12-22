package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Col struct {
	collection mongo.Collection
	Ctx        context.Context
	Result     bson.M
	Conditon   bson.M
	Data       bson.M
}

func (col *Col) Find() {

}
