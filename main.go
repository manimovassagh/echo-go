package main

import (
	"crud-echo/types"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello World")
	dsn := "host=localhost user=postgres password=postgres dbname=cruder port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error connecting to database:", err)
		return
	}
	fmt.Println(db)

	// Initialize the Handler with the database connection
	h := &types.Handler{DB: db}

	e := echo.New()

	// Register routes and pass the handler
	e.GET("/", h.HomeHandler())

	e.Logger.Fatal(e.Start(":4000"))
}
