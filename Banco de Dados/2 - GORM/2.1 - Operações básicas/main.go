package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {

	dsn := "golang:golang@/goexpert?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// cira automaticamente a tabela relacionada ao struct passado no parâmetro
	db.AutoMigrate(&Product{})

	// cria um produto
	// db.Create(&Product{
	// 	Name:  "Adidas Superstar II",
	// 	Price: 799.00,
	// })

	// a ORM também permite fazer inclusão em batch
	// products := []Product{
	// 	{Name: "Adidas courtbase suede", Price: 450.00},
	// 	{Name: "Adidas duramo RC2", Price: 350.00},
	// 	{Name: "Adidas courtbase alpha", Price: 499.00},
	// }

	// db.Create(products)

	// select one
	// var product Product
	// db.First(&product, 3)
	// fmt.Println(product)

	// var product2 Product
	// db.First(&product2, "name like ?", "%duramo%")
	// fmt.Println(product2)

	// select all
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products) -> Limit: limita a 2 linhas por resultado; Offset: faz a paginação dos resultados
	// db.Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	// select where
	// var products []Product
	// db.Where("price > ?", 350).Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	// update
	// var product Product
	// db.First(&product, 3)
	// product.Price = 399
	// db.Save(&product)

	// var product2 Product
	// db.First(&product2, 3)
	// fmt.Println(product2)

	// delete
	// var product Product
	// db.First(&product, 3)
	// db.Delete(&product)

	// var product2 Product
	// db.First(&product2, 3)
	// fmt.Println(product2)
}
