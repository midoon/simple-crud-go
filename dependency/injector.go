package dependency

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/midoon/simple-crud-go/app"
	"github.com/midoon/simple-crud-go/controller"
	"github.com/midoon/simple-crud-go/middleware"
	"github.com/midoon/simple-crud-go/repository"
	"github.com/midoon/simple-crud-go/service"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func InitializedServer() *http.Server {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)
	authMiddleware := middleware.NewAuthMiddleware(router)

	server := NewServer(authMiddleware)
	return server
}