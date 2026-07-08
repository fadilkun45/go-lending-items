// @title Loans Item API
// @version 1.0
// @description REST API untuk sistem peminjaman barang P2P
// @host localhost:8000
// @BasePath /
package main

import (
	"fmt"
	_ "loans-item-go/docs"

	"loans-item-go/config"
	"loans-item-go/controller/category"
	"loans-item-go/controller/item"
	"loans-item-go/controller/loan"
	"loans-item-go/controller/user"
	"loans-item-go/helper"
	"loans-item-go/repository/category"
	"loans-item-go/repository/item"
	"loans-item-go/repository/loan"
	"loans-item-go/repository/loan_item"
	"loans-item-go/repository/user"
	"loans-item-go/service/category"
	"loans-item-go/service/item"
	"loans-item-go/service/loan"
	"loans-item-go/service/user"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	db := config.OpenConnection()
	cfg, err := config.LoadConfig()
	helper.PanicIfError(err)

	// repository
	userRepo := userrepo.NewRepositoryImpl()
	categoryRepo := categoryrepo.NewRepositoryImpl()
	itemRepo := itemrepo.NewRepositoryImpl()
	loanRepo := loanrepo.NewRepositoryImpl()
	loanItemRepo := loanitemrepo.NewRepositoryImpl()

	// service
	userSvc := usersvc.NewServiceImpl(userRepo, db)
	categorySvc := categorysvc.NewServiceImpl(categoryRepo, db)
	itemSvc := itemsvc.NewServiceImpl(itemRepo, db)
	loanSvc := loansvc.NewServiceImpl(loanRepo, itemRepo, loanItemRepo, db)

	// controller
	userCtrl := userctrl.NewController(userSvc)
	categoryCtrl := categoryctrl.NewController(categorySvc)
	itemCtrl := itemctrl.NewController(itemSvc)
	loanCtrl := loanctrl.NewController(loanSvc)

	// router
	mux := http.NewServeMux()
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
	mux.HandleFunc("GET /api/items/owner/{ownerId}", itemCtrl.FindByOwner)
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
		Addr:    "localhost:" + cfg.APP_PORT,
		Handler: mux,
	}

	fmt.Println("Server started on port", cfg.APP_PORT)
	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
