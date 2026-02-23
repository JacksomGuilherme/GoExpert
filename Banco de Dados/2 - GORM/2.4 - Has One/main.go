package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "golang:golang@/goexpert?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// category := Category{Name: "Casual"}
	// db.Create(&category)

	// category2 := Category{Name: "Esportivo"}
	// db.Create(&category2)

	// products := []Product{
	// 	{Name: "Adidas Superstar II", CategoryID: category.ID, Price: 799.00},
	// 	{Name: "Adidas courtbase suede", CategoryID: category.ID, Price: 450.00},
	// 	{Name: "Adidas duramo RC2", CategoryID: category2.ID, Price: 350.00},
	// 	{Name: "Adidas courtbase alpha", CategoryID: category.ID, Price: 499.00},
	// }

	// db.Create(&products)

	// db.Create(&SerialNumber{
	// 	Number:    "123456",
	// 	ProductID: 1,
	// })

	var products2 []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products2)
	for _, p := range products2 {
		fmt.Println(p)
	}

}
