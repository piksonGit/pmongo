package main

import (
	"fmt"

	"github.com/piksonGit/pmongo/db"
	"go.mongodb.org/mongo-driver/bson"
)

type Ranran struct {
	Mouse string
	Brain string
}
type Qijing struct {
	*Ranran
	Name    string
	Age     int
	Address string
}

func main() {
	uri := "mongodb://google:google123@120.27.190.31:27017/googlelinks"
	col := db.Conn(uri, "googlelinks", "test")
	//id := col.InsertOne(bson.M{"name": "亓京", "address": "山东省"})
	//fmt.Println(id)
	res := col.FindOne(bson.M{"_id": "5fe43eb5909d7b2945795f1d"})
	fmt.Println(res)
	//count := col.DeleteOne(bson.M{"name": "亓京"})
	//fmt.Printf("删除了%d条数据", count)
}

func testreturn() Ranran {
	var ranran Ranran = Ranran{"嘴巴", "舌头"}
	return ranran
}
