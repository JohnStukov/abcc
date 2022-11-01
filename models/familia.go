package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Familia es el modelo con los datos de las familias de productos
type Familia struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	NumFamilia    int8               `bson:"numFamilia" json:"numFamilia"`
	NombreFamilia int8               `bson:"nombreFamilia" json:"nombreFamilia"`
}
