package repository

import (
	"github.com/heroku/go-getting-started/model"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetMemories(db *sqlx.DB) ([]model.Memory, error) {
	var memories []model.Memory
	stmt := `select memories.id, memory, longitude, latitude, image, author_id, angle, created_at
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
	stmt := `select memories.id, memory, longitude, latitude, image, author_id, angle, created_at
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

func CreateMemory(db *sqlx.DB, memory *model.Memory) (error) {
	tx := db.MustBegin()
	
	memory_stmt := `insert into memories(memory, image, longitude, latitude, angle, author_id)
					values ($1, $2, $3, $4, $5, $6)
					returning id`
	var memoryId int64
	err := tx.QueryRow(
		memory_stmt,
		memory.Memory, memory.Image, memory.Longitude, memory.Latitude, memory.Angle, memory.AuthorId,
	).Scan(&memoryId)
	if err != nil {
		tx.Rollback()
		return err
	}
	memory.Id = memoryId
	
	// TODO: bulk insertやprepared stmt使えば高速化できるはず
	episode_stmt := `insert into episodes(episode, longitude, latitude, memory_id)
					values ($1, $2, $3, $4)
					returning id`
	var episodeId int64
	for i, e := range memory.Episodes {
		err := tx.QueryRow(episode_stmt, e.Episode, e.Longitude, e.Latitude, memoryId).Scan(&episodeId)
		if err != nil {
			tx.Rollback()
			return err
		}
		memory.Episodes[i].Id = episodeId
	}

	seen_stmt := `insert into author_seen_memory(memory_id, author_id) values ($1, $2)`
	_, err = tx.Exec(seen_stmt, memoryId, memory.AuthorId)
	if err != nil {
		tx.Rollback()
		return err
	}
	memory.Seen = true

	if err = tx.Commit(); err != nil {
		return err
	} else {
		return nil
	}	
}

func SeenMemoryIds(db *sqlx.DB, uid string) ([]int64, error) {
	var memoryIds []int64
	stmt := `select memory_id from author_seen_memory
			join authors on author_seen_memory.author_id = authors.id
			where authors.uuid = $1`
	err := db.Select(&memoryIds, stmt, uid)
	return memoryIds, err
}

func DeleteMemory(db *sqlx.DB, memoryId int) (error) {
	return nil
}

func SeenMemory(db *sqlx.DB, uuid string, memoryId int64) error {
	log.Println("seen memory")
	stmt := `insert into author_seen_memory(memory_id, author_id)
			select $1, id from authors where uuid = $2`
	_, err := db.Exec(stmt, memoryId, uuid)
	return err
}
