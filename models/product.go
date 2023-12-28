package models

type Product struct {
	// nama type `gorm:"type" json:"field"`
	Id         int    `gorm:"primaryKey" json:"id"`
	NamaProduk string `gorm:"type:varchar(255)" json:"nama_produk"`
	Deskripsi  string `gorm:"type:varchar(255)" json:"deskripsi"`
}
