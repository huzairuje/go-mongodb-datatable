package config

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var conn *mongo.Database

func CreateConnectionDB() *mongo.Database {
	dbHost := "localhost"
	dbPort := "27017"
	dbUser := ""
	dbPass := ""
	dbName := "go_datatable"
	dbAuth := ""

	credential := options.Credential{
		AuthSource: dbAuth,
		Username:   dbUser,
		Password:   dbPass,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	URI := fmt.Sprintf("mongodb://%s:%s/?connect=direct", dbHost, dbPort)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI).SetAuth(credential))

	if err != nil {
		logrus.Errorf("failed to open database: %v", err)
	}

	db := client.Database(dbName)
	SetConnection(db)
	return db
}

//GetConnection : Get Available Connection
func GetConnection() (db *mongo.Database, err error) {
	return conn, err
}

//SetConnection : Set Available Connection
func SetConnection(connection *mongo.Database) {
	conn = connection
}