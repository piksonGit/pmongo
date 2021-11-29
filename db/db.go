package db

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"time"

	"github.com/piksonGit/plog/plog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Col自定义一个collection结构体
type Col struct {
	*mongo.Collection
}

//init初始化log系统
func init() {
	plog.SetLog("./mongo_log.txt", "[pmongo]")
}

//Conn连接数据库
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

//Find函数
func (col *Col) Find(filter interface{}, opts ...*options.FindOptions) []bson.M {
	cursor, err := col.Collection.Find(context.Background(), filter, opts...)
	defer cursor.Close(context.Background())
	if err != nil {
		log.Println(err)
	}
	var results []bson.M
	if err = cursor.All(context.Background(), &results); err != nil {
		log.Println(err)
	}
	return results
}

//FindOne函数
func (col *Col) FindOne(filter bson.M, opts ...*options.FindOneOptions) bson.M {
	var result bson.M
	filter = build_id(filter)
	err := col.Collection.FindOne(context.Background(), filter, opts...).Decode(&result)
	if err != nil {
		log.Println(err)
	}
	return result
}

//将字符串类型的_id字段转换成ObjectID类型。
func build_id(filter bson.M) bson.M {
	t := reflect.TypeOf(filter["_id"])
	var ts string
	if t == nil {
		return filter
	}
	ts = t.String()
	if ts == "string" {
		filter["_id"], _ = primitive.ObjectIDFromHex(filter["_id"].(string))
	}

	return filter
}

//InsertOne函数
func (col *Col) InsertOne(data bson.M, opts ...*options.InsertOneOptions) interface{} {
	res, err := col.Collection.InsertOne(context.Background(), data, opts...)
	if err != nil {
		log.Println(err)
	}
	return res.InsertedID
}

//DeleteOne函数
func (col *Col) DeleteOne(filter bson.M, opts ...*options.DeleteOptions) interface{} {
	filter = build_id(filter)
	res, err := col.Collection.DeleteOne(context.Background(), filter, opts...)
	if err != nil {
		log.Println(err)
	}
	return res.DeletedCount
}

//UpdateOne函数
func (col *Col) UpdateOne(filter bson.M, update interface{}, opts ...*options.UpdateOptions) interface{} {
	filter = build_id(filter)
	res, err := col.Collection.UpdateOne(context.Background(), filter, update, opts...)
	if err != nil {
		log.Println(err)
	}
	return res
}

func ReadJson(filepath string) map[string]string {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("读取json配置错误")
	}
	var jsonContent map[string]string = make(map[string]string)
	// Unmarshal第二个参数必须是指针
	err = json.Unmarshal(data, &jsonContent)
	if err != nil {
		fmt.Println("json解析出错", err)
	}
	return jsonContent
}
