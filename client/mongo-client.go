package client

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Data struct {
	Key    string			`bson:"key"`
	Value  string			`bson:"value"`
}

type MongoCli struct {
	Addr	string
	conn    *mgo.Session
	Mode    string
}

func (this *MongoCli) Conn() {
	if this.conn == nil {
		var err error
		this.conn, err = mgo.Dial(this.Addr)

		if this.Mode == "0" {
			this.conn.SetMode(mgo.Eventual, true)
		} else if this.Mode == "1" {
			this.conn.SetMode(mgo.Monotonic, true)
		} else if this.Mode == "2" {
			this.conn.SetMode(mgo.Strong, true)
		}

		if err != nil {
			panic(err) //直接终止程序运行
		}
	}
}

func (this *MongoCli) SetV(key string, value string) error{

	c := this.conn.DB("test").C("test_data")
	err := c.Insert(&Data{Key:key, Value:value})

	if err != nil {
		panic(err)
	}

	return err
}

func (this *MongoCli) GetV(key string) string {

	c := this.conn.DB("test").C("test_data")

	result := Data{}
	err := c.Find(bson.M{"key": key}).One(&result)
	if err != nil {
		panic(err)
	}

	return result.Value
}


func (this *MongoCli) Close() {
	this.conn.Close()
}