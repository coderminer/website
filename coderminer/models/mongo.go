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

var s *mgo.Session

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
