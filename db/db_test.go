package db
import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"fmt"
)
func BenchmarkFind(b *testing.B) {
	uri := "mongodb://google:google123@120.27.190.31:27017/googlelinks"
	col := Conn(uri, "googlelinks", "test")
	//id := col.InsertOne(bson.M{"name": "亓京", "address": "山东省"})
	//fmt.Println(id)
	res := col.FindOne(bson.M{"_id": "5fe43eb5909d7b2945795f1d"})
	fmt.Println(res)
}