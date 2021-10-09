package repository

import (
	"github.com/heroku/go-getting-started/model"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetMemories(db *sqlx.DB) ([]model.Memory, error) {
	var memories []model.Memory
	stmt := `select memories.id, memory, longitude, latitude, image, author_id, angle
			from memories join authors on memories.author_id = authors.id`
	err := db.Select(&memories, stmt)
	if err != nil {
		return nil, err
	}

	err = FillMemoriesEpisodes(db, memories)
	if err != nil {
		return nil, err
	}
	
	return memories, nil
}
// TODO:responseのseen消したい
func GetMyMemories(db *sqlx.DB, uid string) ([]model.Memory, error) {
	var memories []model.Memory
	stmt := `select memories.id, memory, longitude, latitude, image, author_id, angle
			from memories join authors on memories.author_id = authors.id
			where uuid = $1`
	err := db.Select(&memories, stmt, uid)
	if err != nil {
		return nil, err
	}

	err = FillMemoriesEpisodes(db, memories)
	if err != nil {
		return nil, err
	}
	
	return memories, nil
}

//TODO: sliceが配列のポインターを返してるから*[]momeryではく[]memory, 理解が浅いので見返したい
func FillMemoriesEpisodes(db *sqlx.DB, memories []model.Memory) (error) {
	for i, memory := range memories {
		var episodes []model.Episode
		err := db.Select(&episodes, `select id, episode, longitude, latitude from episodes where memory_id = $1`, memory.Id)
		if err != nil {
			return err
		}
		if len(episodes) == 0 {
			memories[i].Episodes = make([]model.Episode, 0)
		} else {
			memories[i].Episodes = episodes
		}
	}
	return nil
}

func CreateMemory(
	db *sqlx.DB,
	memory *model.Memory,
) error {
	tx := db.MustBegin()
	
	memory_stmt := `insert into memories(memory, image, longitude, latitude, angle, author_id)
					values ($1, $2, $3, $4, $5, $6)
					RETURNING id`
	var id int
	err := tx.QueryRow(
		memory_stmt,
		memory.Memory, memory.Image, memory.Longitude, memory.Latitude, memory.Angle, memory.AuthorId,
	).Scan(&id)
	if err != nil {
		tx.Rollback()
		return err
	}

	// TODO: bulk insertやprepared stmt使えば高速化できるはず
	episode_stmt := `insert into episodes(episode, longitude, latitude, memory_id) values ($1, $2, $3, $4)`
	for _, e := range memory.Episodes {
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
	var memoryIds []int64
	stmt := `select memory_id from author_seen_memory
			join authors on author_seen_memory.author_id = authors.id
			where authors.uuid = $1`
	err := db.Select(&memoryIds, stmt, uid)
	return memoryIds, err
}
