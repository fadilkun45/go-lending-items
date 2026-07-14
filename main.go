// @title Loans Item API
// @version 1.0
// @description REST API untuk sistem peminjaman barang P2P
// @host localhost:8000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and the JWT token.
package main

import (
	"fmt"
	_ "loans-item-go/docs"

	"loans-item-go/config"
	categoryctrl "loans-item-go/controller/category"
	itemctrl "loans-item-go/controller/item"
	loanctrl "loans-item-go/controller/loan"
	userctrl "loans-item-go/controller/user"
	"loans-item-go/helper"
	"loans-item-go/middleware"
	categoryrepo "loans-item-go/repository/category"
	itemrepo "loans-item-go/repository/item"
	loanrepo "loans-item-go/repository/loan"
	loanitemrepo "loans-item-go/repository/loan_item"
	userrepo "loans-item-go/repository/user"
	categorysvc "loans-item-go/service/category"
	itemsvc "loans-item-go/service/item"
	loansvc "loans-item-go/service/loan"
	usersvc "loans-item-go/service/user"
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
	userSvc := usersvc.NewServiceImpl(userRepo, db, cfg.JWT_SECRET)
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
	mux.HandleFunc("GET /api/users/me", userCtrl.Me)
	mux.HandleFunc("GET /api/users/{id}", userCtrl.FindById)
	mux.HandleFunc("PUT /api/users/{id}", userCtrl.Update)
	mux.HandleFunc("DELETE /api/users/{id}", userCtrl.Delete)

	// category routes
	mux.HandleFunc("POST /api/categories", categoryCtrl.Create)
	mux.HandleFunc("GET /api/categories", categoryCtrl.FindAll)
	mux.HandleFunc("GET /api/categories/search", categoryCtrl.Search)
	mux.HandleFunc("GET /api/categories/{id}", categoryCtrl.FindById)
	mux.HandleFunc("PUT /api/categories/{id}", categoryCtrl.Update)
	mux.HandleFunc("DELETE /api/categories/{id}", categoryCtrl.Delete)

	// item routes
	mux.HandleFunc("POST /api/items", itemCtrl.Create)
	mux.HandleFunc("GET /api/items", itemCtrl.FindAll)
	mux.HandleFunc("GET /api/items/search", itemCtrl.Search)
	mux.HandleFunc("GET /api/items/owner", itemCtrl.FindByOwner)
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

	handler := middleware.CORS(middleware.Auth(cfg.JWT_SECRET)(mux))

	server := &http.Server{
		Addr:    "localhost:" + cfg.APP_PORT,
		Handler: handler,
	}

	fmt.Println("Server started on port", cfg.APP_PORT)
	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
