package main

import (
	"database/sql"
	"net/http"

	"github.com/firstthumb/todo/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db := initDB("storage.db")
	migrate(db)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/tasks", handlers.GetTasks(db))
	r.POST("/tasks", handlers.PutTask(db))
	r.DELETE("/tasks/:id", handlers.DeleteTask(db))

	r.Run("localhost:8000")
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS tasks(
		id 				INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name 			VARCHAR NOT NULL,
		completed INTEGER NOT NULL DEFAULT 0,
		created 	DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}
