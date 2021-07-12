package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"

	"database/sql"
	_ "github.com/lib/pq"
)

func dbFunc(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        if _, err := db.Exec("CREATE TABLE IF NOT EXISTS ticks (tick timestamp)"); err != nil {
            c.String(http.StatusInternalServerError,
                fmt.Sprintf("Error creating database table: %q\n DB:%s\n", err, os.Getenv("DATABASE_URL")))
            return
        }

        if _, err := db.Exec("INSERT INTO ticks VALUES (now())"); err != nil {
            c.String(http.StatusInternalServerError,
                fmt.Sprintf("Error incrementing tick: %q", err))
            return
        }

        rows, err := db.Query("SELECT tick FROM ticks")
        if err != nil {
            c.String(http.StatusInternalServerError,
                fmt.Sprintf("Error reading ticks: %q", err))
            return
        }

        defer rows.Close()
        for rows.Next() {
            var tick time.Time
            if err := rows.Scan(&tick); err != nil {
                c.String(http.StatusInternalServerError,
                    fmt.Sprintf("Error scanning ticks: %q", err))
                return
            }
            c.String(http.StatusOK, fmt.Sprintf("Read from DB: %s\n", tick.String()))
        }
    }
}

func connectDB() (*sql.DB, error){
	// if e := os.Getenv("DEV"); e == "DEV" {
	// 	db, err :=  sql.Open("postgres", "user= dbname=test password= sslmode=disable host=localhost ")
	// 	return db, err; 
	// } else {
	// 	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// 	return db, err;
	// }
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	return db, err;
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg" : "messages",
		})
	})


	db, err := connectDB()
    if err != nil {
        log.Fatalf("Error opening database: %q", err)
    }

	router.GET("/db", dbFunc(db))
	router.Run(":" + port)
}
