package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Departamentos es el modelo con los datos de los departamentos
type Departamentos struct {
	ID                 primitive.ObjectID `bson:"_id" json:"id"`
	NumDepartamento    int8               `bson:"numDepartamento" json:"numDepartamento"`
	NombreDepartamento int8               `bson:"nombreDepartamento" json:"nombreDepartamento"`
}
