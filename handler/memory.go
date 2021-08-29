package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Memory struct {
	Memory      string    `json:"memory"`
	Longitude   float64   `json:"longitude"`
	Latitude    float64   `json:"latitude"`
	Seen_author []string  `json:"seen_author"`
	Episodes    []Episode `json:"episodes"`
	Image       string    `json:"image"`
	Author      string    `json:"author"`
}

type Episode struct {
	Id        string  `json:"id"`
	Episode   string  `json:"episode"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func GetMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}

func GetMemories(c *gin.Context) {
	e1 := Episode{
		Id:        "first_id",
		Episode:   "subepisode 1Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		Longitude: 136.597816,
		Latitude:  34.860853,
	}
	e2 := Episode{
		Id:        "second_id",
		Episode:   "sub episode2 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		Longitude: 136.599202,
		Latitude:  34.860156,
	}
	m := Memory{
		"main episode1 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		136.601064,
		34.857498,
		[]string{"author1", "author2"},
		[]Episode{e1, e2},
		"https://pbs.twimg.com/media/E6CYtu1VcAIjMvY?format=jpg&name=large",
		"author1",
	}
	m2 := Memory{
		"main episode2 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		136.602276,
		34.856582,
		[]string{"author1", "author3"},
		[]Episode{e1, e2},
		"https://pbs.twimg.com/media/E6FYPWLVIAQvY04?format=jpg&name=small",
		"author2",
	}
	c.JSON(http.StatusOK, gin.H{
		"memories": []Memory{m, m2},
		// "msg": "Memories",
	})
}

func GetMyMemories(c *gin.Context) {
	// uuid := c.Query("uuid")

	e1 := Episode{
		Id:        "first_id",
		Episode:   "subepisode 1Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		Longitude: 136.597816,
		Latitude:  34.860853,
	}
	e2 := Episode{
		Id:        "second_id",
		Episode:   "sub episode2 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		Longitude: 136.599202,
		Latitude:  34.860156,
	}
	m := Memory{
		"main episode1 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		136.601064,
		34.857498,
		[]string{"author1", "author2"},
		[]Episode{e1, e2},
		"https://pbs.twimg.com/media/E6CYtu1VcAIjMvY?format=jpg&name=large",
		"author1",
	}
	m2 := Memory{
		"main episode2 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		136.602276,
		34.856582,
		[]string{"author1", "author3"},
		[]Episode{e1, e2},
		"https://pbs.twimg.com/media/E6FYPWLVIAQvY04?format=jpg&name=small",
		"author2",
	}
	c.JSON(http.StatusOK, gin.H{
		"memories": []Memory{m, m2},
	})
}

func CreateMemory(c *gin.Context) {
	var mb Memory
	if err := c.BindJSON(&mb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"msg": "OK",
	})
}
