package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/otaviobelfort/go/product_service/configs"
	_ "github.com/otaviobelfort/go/product_service/docs"
	"github.com/otaviobelfort/go/product_service/internal/infra/database"
	"github.com/otaviobelfort/go/product_service/internal/infra/webserver/handlers"
	httpswagger "github.com/swaggo/http-swagger"
)

// crie os principais parametros e decrições para o swagger, exemplo: @title, @version, @description
// @title Product Service API
// @version 1.0
// @description This is a sample server for Product Service.
// @termsOfService http://swagger.io/terms/

// @host localhost:8000
// @BasePath /
// @securityDefinitions.api_key ApiKeyAuth
// @in header
// @name Authorization

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		print("Error loading config\n", err, &config.DBHost)
	}
	//
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	db, err := configs.ConnectPostgres()
	if err != nil {
		panic(err)
	}

	//db.AutoMigrate(&database.Product{}, &database.User{})

	if err != nil {
		print("Error migrating database", err)
	}

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, config.TokenAuth, config.JWTExpiration)

	r := chi.NewRouter()

	// configure o LOGGER
	r.Use(middleware.Logger)
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/{id}", productHandler.GetProduct)
		r.Post("/", productHandler.CreateProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Get("/users/{id}", userHandler.GetUserEmail)
	r.Post("/users/token", userHandler.GetJwtToken)

	r.Get("/docs/*", httpswagger.Handler(httpswagger.URL("http://localhost:8000/docs/doc.json")))

	err = http.ListenAndServe(":8000", r)
	if err != nil {
		print("Error starting server", err)
	}
}
