package repository

import (
	"github.com/heroku/go-getting-started/model"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"log"
)

func GetMemories(db *sqlx.DB) ([]model.Memory, error) {
	log.Print("This is repository\n\n\n")

	var memories []model.Memory
	stmt := `select memories.id, memory, longitude, latitude, image, author_id, angle
			from memories join authors on memories.author_id = authors.id`
	err := db.Select(&memories, stmt)
	if err != nil {
		return nil, err
	}
	log.Println(memories)

	for i, memory := range memories {
		var episodes []model.Episode
		err = db.Select(&episodes, `select id, episode, longitude, latitude from episodes where memory_id = $1`, memory.Id)
		if err != nil {
			return nil, err
		}
		if len(episodes) == 0 {
			memories[i].Episodes = make([]model.Episode, 0)
		} else {
			memories[i].Episodes = episodes
		}
	}
	log.Println(memories)
	
	return memories, nil
}

func GetMyMemories(db *sqlx.DB, uuid string) ([]model.Memory, error) {
	e1 := model.Episode{
		Id:        1,
		Episode:   "subepisode 1Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		Longitude: 136.597816,
		Latitude:  34.860853,
	}
	e2 := model.Episode{
		Id:        2,
		Episode:   "sub episode2 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		Longitude: 136.599202,
		Latitude:  34.860156,
	}
	m := model.Memory{
		Memory:    "main episode1 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		Longitude: 136.601064,
		Latitude:  34.857498,
		Episodes:  []model.Episode{e1, e2},
		Image:     "https://pbs.twimg.com/media/E6CYtu1VcAIjMvY?format=jpg&name=large",
		AuthorId:  1,
		Angle:  38.54,
		Seen:  true,
	}
	m2 := model.Memory{
		Memory:      "main episode2 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		Longitude:   136.602276,
		Latitude:    34.856582,
		Episodes:    []model.Episode{e1, e2},
		Image:       "https://pbs.twimg.com/media/E6FYPWLVIAQvY04?format=jpg&name=small",
		AuthorId:    2,
		Angle:    67.12,
		Seen:    false,
	}

	return []model.Memory{m, m2}, nil
}

func CreateMemory(db *sqlx.DB, m model.Memory) error {
	tx := db.MustBegin()
	
	memory_stmt := `insert into memories(memory, image, longitude, latitude, author_id, angle) values ($1, $2, $3, $4, $5, $6)
	RETURNING id`
	var id int
	err := tx.QueryRow(memory_stmt, m.Memory, m.Image, m.Longitude, m.Latitude, m.AuthorId, m.Angle).Scan(&id)
	if err != nil {
		tx.Rollback()
		return err
	}

	episode_stmt := `insert into episodes(episode, longitude, latitude, memory_id) values ($1, $2, $3, $4)`
	for _, e := range m.Episodes {
		_, err := tx.Exec(episode_stmt, e.Episode, e.Longitude, e.Latitude, id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	return err
}

func SeenMemoryIds(db *sqlx.DB, uid string) ([]int64, error) {
	// var memoryIds []int64
	// stmt := `select memory_id from author_seen_memory where author_id=$1 `
	// err := db.Select(&memoryIds, stmt, uid)
	// return memoryIds, err
	return []int64{1, 2, 3}, nil
}
