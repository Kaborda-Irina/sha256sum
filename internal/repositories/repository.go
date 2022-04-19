package repositories

import (
	"github.com/Kaborda-Irina/sha256sum/internal/core/ports"
	"github.com/jmoiron/sqlx"
)

type AppRepository struct {
	ports.IHashRepository
}

func NewAppRepository(db *sqlx.DB) *AppRepository {
	return &AppRepository{
		IHashRepository: NewHashRepository(db),
	}
}
