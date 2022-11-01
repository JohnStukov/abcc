package routes

import (
	"encoding/json"
	"github.com/JohnStukov/examen1/db"
	"github.com/JohnStukov/examen1/models"
	"net/http"
	"time"
)

// RegistrarProducto crea el registro del producto en el sistema
func RegistrarProducto(w http.ResponseWriter, r *http.Request) {
	var p models.Productos
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	//revisamos si hay SKU's repetidos
	_, sku, _ := db.RevisarSiExisteProducto(p.SKU)
	if sku {
		http.Error(w, "Ya existe un producto con este sku  ", 400)
		return
	}
	if len(p.Articulo) < 0 {
		http.Error(w, "Campo del articulo vacio ", 400)
		return
	}
	if len(p.Marca) < 0 {
		http.Error(w, "Campo de la marca vacia ", 400)
		return
	}
	if len(p.Modelo) < 0 {
		http.Error(w, "Campo del modelo vacio ", 400)
		return
	}
	if (p.Departamento) <= 0 {
		http.Error(w, "Campo del departamento vacio ", 400)
		return
	}
	if (p.Clase) <= 0 {
		http.Error(w, "Campo de la Clase vacio ", 400)
		return
	}
	if (p.Familia) <= 0 {
		http.Error(w, "Campo de la familia vacio ", 400)
		return
	}
	if p.Stock < 0 {
		http.Error(w, "Stock no puede ser menor a 0 ", 400)
		return
	}
	if p.Cantidad > p.Stock {
		http.Error(w, "Cantidad no puede ser mayor al stock ", 400)
		return
	}
	p.FechaCreacion = time.Now()
	//revisar si fue correcto el registro
	_, status, err := db.InsertoProducto(p)
	if err != nil {
		http.Error(w, "Error al intentar el registro del producto "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se logro el registro del producto| ", 400)
		return
	}
	//se registra el producto
	w.WriteHeader(http.StatusCreated)
}

// EliminacionFisicoProducto Pide como parametro el id del producto para su eliminacion fisica
func EliminacionFisicoProducto(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID ", http.StatusBadRequest)
		return
	}
	_, err := db.BuscarProducto(ID)
	if err != nil {
		http.Error(w, "dato no encontrado "+err.Error(), 400)
		return
	}
	err = db.BorradoFisicoProducto(ID)
	if err != nil {
		http.Error(w, "ocurrio un error al eliminar el producto "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// EliminacionLogicaProducto Pide como parametro el sku del producto para su eliminacion logica
func EliminacionLogicaProducto(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID ", http.StatusBadRequest)
		return
	}
	_, err := db.BuscarProducto(ID)
	if err != nil {
		http.Error(w, "dato no encontrado "+err.Error(), 400)
		return
	}
	err = db.BorradoLogicoProducto(ID)
	if err != nil {
		http.Error(w, "ocurrio un error al eliminar el producto "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// ModificarProducto pide el id del producto para su modificacion de los datos
func ModificarProducto(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var u models.Productos
	//grabamos el body
	err := json.NewDecoder(r.Body).Decode(&u)
	//revisamos q el body no esta vacio o que no contiene errores de diseño el documento
	if err != nil {
		http.Error(w, "datos incorrectos o peticion vacia. "+err.Error(), 400)
		return
	}
	if len(ID) < 1 {
		http.Error(w, "debe enviar el parametro ID ", http.StatusBadRequest)
		return
	}
	user, err := db.BuscarProducto(ID)
	if err != nil {
		http.Error(w, "error al encontrar el registro "+err.Error(), 400)
		return
	}
	if len(user.Marca) == 0 {
		http.Error(w, "Campo del nombre vacio ", 400)
		return
	}
	if len(user.Modelo) == 0 {
		http.Error(w, "Campo del apellido vacio ", 400)
		return
	}
	if (user.Departamento) == 0 {
		http.Error(w, "El email de user es requerido ", 400)
		return
	}
	var status bool
	status, err = db.ModificoProducto(u, ID)
	if err != nil {
		http.Error(w, "ocurrio un error al modificar el registro, intente de nuevo "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "no se logro modificar el registro del producto "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// VerProducto pide el id del producto para mostrar su datos
func VerProducto(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	client, err := db.BuscarProducto(ID)
	if err != nil {
		http.Error(w, "error al encontrar el registro "+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)
}

// ListarProducto muestra todos los clientes registrados
func ListarProducto(w http.ResponseWriter, r *http.Request) {
	result, status := db.ListoProductos()
	if status != false {
		http.Error(w, "error al leer productos ", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
