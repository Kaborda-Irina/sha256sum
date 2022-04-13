package repositories

import (
	"context"
	"fmt"
	"github.com/Kaborda-Irina/sha256sum/internal/core/models"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
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

//SaveHashSum saves the data to the database and overwrites it if necessary
func (hr HashRepository) SaveHashSum(ctx context.Context, inputHashSum models.HashData) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := fmt.Sprintf(`SELECT checkHashSum($1, $2, $3, $4);`)
	hash := fmt.Sprintf("%x", inputHashSum.Hash)
	_, err := hr.db.Exec(query, inputHashSum.FileName, inputHashSum.FullFilePath, hash, inputHashSum.Algorithm)
	if err != nil {
		return err
	}

	return nil
}

//GetHashSum retrieves data from the database using the path and algorithm
func (hr HashRepository) GetHashSum(ctx context.Context, filePath string, algorithm string) (models.HashDataFromDB, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := fmt.Sprintf("SELECT id,fileName,fullFilePath,hashSum,algorithm FROM %s WHERE fullFilePath=$1 and algorithm=$2", nameTable)
	row := hr.db.QueryRow(query, filePath, algorithm)

	var hashDataFromDB models.HashDataFromDB
	err := row.Scan(&hashDataFromDB.Id, &hashDataFromDB.FileName, &hashDataFromDB.FullFilePath, &hashDataFromDB.Hash, &hashDataFromDB.Algorithm)
	if err != nil {
		return models.HashDataFromDB{}, err
	}

	return hashDataFromDB, nil
}

func (hr HashRepository) delete() {
	result, err := hr.db.Exec("delete from hashFiles")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(result.RowsAffected())
}
