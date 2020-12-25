package db

import (
	"context"
	"log"
	"time"

	"github.com/piksonGit/plog/plog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Col struct {
	*mongo.Collection
}

func init() {
	plog.SetLog("./mongo_log.txt", "[pmongo]")
}
func Conn(uri string, databaseName string, collectionName string) Col {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(databaseName).Collection(collectionName)
	col := Col{collection}
	return col
}

func (col *Col) Find(filter interface{}, opts ...*options.FindOptions) []bson.M {
	cursor, err := col.Collection.Find(context.TODO(), filter, opts...)
	if err != nil {
		log.Println(err)
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Println(err)
	}
	return results
}

func (col *Col) FindOne(filter bson.M, opts ...*options.FindOneOptions) bson.M {
	var result bson.M
	if filter["_id"] != nil {
		filter["_id"], _ = primitive.ObjectIDFromHex(filter["_id"].(string))
	}
	err := col.Collection.FindOne(context.TODO(), filter, opts...).Decode(&result)
	if err != nil {
		log.Println(err)
	}
	return result
}
func (col *Col) InsertOne(data bson.M, opts ...*options.InsertOneOptions) interface{} {
	res, err := col.Collection.InsertOne(context.TODO(), data, opts...)
	if err != nil {
		log.Println(err)
	}
	return res.InsertedID
}

func (col *Col) DeleteOne(filter interface{}, opts ...*options.DeleteOptions) interface{} {
	res, err := col.Collection.DeleteOne(context.TODO(), filter, opts...)
	if err != nil {
		log.Println(err)
	}
	return res.DeletedCount
}
func (col *Col) UpdateOne(filter interface{}, update interface{}, opts *options.UpdateOptions) interface{} {
	res, err := col.Collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Println(err)
	}
	return res
}
