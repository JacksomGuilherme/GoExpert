package main

import (
	"fmt"
	"net/http"

	"github.com/JacksomGuilherme/GoExpert/APIs/configs"
	"github.com/JacksomGuilherme/GoExpert/APIs/internal/entity"
	"github.com/JacksomGuilherme/GoExpert/APIs/internal/entity/infra/database"
	"github.com/JacksomGuilherme/GoExpert/APIs/internal/entity/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config := configs.LoadConfig(".")

	dsn := fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProductRepository(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products", productHandler.GetProducts)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	http.ListenAndServe(":8080", r)
}
