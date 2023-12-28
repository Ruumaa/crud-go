package main

import (
	"gin-gorm/controllers/produkcontroller"
	"gin-gorm/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// connect dengan function yang ada di setup
	models.ConnectDatabase()

	// router.METHOD("/api/example", controller.Function)
	router.GET("/api/produk/", produkcontroller.GetAll) //func wajib capitalize
	router.GET("/api/produk/:id", produkcontroller.GetById)
	router.POST("/api/produk/", produkcontroller.Create)
	router.PUT("/api/produk/:id", produkcontroller.Update)
	router.DELETE("/api/produk/:id", produkcontroller.Delete)

	// router.RUN() untuk menjalankan aplikasi
	router.Run()

}
