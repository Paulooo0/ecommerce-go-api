package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"

	"github.com/Paulooo0/ecommerce/goapi/internal/database"
	"github.com/Paulooo0/ecommerce/goapi/internal/service"
	"github.com/Paulooo0/ecommerce/goapi/internal/webserver"
)

func main() {
	db, err := sql.Open("mysql", "root:O0etybU6g91@tcp(localhost:3306)/ecommerce")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)

	webCategoryHandler := webserver.NewWebCategotyHandler(categoryService)
	webProductHandler := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Get("/category/{id}", webCategoryHandler.GetCategory)
	c.Get("/category", webCategoryHandler.GetCategories)
	c.Post("/category", webCategoryHandler.CreateCategory)

	c.Get("/product/{id}", webProductHandler.GetProduct)
	c.Get("/product/category/{categoryID}", webProductHandler.GetPrductByCategoryID)
	c.Get("/product", webProductHandler.GetProducts)
	c.Post("/product", webProductHandler.CreateProduct)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", c)
}
