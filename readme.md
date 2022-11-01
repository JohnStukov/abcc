# Prueba tecnica

## Comenzando

_Estas instrucciones te permitirán obtener una copia del proyecto en funcionamiento en tu máquina local para propósito de evaluación._


### Pre-requisitos

```
Golang
Postman
```

### Instalación

* 1.- Instalar el lenguaje de programación golang, disponible en su página web: [Golang](https://golang.org)
* 2.- Para corroborar la instalación, ejecutar el siguiente comando:
```
go -v
```
Si el comando muestra la version de go, la instalación esta correcta.

* 3.- Instalar Postman desde su pagina: [Postman](https://postman.com)
* 4.- Docker para ejecutar la base de datos [Docker](https://docker.com)


## Despliegue

* Clonar este repositorio en tu equipo y ejecutarlo.
* La base de datos en MongoDB se crea en automatico cuando se ejecuta el servidor 

## Docker
* Este comando creara un contenedor en docker de MongoDB
* docker run -d --name mongo -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=123456 -p 27017:27017 mongo


## Endpoints

* Para probar los endpoints del proyecto, se sugiere ocupar Postman.
* Los endpoints del proyecto se encuentran en el archivo [abcc.api_collection.json](abcc.api_collection.json)
* Importar el archivo [abcc.api_collection.json](abcc.api_collection.json) en el programa [Postman](https://postman.com)

### localhost:8081/producto/ METOHOD POST 
* Registra un producto.

### localhost:8081/producto/ METHOD GET
* Consulta todos los productos existentes

### localhost:8081/producto/id METHOD GET
* Consulta los datos de un producto en especifico

### localhost:8081/producto/id METHOD PUT
* Modifica un producto

### localhost:8081/producto/l/ METHOD PUT
* Elimina de forma logica un producto

### localhost:8081/producto/f/ METHOD DELETE
* Elimina de forma fisica un producto

## Construido con

* [Goland](https://www.jetbrains.com/es-es/go/) - IDE de programación para Golang


## Autores ✒️

* **Juan Salomón Sains Hdez** - *Documentación y programación* - [JohnStukov](https://github.com/JohnStukov)



---
⌨️ con ❤️ por [JohnStukov](https://github.com/JohnStukov) 😊