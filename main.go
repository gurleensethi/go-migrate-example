package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Post struct {
	Title string
	Body  string
}

func main() {
	// connect to database
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5454/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// connect to check connection
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping db", err)
	}

	// insert a new post
	_, err = db.Exec("INSERT INTO posts VALUES($1, $2)", "Title", "Body")
	if err != nil {
		log.Fatal(err)
	}

	// query all the posts
	result, err := db.Query("SELECT * FROM posts;")
	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()

	// iterate over all the posts and print them
	for result.Next() {
		var p Post

		err := result.Scan(&p.Title, &p.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", p)
	}
}
