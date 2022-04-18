package services

import (
	"context"
	"fmt"
	"github.com/Kaborda-Irina/sha256sum/internal/core/models"
	"github.com/Kaborda-Irina/sha256sum/internal/core/ports"
	"log"
	"time"
)

type HashService struct {
	hashRepository ports.IHashRepository
}

//NewHashService creates a new struct HashService
func NewHashService(hashRepository ports.IHashRepository) ports.IHashService {
	return HashService{
		hashRepository,
	}
}
func (hs HashService) SaveHashData(ctx context.Context, allHashData []models.HashData) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := hs.hashRepository.SaveHashData(ctx, allHashData)
	if err != nil {
		return err
	}
	return nil
}

//GetHashSum accesses the repository to get data from the database
func (hs HashService) GetHashSum(ctx context.Context, dirFiles string, algorithm string) ([]models.HashDataFromDB, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	hash, err := hs.hashRepository.GetHashSum(ctx, dirFiles, algorithm)
	if err != nil {
		log.Printf("hash service didn't get hash sum %s", err)
		return nil, err
	}

	return hash, nil
}

//ChangedHashes checks if the current data has changed with the data stored in the database
func (hs HashService) ChangedHashes(currentHashData []models.HashData, hashSumFromDB []models.HashDataFromDB) ([]models.ChangedHashes, []models.DeletedHashes, []models.AddedHashes, error) {
	var changeResult []models.ChangedHashes
	var deletedResult []models.DeletedHashes
	var addedResult []models.AddedHashes
	var trigger bool

	for _, dataFromDB := range hashSumFromDB {
		trigger = false
		for _, dataCurrent := range currentHashData {
			if dataFromDB.FullFilePath == dataCurrent.FullFilePath {
				if dataFromDB.Hash != fmt.Sprintf("%x", dataCurrent.Hash) {
					changeResult = append(changeResult, models.ChangedHashes{
						FileName:    dataFromDB.FileName,
						FilePath:    dataFromDB.FullFilePath,
						OldChecksum: dataFromDB.Hash,
						NewChecksum: fmt.Sprintf("%x", dataCurrent.Hash),
					})
				}
				trigger = true
				break
			}
		}

		if !trigger {
			deletedResult = append(deletedResult, models.DeletedHashes{
				FileName:    dataFromDB.FileName,
				FilePath:    dataFromDB.FullFilePath,
				OldChecksum: dataFromDB.Hash,
				Algorithm:   dataFromDB.Algorithm,
			})
		}
	}

	for _, dataCurrent := range currentHashData {
		trigger = false
		for _, dataFromDB := range hashSumFromDB {
			if dataCurrent.FullFilePath == dataFromDB.FullFilePath {
				trigger = true
				break
			}
		}

		if !trigger {
			addedResult = append(addedResult, models.AddedHashes{
				FileName:    dataCurrent.FileName,
				FilePath:    dataCurrent.FullFilePath,
				NewChecksum: fmt.Sprintf("%x", dataCurrent.Hash),
				Algorithm:   dataCurrent.Algorithm,
			})
		}
	}

	if len(deletedResult) > 0 {
		err := hs.hashRepository.UpdateDeletedItems(deletedResult)
		if err != nil {
			return []models.ChangedHashes{}, []models.DeletedHashes{}, []models.AddedHashes{}, err
		}
	}
	return changeResult, deletedResult, addedResult, nil
}
