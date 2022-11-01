package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Clase es el modelo con los datos de las clases
type Clase struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	NumClase    int8               `bson:"numClase" json:"numClase"`
	NombreClase int8               `bson:"nombreClase" json:"nombreClase"`
}
