package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MongoClient = ConnectDB() //retorna la conexion
var ClientOptions = options.Client().ApplyURI("mongodb+srv://db_user:user2023@cluster0.voam20v.mongodb.net/go_twitter")

//ConnectDB to connect with mongodb
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), ClientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	//una vez que se tiene el cliente se realiza un ping para ver si la conexcion esta disponible

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Success Connection")

	return client
}

//CheckConnection - to check the connection with mongodb
func CheckConnection() int {

	err := MongoClient.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}

	return 1
}
