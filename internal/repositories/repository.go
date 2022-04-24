package repositories

import (
	"github.com/Kaborda-Irina/sha256sum/internal/core/ports"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AppRepository struct {
	ports.IHashRepository
	logger *logrus.Logger
}

func NewAppRepository(db *sqlx.DB, logger *logrus.Logger) *AppRepository {
	return &AppRepository{
		IHashRepository: NewHashRepository(db, logger),
		logger:          logger,
	}
}
