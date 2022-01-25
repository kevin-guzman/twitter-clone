package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = conectarDB()

const connectionString = 
	"mongodb+srv://kevinsizz:passUltraSecretXD1234.$@realmcluster.wflku.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

var clientOptions = options.Client().ApplyURI(connectionString)

func conectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión exitosa")
	return client
}

func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	return true
}
