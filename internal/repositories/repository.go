package repositories

import (
	"context"
	"fmt"
	"github.com/Kaborda-Irina/sha256sum/internal/core/models"
	"github.com/jmoiron/sqlx"
	"log"
)

const nameTable = "hashFiles"

type HashRepository struct {
	db *sqlx.DB
}

func NewHashRepository(db *sqlx.DB) *HashRepository {
	return &HashRepository{
		db,
	}
}

func (hr HashRepository) Ping(_ context.Context) error {
	log.Println("start repository was initialized")
	return hr.db.Ping()
}

func (hr HashRepository) SaveHashSum(inputHashSum models.HashSum, ctx context.Context) error {
	query := fmt.Sprintf("INSERT INTO %s (fileName,hashSum) VALUES ($1,$2)", nameTable)
	hash := fmt.Sprintf("%x\n", inputHashSum.Hash)
	_, err := hr.db.Exec(query, inputHashSum.FileName, hash)

	if err != nil {
		return err
	}
	log.Println("Hash sum successfully written to the database")
	return nil
}
