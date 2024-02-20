package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"runtime"
	"time"
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

func f() {
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
}
func main() {
	idx := 0
	for {
		fmt.Printf("round %d\n", idx)
		f()
		runtime.GC()
		time.Sleep(time.Second * 5)
		idx++
		if idx == 5 {
			time.Sleep(time.Hour * 111)
		}
	}
}
