package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var insert string

func init() {
	insert = "INSERT INTO dummy_table (f1,f2,f3,f4,f5,f6,f7,f8,f9,f10) VALUES "
	for i := 0; i < 9999; i += 1 {
		insert += fmt.Sprintf(`("%s","%s","%s","%s","%s","%s","%s","%s","%s","%s")`, "1111111111", "1111111111", "1111111111", "1111111111", "1111111111", "1111111111", "1111111111", "1111111111", "1111111111", "1111111111")
		insert += ","
	}
	insert = insert[:len(insert)-1]
}

func main() {
	r := gin.Default()
	r.GET("/run", func(c *gin.Context) {
		db, err := sql.Open("sqlite3", "file::memory:")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		if _, err := db.Exec(`CREATE TABLE dummy_table (id INT PRIMARY KEY,f1 VARCHAR(50),f2 VARCHAR(50),f3 VARCHAR(50),f4 VARCHAR(50),f5 VARCHAR(50),f6 VARCHAR(50),f7 VARCHAR(50),f8 VARCHAR(50),f9 VARCHAR(50),f10 VARCHAR(50));`); err != nil {
			panic(err)
		}
		if _, err := db.Exec(insert); err != nil {
			panic(err)
		}
		c.String(200, "OK")
	})
	r.Run(":8080")
}
