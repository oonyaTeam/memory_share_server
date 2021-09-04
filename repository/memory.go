package repository

import (
	"github.com/heroku/go-getting-started/model"

	"database/sql"
	_ "github.com/lib/pq"

	"log"
)

func GetMemories(db *sql.DB) ([]model.Memory, error) {
	log.Print("This is repository\n\n\n")

	e1 := model.Episode{
		Id:        "first_id",
		Episode:   "subepisode 1Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		Longitude: 136.597816,
		Latitude:  34.860853,
	}
	e2 := model.Episode{
		Id:        "second_id",
		Episode:   "sub episode2 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		Longitude: 136.599202,
		Latitude:  34.860156,
	}
	m := model.Memory{
		Memory:      "main episode1 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		Longitude:   136.601064,
		Latitude:    34.857498,
		Seen_author: []string{"author1", "author2"},
		Episodes:    []model.Episode{e1, e2},
		Image:       "https://pbs.twimg.com/media/E6CYtu1VcAIjMvY?format=jpg&name=large",
		Author:      "author1",
	}
	m2 := model.Memory{
		Memory:      "main episode2 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		Longitude:   136.602276,
		Latitude:    34.856582,
		Seen_author: []string{"author1", "author3"},
		Episodes:    []model.Episode{e1, e2},
		Image:       "https://pbs.twimg.com/media/E6FYPWLVIAQvY04?format=jpg&name=small",
		Author:      "author2",
	}
	return []model.Memory{m, m2}, nil
}

func GetMyMemories(db *sql.DB, uuid string) ([]model.Memory, error) {
	e1 := model.Episode{
		Id:        "first_id",
		Episode:   "subepisode 1Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		Longitude: 136.597816,
		Latitude:  34.860853,
	}
	e2 := model.Episode{
		Id:        "second_id",
		Episode:   "sub episode2 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		Longitude: 136.599202,
		Latitude:  34.860156,
	}
	m := model.Memory{
		Memory:      "main episode1 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		Longitude:   136.601064,
		Latitude:    34.857498,
		Seen_author: []string{"author1", "author2"},
		Episodes:    []model.Episode{e1, e2},
		Image:       "https://pbs.twimg.com/media/E6CYtu1VcAIjMvY?format=jpg&name=large",
		Author:      "author1",
	}
	m2 := model.Memory{
		Memory:      "main episode2 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		Longitude:   136.602276,
		Latitude:    34.856582,
		Seen_author: []string{"author1", "author3"},
		Episodes:    []model.Episode{e1, e2},
		Image:       "https://pbs.twimg.com/media/E6FYPWLVIAQvY04?format=jpg&name=small",
		Author:      "author2",
	}

	return []model.Memory{m, m2}, nil
}

func CreateMemory(db *sql.DB, m model.Memory) error {
	// insert
	return nil
}