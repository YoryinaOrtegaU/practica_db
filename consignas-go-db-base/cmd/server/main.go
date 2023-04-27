package main

import (
	"github.com/bootcamp-go/consignas-go-db.git/cmd/server/handler"
	"github.com/bootcamp-go/consignas-go-db.git/internal/product"
	"github.com/bootcamp-go/consignas-go-db.git/internal/product2"
	"github.com/gin-gonic/gin"

	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func main() {

	//storage := store.NewJsonStore("./products.json")
	//-------------------------------------
	databaseConfig := mysql.Config{
		User:   "user1",
		Passwd: "secret_password",
		Addr:   "127.0.0.1:3306",
		DBName: "my_db",
	}

	database, err := sql.Open("mysql", databaseConfig.FormatDSN())
	if err != nil {
		panic(err)
	}

	defer database.Close()

	if err = database.Ping(); err != nil {
		panic(err)
	}

	/* repository := &product2.RepositoryImpl{
		Database: database,
	}

	repository.GetByID(1) */

	//-------------------------------------
	//repo := product.NewRepository(storage)
	repo := product2.NewRepository2(database)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET(":id", productHandler.GetByID())
		products.POST("", productHandler.Post())
		products.DELETE(":id", productHandler.Delete())
		products.PATCH(":id", productHandler.Patch())
		products.PUT(":id", productHandler.Put())
	}

	r.Run(":8080")
}
