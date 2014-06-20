package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

var (
	mgoSession *mgo.Session
	dbName     = "goelia"
)

func getMgoSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
	}
	return mgoSession.Clone()
}

func withCollection(collectionName string, fn func(*mgo.Collection) error) error {
	session := getMgoSession()
	defer session.Close()
	collection := session.DB(dbName).C(collectionName)
	return fn(collection)
}

type User struct {
	Id       bson.ObjectId `json:"id,omitemty" bson:"_id"`
	WeixinId string        `json:"weixin_id" bson:"weixin_id"`
	Name     string        `json:"name,omitemty"`
	Contact  `json:"contact,omitempty" bson:"contact"`
	Chanels  []Chanel `json:"chanels,omitemty" bson:"chanels"`
}

type Contact struct {
	// Id     bson.ObjectId `json:"id,omitemty" bson:"_id"`
	Gender string `json:"gender" bson:"gender"`
	Mobile string `json:"mobile" bson:"mobile"`
}

type Chanel struct {
	Id    bson.ObjectId `json:"id,omitemty" bson:"_id"`
	Name  string        `json:"name,omitemty" bson:"name"`
	Image string        `json:"image,omitemty" bson:"image"`
	// ParentId bson.ObjectId `json:"parent_id,omitemty" bson:"parent_id"`
}

func (u *User) GetByWeixinId(weixinId string) error {
	searchFn := func(c *mgo.Collection) error {
		log.Println("weixinId:", weixinId)
		return c.Find(bson.M{"weixin_id": weixinId}).One(&u)
	}
	err := withCollection("user", searchFn)
	log.Println("search user:", u)
	return err

}

func (u *User) Create() error {
	// u.Id = bson.NewObjectId()
	// u.Contact.Id = bson.NewObjectId()
	fn := func(c *mgo.Collection) error {
		return c.Insert(&u)
	}
	return withCollection("user", fn)
}
