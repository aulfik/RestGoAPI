package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	host := "localhost"
	port := 5432
	user := "aulia"
	password := "abc123"
	dbname := "mydb"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		var (
			username string
			email    string
		)
		rows, err := db.Query("select * from user_data")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&username, &email)
			if err != nil {
				log.Fatal(err)
			}
			c.JSON(200, gin.H{
				"username": username,
				"email":    email,
			})
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
