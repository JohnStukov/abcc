package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Productos es el modelo con los datos de los productos
type Productos struct {
	ID                 primitive.ObjectID `bson:"_id" json:"id"`
	SKU                int                `bson:"sku" json:"sku"`
	Articulo           string             `bson:"Articulo" json:"Articulo"`
	Marca              string             `bson:"Marca" json:"Marca"`
	Modelo             string             `bson:"Modelo" json:"Modelo"`
	Departamento       int8               `bson:"departamento" json:"departamento"`
	Clase              int8               `bson:"clase" json:"clase"`
	Familia            int8               `bson:"familia" json:"familia"`
	Stock              int8               `bson:"stock" json:"stock"`
	Cantidad           int8               `bson:"cantidad" json:"cantidad"`
	Descontinuado      int8               `bson:"descontinuado" json:"descontinuado"`
	FechaCreacion      time.Time          `bson:"fechaCreacion" json:"fechaCreacion"`
	FechaActualizacion time.Time          `bson:"fechaActualizacion" json:"fechaActualizacion"`
	FechaBaja          time.Time          `bson:"fechaBaja" json:"fechaBaja"`
}
