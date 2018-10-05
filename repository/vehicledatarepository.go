package repository

import (
	"fmt"
	"github.com/gocql/gocql"
	"wim-api/domain"
)

var Session *gocql.Session

func getCluster() *gocql.Session {
	var err error
	cluster := gocql.NewCluster("155.238.156.52")
	cluster.Keyspace = "wim"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	return Session
}

func AddVehicleData(data domain.VehicleData) (string, error) {
	fmt.Println(" **** Creating new data ****\n", data)
	defer getCluster().Close()
	if err := Session.Query("INSERT INTO vehicledata(id, weight,vehiclespeed) values(?,?,?)", data.ID, data.Weight, data.VehicleSpeed).Exec(); err != nil {

		fmt.Println("Error while inserting Vehicle Data", err)
		fmt.Println(err)
		return "", err
	}

	return "Success", nil
}
