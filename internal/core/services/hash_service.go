package services

import (
	"context"
	"fmt"
	"github.com/Kaborda-Irina/sha256sum/internal/core/models"
	"github.com/Kaborda-Irina/sha256sum/internal/core/ports"
	"github.com/Kaborda-Irina/sha256sum/pkg/api"
	"github.com/sirupsen/logrus"
	"time"
)

type HashService struct {
	hashRepository ports.IHashRepository
	logger         *logrus.Logger
}

//NewHashService creates a new struct HashService
func NewHashService(hashRepository ports.IHashRepository, logger *logrus.Logger) *HashService {
	return &HashService{
		hashRepository: hashRepository,
		logger:         logger,
	}
}

//SaveHashData accesses the repository to save data to the database
func (hs HashService) SaveHashData(ctx context.Context, allHashData []api.HashData) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := hs.hashRepository.SaveHashData(ctx, allHashData)
	if err != nil {
		hs.logger.Error(err)
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
		hs.logger.Error("hash service didn't get hash sum %s", err)
		return nil, err
	}

	return hash, nil
}

//ChangedHashes checks if the current data has changed with the data stored in the database
func (hs HashService) ChangedHashes(currentHashData []api.HashData, hashSumFromDB []models.HashDataFromDB) error {
	var deletedResult []models.DeletedHashes
	var trigger bool
	var count int

	for _, dataFromDB := range hashSumFromDB {
		trigger = false
		for _, dataCurrent := range currentHashData {
			if dataFromDB.FullFilePath == dataCurrent.FullFilePath {
				if dataFromDB.Hash != dataCurrent.Hash {
					count++
					fmt.Printf("Changed: file - %s the path %s, old hash sum %s, new hash sum %s\n",
						dataFromDB.FileName, dataFromDB.FullFilePath, dataFromDB.Hash, dataCurrent.Hash)
				}
				trigger = true
				break
			}
		}

		if !trigger {
			count++
			fmt.Printf("Deleted: file - %s the path %s hash sum %s\n",
				dataFromDB.FileName, dataFromDB.FullFilePath, dataFromDB.Hash)
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
			count++
			fmt.Printf("Added: file - %s the path %s hash sum %s\n",
				dataCurrent.FileName, dataCurrent.FullFilePath, dataCurrent.Hash)
		}
	}

	if len(deletedResult) > 0 {
		err := hs.hashRepository.UpdateDeletedItems(deletedResult)
		if err != nil {
			hs.logger.Error(err)
			return err
		}
	}

	if count == 0 {
		fmt.Println("Files didn't changed, added or deleted")
	}
	return nil
}
