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
	"github.com/go-chi/jwtauth"
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

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	/* ========== PRODUCT ENDPOINTS ========== */
	productDB := database.NewProductRepository(db)
	productHandler := handlers.NewProductHandler(productDB)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})
	/* ======================================= */

	/* =========== USER ENDPOINTS =========== */
	userDB := database.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userDB, config.TokenAuth, config.JWTExpiresIn)

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)
	/* ====================================== */

	http.ListenAndServe(":8080", r)
}
