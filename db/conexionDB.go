package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// MongoCN es el objeto de conexion a la db
var MongoCN = ConectDB()
var clientOption = options.Client().ApplyURI("mongodb://root:123456@localhost:27017")

// docker run -d --name mongo -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=123456 -p 27017:27017 mongo

// ConetarDB ES LA FUNCION Q ME PERMITE CONECTAR A LA BASE DE DATOS
func ConectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa con la DB")
	return client
}

// RealizaPing realiza un ping a la db para verificar que este activo el servicio
func RealizaPing() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
