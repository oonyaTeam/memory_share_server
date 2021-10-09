package usecase

import (
	"github.com/heroku/go-getting-started/model"
	"github.com/heroku/go-getting-started/repository"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type MemoryUseCase struct {
	db *sqlx.DB
}

func NewMemoryUseCase(db *sqlx.DB) *MemoryUseCase {
	return &MemoryUseCase{
		db: db,
	}
}

func (m *MemoryUseCase) GetMemories(uid string) ([]model.Memory, error) {
	memories, err := repository.GetMemories(m.db)
	if err != nil {
		return nil, err
	}

	seenMemoryList, err := repository.SeenMemoryIds(m.db, uid)
	if err != nil {
		return nil, err
	}
	
	for i, memory := range memories {
		if contains(seenMemoryList, memory.Id) {
			memories[i].Seen = true
		}
	}
	return memories, nil
}
func contains(s []int64, e int64) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func (m *MemoryUseCase) GetMyMemories(uid string) ([]model.Memory, error) {
	memories, err := repository.GetMyMemories(m.db, uid)
	if err != nil {
		return nil, err
	}

	return memories, nil
}

func (m *MemoryUseCase) CreateMemories(
	memory *model.Memory,
	uid string,
) (error) {
	authorId, err := repository.GetAuthorId(m.db, uid)
	if err != nil {
		return err
	}
	memory.AuthorId = authorId
	err = repository.CreateMemory(
		m.db, memory,
	)
	return err
}