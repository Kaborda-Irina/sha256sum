package postrges

import (
	config "github.com/Kaborda-Irina/sha256sum/internal/configs"

	"github.com/jmoiron/sqlx"
	// postgres driver for Go's database/sqlx
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// Initialize postgres database
func Initialize(cfg *config.Config, logger *logrus.Logger) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.URI)
	if err != nil {
		logger.Error("error open db: ", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return db, nil
}
