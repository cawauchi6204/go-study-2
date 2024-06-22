package main

import (
	"database/sql"
	"fmt"
	"log"

	"go-study/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	return db, err
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	taskController := controller.TaskController{}

	e.GET("/tasks", taskController.Get)

	e.POST("/tasks", taskController.Create)

	e.Start(":8080")
}
