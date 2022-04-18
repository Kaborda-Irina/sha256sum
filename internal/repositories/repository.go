package repositories

import (
	"context"
	"fmt"
	"github.com/Kaborda-Irina/sha256sum/internal/core/models"
	"github.com/jmoiron/sqlx"
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

//SaveHashData iterates through all elements of the slice and triggers the save to database function
func (hr HashRepository) SaveHashData(ctx context.Context, allHashData []models.HashData) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	start := time.Now()
	tx, err := hr.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf(`
		INSERT INTO hashFiles (fileName,fullFilePath,hashSum,algorithm) 
		VALUES($1,$2,$3,$4) ON CONFLICT (fullFilePath,algorithm) DO UPDATE SET hashSum=EXCLUDED.hashSum`)

	for _, hash := range allHashData {
		_, err = tx.Exec(query, hash.FileName, hash.FullFilePath, fmt.Sprintf("%x", hash.Hash), hash.Algorithm)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	fmt.Println(time.Since(start).Seconds())
	return tx.Commit()

}

//GetHashSum retrieves data from the database using the path and algorithm
func (hr HashRepository) GetHashSum(ctx context.Context, dirFiles string, algorithm string) ([]models.HashDataFromDB, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var allHashDataFromDB []models.HashDataFromDB

	query := fmt.Sprintf("SELECT id,fileName,fullFilePath,hashSum,algorithm FROM %s WHERE fullFilePath LIKE $1 and algorithm=$2", nameTable)

	rows, err := hr.db.Query(query, "%"+dirFiles+"%", algorithm)
	if err != nil {
		return []models.HashDataFromDB{}, err
	}
	for rows.Next() {
		var hashDataFromDB models.HashDataFromDB
		err := rows.Scan(&hashDataFromDB.Id, &hashDataFromDB.FileName, &hashDataFromDB.FullFilePath, &hashDataFromDB.Hash, &hashDataFromDB.Algorithm)
		if err != nil {
			return []models.HashDataFromDB{}, err
		}
		allHashDataFromDB = append(allHashDataFromDB, hashDataFromDB)
	}

	return allHashDataFromDB, nil
}

//UpdateDeletedItems changes the deleted field to true in the database for each row if the file name has been deleted
func (hr HashRepository) UpdateDeletedItems(deletedItems []models.DeletedHashes) error {
	tx, err := hr.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`UPDATE %s SET deleted = true WHERE fullFilePath=$1 AND algorithm=$2`, nameTable)

	for _, item := range deletedItems {
		_, err := tx.Exec(query, item.FilePath, item.Algorithm)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
