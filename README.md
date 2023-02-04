# API REST de Productos

## prueba-técnica

============================================

**Descripción:** Este proyecto tiene como objetivo desarrollar una api rest en el lenguaje Golang que exponga endpoints CRUD para gestionar productos y volúmenes.

=============================================

**Tecnologías usadas:**

- Golang/Go
- Gin Framework
- GORM ORM
- Ginkgo & Gomega
- MariaDB
- Postman

=============================================

**Instalación:** Después de clonar el repositorio ejecutar los siguiente comandos en la terminal

`go mod tidy`

`go run main.go`

**Dependencias**

- GORM y Driver SQLite

  `go get -u gorm.io/gorm`

  `go get -u gorm.io/driver/sqlite`
- Gin

  `go get -u github.com/gin-gonic/gin`
- Datatypes GORM

  `go get -u gorm.io/datatypes`
- Ginkgo y Gomega

  `go get github.com/onsi/ginkgo/ginkgo`

  `go get github.com/onsi/gomega/...`

=============================================

### La API se levantara en el puerto :9098

**Los endpoints expuestos al levantar la aplicación son:**

- **POST** /products/create
- **GET** /products/list
- **GET** /products/find/:id
- **PUT** /products/update/:id
- **DELETE** /products/delete/:id

---

**Para acceder al siguiente endpoint deberá autenticarse:**

    Usuario     Contraseña
    "user1":    "root",
    "user2":    "1234",

- **GET** /volumes

## Comando docker base de datos

```
docker run --name prueba-tecnica -p 127.0.0.1:3310:3306 -e MYSQL_ROOT_PASSWORD=qwerty -e MYSQL_DATABASE=prueba_tecnica -d mariadb
```
