package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskController struct{}

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func (t *TaskController) Get(c echo.Context) error {
	// tasks, err := usecase.GetTasks()
	return c.JSON(http.StatusOK, nil)
}

func (t *TaskController) Create(c echo.Context) error {
	var task Task
	// バインドできなかった場合は400を返す
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	println(task.Title)

	return c.JSON(http.StatusOK, nil)

	// created, err := usecase.CreateTask(taske)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, nil)
	// }
}
