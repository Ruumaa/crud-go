// untuk konek ke database
// import (
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
//   )

//   dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
//   db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// handle connect database
	database, err := gorm.Open(postgres.Open("host=localhost user=postgres password=1234 dbname=go_restapi port=5432 sslmode=disable"))
	if err != nil {
		panic(err)
	}
	// membuat auto migrate
	database.AutoMigrate(&Product{}) //& adalah pointer ke file product

	// memasukkan hasil dari database migrate ke DB
	DB = database
}
