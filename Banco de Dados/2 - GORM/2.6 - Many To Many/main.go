package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories"`
	gorm.Model
}

func main() {
	dsn := "golang:golang@/goexpert?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	category := Category{Name: "Casual"}
	db.Create(&category)

	category2 := Category{Name: "Esportivo"}
	db.Create(&category2)

	category3 := Category{Name: "Calçados"}
	db.Create(&category3)

	products := []Product{
		{Name: "Adidas Superstar II", Categories: []Category{category3, category}, Price: 799.00},
		{Name: "Adidas courtbase suede", Categories: []Category{category3, category}, Price: 450.00},
		{Name: "Adidas duramo RC2", Categories: []Category{category3, category2}, Price: 350.00},
		{Name: "Adidas courtbase alpha", Categories: []Category{category3, category}, Price: 499.00},
	}

	db.Create(&products)

	var categories []Category

	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, c := range categories {
		fmt.Println(c.Name, ":")

		for _, p := range c.Products {
			fmt.Println("- ", p.Name)
		}
	}

}
