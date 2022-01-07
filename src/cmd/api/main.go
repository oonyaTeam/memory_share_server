package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oonyaTeam/memory_share_server/firebase"
	"github.com/oonyaTeam/memory_share_server/handler"
	"github.com/oonyaTeam/memory_share_server/middleware"
	"github.com/oonyaTeam/memory_share_server/usecase"
	_ "github.com/heroku/x/hmetrics/onload"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/gin-contrib/cors"
)

func connectDB() (*sqlx.DB, error) {
	if e := os.Getenv("DEV"); e == "DEV" {
		db, err := sqlx.Open("postgres", os.Getenv("POSTGRESQL_URL"))
		return db, err;
	} else {
		db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
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
	memoryUseCase := usecase.NewMemoryUseCase(db)
	memoryHandler := handler.NewMemoryHandler(memoryUseCase)
	authorUseCase := usecase.NewAuthorUseCase(db)
	authorHandler := handler.NewAuthorHandler(authorUseCase)

	
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
		authRouter.GET("/memories/me", memoryHandler.GetMyMemories)
		authRouter.POST("/memories", memoryHandler.CreateMemory)
		authRouter.DELETE("/memories", memoryHandler.DeleteMemory)
		// footprintsとかの方がスッキリするが分かりづらいので妥協
		authRouter.POST("/memories/seen", memoryHandler.SeenMemory)

		authRouter.POST("/author", authorHandler.RegisterAuthor)
	}

	router.Run(":" + port)
}
