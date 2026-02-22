package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "golang:golang@/goexpert?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	// products := []Product{
	// 	{Name: "Adidas Superstar II", Price: 799.00},
	// 	{Name: "Adidas courtbase suede", Price: 450.00},
	// 	{Name: "Adidas duramo RC2", Price: 350.00},
	// 	{Name: "Adidas courtbase alpha", Price: 499.00},
	// }

	// db.Create(products)

	var p Product
	db.Find(&p, 1)
	p.Price = 599
	db.Save(&p)

	var p2 Product
	db.Find(&p2, 2)
	db.Delete(&p2)
}
