package postrges

import (
	"fmt"
	config "github.com/Kaborda-Irina/sha256sum/internal/configs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//Initialize postgres database
func Initialize(cfg config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.URI)
	if err != nil {
		fmt.Printf("error open db: %v\n", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
