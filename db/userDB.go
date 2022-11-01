package db

import (
	"context"
	"fmt"
	"github.com/JohnStukov/examen1/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// Variables
var db = MongoCN.Database("abcc")
var col = db.Collection("productos")

// InsertoProducto es la parada final con la BD para insertar los datos del producto
func InsertoProducto(c models.Productos) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	c.ID = primitive.NewObjectID()
	result, err := col.InsertOne(ctx, c)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}

// ModificoProducto permite modificar los datos del producto dependiendo de que valores se envien, seran los que se modificaran
func ModificoProducto(c models.Productos, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	//armar el registro de la actualizacion de la db
	registro := make(map[string]interface{})
	if len(c.Articulo) > 0 {
		registro["Articulo"] = c.Articulo
	}
	if len(c.Marca) > 0 {
		registro["Marca"] = c.Marca
	}
	if len(c.Modelo) > 0 {
		registro["Modelo"] = c.Modelo
	}
	if c.Departamento > 0 {
		registro["departamento"] = c.Departamento
	}
	if c.Clase > 0 {
		registro["clase"] = c.Clase
	}
	if c.Familia > 0 {
		registro["familia"] = c.Familia
	}
	if c.Stock > 0 {
		registro["stock"] = c.Stock
	}
	if c.Cantidad > 0 {
		registro["cantidad"] = c.Cantidad
	}
	if c.Descontinuado == 1 {
		registro["descontinuado"] = c.Descontinuado

	}
	registro["fechaActualizacion"] = time.Now()
	updateString := bson.M{
		"$set": registro,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}
	//db
	_, err := col.UpdateOne(ctx, filtro, updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}

// BuscarProducto verifica que exista el producto en el sistema
func BuscarProducto(ID string) (models.Productos, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	var client models.Productos
	objID, _ := primitive.ObjectIDFromHex(ID)
	condicion := bson.M{
		"_id": objID,
	}
	err := col.FindOne(ctx, condicion).Decode(&client)
	client.Articulo = ""
	if err != nil {
		fmt.Println("registro no encontrado " + err.Error())
		return client, err
	}
	return client, nil
}

// RevisarSiExisteProducto recibe el sku como parametro y revisa si ya existe en la db
func RevisarSiExisteProducto(sku int) (models.Productos, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	condicion := bson.M{
		"sku": sku}
	var resultado models.Productos
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}

// ListoProductos lista todos los productos registrados en el sistema,
func ListoProductos() ([]*models.Productos, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var resultados []*models.Productos
	findOptions := options.Find()
	query := bson.M{}
	cursor, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return resultados, false
	}
	for cursor.Next(ctx) {
		var s models.Productos
		err := cursor.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return resultados, false
		}
		resultados = append(resultados, &s)
	}
	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return resultados, false
	}
	cursor.Close(ctx)
	return resultados, false
}

// BorradoLogicoProducto permite borrar el producto de forma logica
func BorradoLogicoProducto(ID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	//armar el registro de la actualizacion de la db
	registro := make(map[string]interface{})
	registro["fechaBaja"] = time.Now()
	updateString := bson.M{
		"$set": registro,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}
	//db
	_, err := col.UpdateOne(ctx, filtro, updateString)
	if err != nil {
		return err
	}
	return nil
}

// BorradoFisicoProducto recibe un id como parametro para eliminar un producto
func BorradoFisicoProducto(ID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(ID)
	condicion := bson.M{
		"_id": objID,
	}
	_, err := col.DeleteOne(ctx, condicion)
	return err
}
