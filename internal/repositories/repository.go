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
	log.Println("hash repository was initialized")
	return hr.db.Ping()
}

func (hr HashRepository) SaveHashSum(inputHashSum models.HashSum, ctx context.Context) error {
	query := fmt.Sprintf("INSERT INTO %s (fileName,fullFilePath,hashSum) VALUES ($1,$2,$3)", nameTable)
	hash := fmt.Sprintf("%x", inputHashSum.Hash)
	_, err := hr.db.Exec(query, inputHashSum.FileName, inputHashSum.FullFilePath, hash)

	if err != nil {
		return err
	}
	log.Println("Hash sum successfully written to the database")
	return nil
}

func (hr HashRepository) GetHashSum(filePath string, ctx context.Context) (models.HashSumFromDB, error) {
	query := fmt.Sprintf("SELECT id,fileName,fullFilePath,hashSum FROM %s WHERE fullFilePath=$1", nameTable)
	row := hr.db.QueryRow(query, filePath)

	var newHashSum models.HashSumFromDB
	err := row.Scan(&newHashSum.Id, &newHashSum.FileName, &newHashSum.FullFilePath, &newHashSum.Hash)
	if err != nil {
		return models.HashSumFromDB{}, err
	}

	return newHashSum, nil

}

func (hr HashRepository) delete() {
	result, err := hr.db.Exec("delete from hashFiles")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())
}
