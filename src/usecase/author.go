package usecase

import (
	"github.com/oonyaTeam/memory_share_server/repository"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type AuthorUseCase struct {
	db *sqlx.DB
}

func NewAuthorUseCase(db *sqlx.DB) *AuthorUseCase {
	return &AuthorUseCase{
		db: db,
	}
}

func (m *AuthorUseCase) RegisterAuthor(uid string) (error) {
	err := repository.RegisterAuthor(m.db, uid)
	return err
}


