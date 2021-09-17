package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/handler"
	"github.com/heroku/go-getting-started/firebase"
	"github.com/heroku/go-getting-started/middleware"
	_ "github.com/heroku/x/hmetrics/onload"

	"database/sql"
	_ "github.com/lib/pq"

	"github.com/gin-contrib/cors"
)

func dbFunc(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS ticks (tick timestamp)"); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error creating database table: %q\n DB:%s\n", err, os.Getenv("DATABASE_URL")))
			return
		}

		// if _, err := db.Exec("INSERT INTO ticks VALUES (now())"); err != nil {
		// 	c.String(http.StatusInternalServerError,
		// 		fmt.Sprintf("Error incrementing tick: %q", err))
		// 	return
		// }

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

func connectDB() (*sql.DB, error) {
	if e := os.Getenv("DEV"); e == "DEV" {
		db, err :=  sql.Open("postgres", "user=postgres dbname=test password=Arrow0816 sslmode=disable host=localhost ")
		return db, err;
	} else {
		db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
		return db, err;
	}
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	// router.Use(gin.Logger())
	router.Use(cors.New(cors.Config{
        AllowMethods: []string{
            "POST",
            "GET",
            "OPTIONS",
            "PUT",
            "DELETE",
        },
        AllowHeaders: []string{
            "Access-Control-Allow-Headers",
            "Content-Type",
            "Content-Length",
            "Accept-Encoding",
            "Authorization",
			"Origin",
        },
        AllowOrigins: []string{
            "*",
        },
        MaxAge: 24 * time.Hour,
    }))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "OK",
		})
	})

	db, err := connectDB()
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	memoryHandler := handler.NewMemoryHandler(db)
	authorHandler := handler.NewAuthorHandler(db)

	
	firebase.CreateFirebaseJson()
	auth, err :=  firebase.InitializeAppWithRefreshToken()
	if err != nil {
		log.Println(err)
		log.Fatal("firebase死んでるけど大丈夫そ？")
	}
	authMiddleware := middleware.NewAuth(auth)
	
	// authをするGroup
	authRouter := router.Group("/", authMiddleware.AuthRequired)
	{
		authRouter.GET("/memories", memoryHandler.GetMemories)
		authRouter.GET("/mymemories", memoryHandler.GetMyMemories)
		authRouter.POST("/create-memory", memoryHandler.CreateMemory)

		authRouter.POST("/author", authorHandler.RegisterAuthor)
		authRouter.POST("/seen-memory", authorHandler.SeenMemory)
		
		authRouter.GET("/get1", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "get1",
			})
		})
	}

	router.GET("/db", dbFunc(db))
	router.Run(":" + port)
}
