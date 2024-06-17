package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	// _をつけているのは、パッケージをインポートしているが、そのパッケージを使っていないため
	// 内部的に使用しているため、エラーが出ないようにするため
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

// jsonタグをつけることで、json形式での出力時に指定した名前で出力される
// jsonタグをつけないと、フィールド名で出力される
// つまり大文字で始まるフィールド名は、jsonタグをつけないと大文字で出力される
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	db := initDB("example.db")
	e := echo.New()
	e.Use(middleware.Logger())

	e.POST("/users", func(c echo.Context) error {
		name := c.FormValue("name")
		age, _ := strconv.Atoi(c.FormValue("age"))

		result, err := db.Exec(`
			INSERT INTO users (name, age) VALUES (?, ?)`, name, age)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		id, _ := result.LastInsertId()
		// &Userとすることで、Userのポインタを返す
		return c.JSON(http.StatusOK, &User{ID: int(id), Name: name, Age: age})
	})

	e.GET("/users", func(c echo.Context) error {
		rows, err := db.Query("SELECT id, name, age FROM users")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		defer rows.Close()

		users := []User{}
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			users = append(users, user)
		}
		return c.JSON(http.StatusOK, users)
	})

	e.Start(":8080")
	// db, err := sql.Open("sqlite3", "./example.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// createTableSQL := `CREATE TABLE IF NOT EXISTS users (
	// 			id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 			name TEXT NOT NULL,
	// 			age INTEGER NOT NULL);
	// 	`

	// _, err = db.Exec(createTableSQL)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// println("Table created successfully")
}
