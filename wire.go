//go:build wireinject
// +build wireinject

package main

// NOTE: ini terdapat error di bagian generate dependency menggunakan google wire, mohon riset lebih lanjut

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"github.com/midoon/simple-crud-go/app"
	"github.com/midoon/simple-crud-go/controller"
	"github.com/midoon/simple-crud-go/middleware"
	"github.com/midoon/simple-crud-go/repository"
	"github.com/midoon/simple-crud-go/service"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRespository), new(repository.CategoryRespositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(controller.CategoryControllerImpl)),	
)

func InitializedServer() *http.Server{
	wire.Build(app.NewDB, validator.New, categorySet, app.NewRouter, wire.Bind(new(http.Handler), new(*httprouter.Router)),middleware.NewAuthMiddleware, NewServer)
	return nil
}

