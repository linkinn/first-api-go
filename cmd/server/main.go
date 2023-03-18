package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/linkinn/first-api/configs"
	_ "github.com/linkinn/first-api/docs"
	"github.com/linkinn/first-api/internal/entity"
	"github.com/linkinn/first-api/internal/infra/database"
	"github.com/linkinn/first-api/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title Go Expert API
// @version 1.0
// @description Product API
// @termsOfService http://swagger.io/terms/
//
// @contact.name Fillipi Nascimento
// @contact.url https://fillipinascimento.com
// @contact.email administrator@fillipinascimento.com
//
// @license.name Fillipi Nascimento License
// @license.url https://fillipinascimento.com
//
// @host localhost:9090
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// -> Environments
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	// -> Databse Config
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	r := chi.NewRouter()

	// -> Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("JwtExperiesIn", configs.JWTExperesIn))

	// -> User Routes
	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)
	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)

	// -> Product Routes
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandlerr(productDB)
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	// -> Swag Routes
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:9090/docs/doc.json")))

	// -> Starting Server
	http.ListenAndServe(":9090", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
