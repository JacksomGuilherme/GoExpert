package main

import (
	"fmt"
	"net/http"

	"github.com/JacksomGuilherme/GoExpert/APIs/configs"
	_ "github.com/JacksomGuilherme/GoExpert/APIs/docs"
	"github.com/JacksomGuilherme/GoExpert/APIs/internal/entity"
	"github.com/JacksomGuilherme/GoExpert/APIs/internal/infra/database"
	"github.com/JacksomGuilherme/GoExpert/APIs/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with authentication.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Jacksom Guilherme
// @contact.url    https://www.github.com/JacksomGuilherme

// @license.name  Full Cycle License
// @license.url   http://www.fullcycle.com.br

// @host      localhost:8080
// @BasePath  /
// @securityDefinitions.apiKey  ApiKeyAuth
// @in header
// @name Authorization
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
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", config.TokenAuth))
	r.Use(middleware.WithValue("JwtExpiresIn", config.JWTExpiresIn))

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
	userHandler := handlers.NewUserHandler(userDB)

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)
	/* ====================================== */

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))
	http.ListenAndServe(":8080", r)
}
