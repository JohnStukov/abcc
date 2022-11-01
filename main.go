package main

import (
	"github.com/JohnStukov/examen1/db"
	"github.com/JohnStukov/examen1/handlers"
	"log"
)

// main ejecuta la API y pasa el control al handler
func main() {
	if db.RealizaPing() == 0 {
		log.Fatal("Sin conexion a la db")
		return
	}
	handlers.Manejador()
}
