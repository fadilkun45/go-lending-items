// @title Loans Item API
// @version 1.0
// @description REST API untuk sistem peminjaman barang P2P
// @host localhost:8000
// @BasePath /
package main

import (
	_ "loans-item-go/docs"

	"loans-item-go/config"
	"loans-item-go/controller"
	"loans-item-go/helper"
	"loans-item-go/repository"
	"loans-item-go/service"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {

	db := config.OpenConnection()

	// repository
	userRepo := repository.NewUserRepositoryImpl()
	categoryRepo := repository.NewCategoryRepositoryImpl()
	itemRepo := repository.NewItemRepositoryImpl()
	loanRepo := repository.NewLoanRepositoryImpl()
	loanItemRepo := repository.NewLoanItemRepositoryImpl()

	// service
	userSvc := service.NewUserServiceImpl(userRepo, db)
	categorySvc := service.NewCategoryServiceImpl(categoryRepo, db)
	itemSvc := service.NewItemServiceImpl(itemRepo, db)
	loanSvc := service.NewLoanServiceImpl(loanRepo, itemRepo, loanItemRepo, db)

	// controller
	userCtrl := controller.NewUserController(userSvc)
	categoryCtrl := controller.NewCategoryController(categorySvc)
	itemCtrl := controller.NewItemController(itemSvc)
	loanCtrl := controller.NewLoanController(loanSvc)

	// router
	mux := http.NewServeMux()

	// swagger
	mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	// user routes
	mux.HandleFunc("POST /api/users/register", userCtrl.Register)
	mux.HandleFunc("POST /api/users/login", userCtrl.Login)
	mux.HandleFunc("GET /api/users", userCtrl.FindAll)
	mux.HandleFunc("GET /api/users/{id}", userCtrl.FindById)
	mux.HandleFunc("PUT /api/users/{id}", userCtrl.Update)
	mux.HandleFunc("DELETE /api/users/{id}", userCtrl.Delete)

	// category routes
	mux.HandleFunc("POST /api/categories", categoryCtrl.Create)
	mux.HandleFunc("GET /api/categories", categoryCtrl.FindAll)
	mux.HandleFunc("GET /api/categories/{id}", categoryCtrl.FindById)
	mux.HandleFunc("PUT /api/categories/{id}", categoryCtrl.Update)
	mux.HandleFunc("DELETE /api/categories/{id}", categoryCtrl.Delete)

	// item routes
	mux.HandleFunc("POST /api/items", itemCtrl.Create)
	mux.HandleFunc("GET /api/items", itemCtrl.FindAll)
	mux.HandleFunc("GET /api/items/{id}", itemCtrl.FindById)
	mux.HandleFunc("PUT /api/items/{id}", itemCtrl.Update)
	mux.HandleFunc("DELETE /api/items/{id}", itemCtrl.Delete)

	// loan routes
	mux.HandleFunc("POST /api/loans", loanCtrl.Create)
	mux.HandleFunc("GET /api/loans", loanCtrl.FindAll)
	mux.HandleFunc("GET /api/loans/{id}", loanCtrl.FindById)
	mux.HandleFunc("GET /api/loans/borrower/{borrower_id}", loanCtrl.FindByBorrowerID)
	mux.HandleFunc("PUT /api/loans/{id}", loanCtrl.Update)
	mux.HandleFunc("DELETE /api/loans/{id}", loanCtrl.Delete)

	server := &http.Server{
		Addr:    "localhost:8000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
