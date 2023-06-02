package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	//connStr := "postgres://postgres:postgres@localhost:5432/dbdata"
	connStr1 := "postgres://dbdata:dbdpsw@localhost:5432/dbdata"
	db, err := sql.Open("postgres", connStr1)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	user, err := db.Query("SELECT id,first_name, last_name FROM person")

	for user.Next() {
		var (
			id         int
			first_name string
			last_name  string
		)

		err := user.Scan(&id, &first_name, &last_name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d %s %s \n", id, first_name, last_name)
	}

}
