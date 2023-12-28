package produkcontroller

import (
	"gin-gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//membuat function untuk manipulasi database

func GetAll(c *gin.Context) {
	// membuat slice untuk menampung
	var products []models.Product

	// mengambil semua product yang ada di database menggunakan pointer ke products
	models.DB.Find(&products)
	// mengirim produk ke response JSON
	// c.JSON(http.status,gin.H{"data": products} )
	c.JSON(http.StatusOK, gin.H{"produk": products})
}

func GetById(c *gin.Context) {
	var produk models.Product
	id := c.Param("id")

	if err := models.DB.First(&produk, id).Error; err != nil {
		// switch ke error
		switch err {
		// jika casenya notFound
		case gorm.ErrRecordNotFound:
			// mengirimkan response json jika data kosong
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Produk tidak ditemukan"})
			// return agar keluar
			return
		default:
			// response jika error database
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	// kirim response jika berhasil
	c.JSON(http.StatusOK, gin.H{"produk": produk})
}

func Create(c *gin.Context) {

	var produk models.Product

	// ShouldBindJSON ->  untuk mengikat atau mengisi struktur data (produk dalam contoh Anda) dari data JSON yang diterima dalam badan permintaan HTTP.
	// jadi jika fieldnya tidak sesuai pointer maka akan error
	if err := c.ShouldBindJSON(&produk); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// menambahkan ke database
	models.DB.Create(&produk)
	// kirim response yang sudah dikirim ke database
	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dibuat", "produk": produk})
}

func Update(c *gin.Context) {
	var produk models.Product
	id := c.Param("id")

	// handle error
	if err := c.ShouldBindJSON(&produk); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// handle jika tidak ada perubahan atau tidak ditemukan idnya
	if models.DB.Model(&produk).Where("id = ?", id).Updates(&produk).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menemukan produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update produk berhasil"})
}

func Delete(c *gin.Context) {
	var produk models.Product
	id := c.Param("id")

	// input := map[string]string{"id": ""}

	if models.DB.Where("id = ?", id).Delete(&produk).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menemukan produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
