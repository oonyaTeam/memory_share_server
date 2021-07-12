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

type Memory struct{
	Memory string `json:"memory"`
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
	Seen_author []string `json:"seen_author"`
	Episodes []Episode `json:"episodes"`
	Image string `json:"image"`
	Author string `json:"author"`
}

type Episode struct{
	Id string `json:"id"`
	Episode string `json:"episode"`
	Distance int `json:"distance"`
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
	e1 := Episode{
		Id :"first_id",
		Episode: "subepisode 1Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		Distance : 30,
	}
	e2 := Episode{
		"second_id",
		"sub episode2 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		50,
	}
	m := Memory{
		"main episode1 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		30.5,
		40.5,
		[]string{"author1", "author2"},
		[]Episode{e1, e2},
		"https://pbs.twimg.com/media/E6CYtu1VcAIjMvY?format=jpg&name=large",
		"author1",
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg" : "OK",
		})
	})
	
	router.GET("/memories", func(c *gin.Context) {
		c.JSON(http.StatusOK, m)
	})


	db, err := connectDB()
    if err != nil {
        log.Fatalf("Error opening database: %q", err)
    }

	router.GET("/db", dbFunc(db))
	router.Run(":" + port)
}
