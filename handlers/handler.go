package handlers

import (
	"github.com/JohnStukov/examen1/middleW"
	"github.com/JohnStukov/examen1/routes"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"os"
)

func Manejador() {
	router := mux.NewRouter()
	router.Use(
		middleware.Logger,    //log every http request
		middleware.Recoverer, //recover if panic occurs
	)
	//-----------------------------RUTAS------------------------------------//
	//-----------------------------Productos----------------------------------//
	router.HandleFunc("/producto/", middleW.RevisarDB(routes.RegistrarProducto)).Methods("POST")
	router.HandleFunc("/producto/", middleW.RevisarDB((routes.ListarProducto))).Methods("GET")
	router.HandleFunc("/producto/id", middleW.RevisarDB((routes.VerProducto))).Methods("GET")
	router.HandleFunc("/producto/id", middleW.RevisarDB((routes.ModificarProducto))).Methods("PUT")
	router.HandleFunc("/producto/l/", middleW.RevisarDB((routes.EliminacionLogicaProducto))).Methods("PUT")
	router.HandleFunc("/producto/f/", middleW.RevisarDB((routes.EliminacionFisicoProducto))).Methods("DELETE")
	//----------------------------------------------------------------------//
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}
	//da acceso desde cualquier IP
	handler := cors.AllowAll().Handler(router)
	//pone a escuchar el puerto elegido y pasa el control al handler
	http.ListenAndServe(":"+PORT, handler)
}
