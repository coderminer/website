package models

import (
	"log"
	"time"

	"github.com/globalsign/mgo"
)

const (
	dbhost    = "127.0.0.1:27017"
	authdb    = "admin"
	authuser  = "user"
	authpass  = "123456"
	mechanism = "SCRAM-SHA-1"
	timeout   = 60 * time.Second
	poollimit = 4096
	direct    = true
)

var (
	globalMS *mgo.Session
)

var s *mgo.Session

func init() {
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{dbhost},
		Direct:    direct,
		Timeout:   timeout,
		Source:    authdb,
		Username:  authuser,
		Password:  authpass,
		PoolLimit: poollimit,
		Mechanism: mechanism,
	}
	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalf("Create Session: %s\n", err)
	}
	globalMS = s
}

func connect(database, collection string) (*mgo.Session, *mgo.Collection) {
	ms := globalMS.Copy()
	c := ms.DB(database).C(collection)
	return ms, c
}

func IsEmpty(db, collection string) bool {
	ms, c := connect(db, collection)
	defer ms.Close()
	count, err := c.Count()
	if err != nil {
		log.Fatalln(err)
	}
	return count == 0
}

func Insert(db, collection string, docs interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Insert(docs)
}

func FindOne(db, collection string, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(selector).One(result)
}

func FindAll(db, collection string, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(selector).All(result)
}

func FindIter(db, collection string, selector interface{}) (*mgo.Iter, *mgo.Session) {
	ms, c := connect(db, collection)
	return c.Find(selector).Iter(), ms
}

func Count(db, collection string, query interface{}) (int, error) {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Count()
}

func Update(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	err := c.Update(selector, update)
	return err
}

func Upsert(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	_, err := c.Upsert(selector, update)
	return err
}

func Remove(db, collection string, selector interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Remove(selector)
}

func GetSession() *mgo.Session {
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{dbhost},
		Direct:    direct,
		Timeout:   timeout,
		Source:    authdb,
		Username:  authuser,
		Password:  authpass,
		PoolLimit: poollimit,
		Mechanism: mechanism,
	}

	s, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		log.Fatalf("Create Session: %s\n", err)
		panic(err)
	}
	s.SetMode(mgo.Monotonic, true)
	return s
}
