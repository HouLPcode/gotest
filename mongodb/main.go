package main

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

func main() {
	url := "mongodb://platform-session:thoquon9eiCh@192.168.3.18:27017/platform-session"
	mgoInfo, err := mgo.ParseURL(url)
	if err != nil {
		panic(err)
	}
	mgoInfo.Timeout = time.Second * 5

	session, err := mgo.DialWithInfo(mgoInfo)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	session.SetPoolLimit(512)

	collection := session.DB("platform-session").C("hlp-test")

	// 注意：结构体一定要是首字母大写
	type mdata struct {
		Name string
		Age  int
		ID   bson.ObjectId `bson:"_id",inline`
	}
    d := &mdata{Name:"zhao",Age:16,ID:bson.NewObjectId()}
	err = collection.Insert(d)
	if err != nil{
		panic(err)
	}

	//err = collection.Insert(bson.M{"Name":"zhao","Age":13})
	//if err != nil{
	//	panic(err)
	//}
	cnt, _ := collection.Count()
	fmt.Println("collection count ", cnt)

	rnt := []mdata{}
	// 注意切片是值类型？？？
	collection.Find(bson.M{"name": "zhao"}).All(&rnt)
	fmt.Println(rnt)

	//collection.RemoveId("")

	//err = c.Find(query).One(&result)
	//fmt.Println(err)

	//session

	//time.Sleep(time.Second)
	session.Close()
}
