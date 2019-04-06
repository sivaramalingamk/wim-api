package repository

import "github.com/gocql/gocql"

var Session *gocql.Session

func getCluster() *gocql.Session {
	var err error
	cluster := gocql.NewCluster("scylla.scyla:9042")
	cluster.Keyspace = "wim"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	return Session
}
