package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Williberto14/powercrud/conf"
	"github.com/Williberto14/powercrud/model"
	"github.com/Williberto14/powercrud/utils"
	"github.com/gin-gonic/gin"
)

var db = conf.ConfDatabase()

func main() {

	IsDevelopEnv := flag.Bool("devEnv", false, "")
	flag.Parse()
	log.Println("Deployed for development environment:", *IsDevelopEnv)

	// r := gin.Default()

	// authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	// 	"user1": "root",
	// 	"user2": "1234",
	// }))

	// authorized.GET("/volumes", GetVolumes)

	// r.POST("/products/create", CreateProduct())
	// r.GET("/products/list", GetProducts)
	// r.GET("/products/find/:id", GetProduct)
	// r.PUT("/products/update/:id", UpdateProduct)
	// r.DELETE("/products/delete/:id", DeleteProduct)

	// r.Run(":9098")
}

//=====================================================================================

func CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := model.DefaultResponse{}
		product := model.Product{}

		errBind := c.BindJSON(&product)
		if errBind != nil {
			res.NewDefault("WRONG_DATA", "error in the data structure")
			c.JSON(http.StatusBadRequest, res)
			return
		}

		product.ProductId = utils.GenerateUuid()
		db.Create(&product)
		res.NewDefault("OK", "product created succesfully")
		c.JSON(200, res)
		return
	}
}

func GetProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product model.Product
	if err := db.Where("product_id = ?", id).First(&product).Error; err != nil {
		c.JSON(404, "No existe un producto con id: "+id)
		fmt.Println(err)
	} else {
		c.JSON(200, product)
	}
}

func GetProducts(c *gin.Context) {
	var products []model.Product
	if err := db.Find(&products).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, products)
	}
}

func UpdateProduct(c *gin.Context) {
	var product model.Product
	id := c.Params.ByName("id")
	if err := db.Where("product_id  = ?", id).First(&product).Error; err != nil {
		c.JSON(404, "Error! No se pudo actualizar el producto con id: "+id)
		fmt.Println(err)
	} else {
		c.BindJSON(&product)
		db.Save(&product)
		c.JSON(200, product)
	}
}

func DeleteProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product model.Product
	d := db.Where("product_id  = ?", id)
	if err := db.Where("product_id  = ?", id).First(&product).Error; err != nil {
		c.JSON(404, "No existe un producto con id: "+id)
	} else {
		d.Delete(&product)
		c.JSON(200, gin.H{"Producto con id #" + id: "eliminado con exito"})
	}
}

func GetVolumes(c *gin.Context) {
	jsonFile, _ := ioutil.ReadFile("./volumes/volumen_list.json")
	var data interface{}
	json.Unmarshal(jsonFile, &data)
	c.JSON(http.StatusOK, data)
}
