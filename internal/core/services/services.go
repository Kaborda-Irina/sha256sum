package services

import (
	"context"
	"github.com/Kaborda-Irina/sha256sum/internal/core/models"
	"github.com/Kaborda-Irina/sha256sum/internal/core/ports"
	"log"
	"time"
)

type HashService struct {
	hashRepository ports.IHashRepository
}

func NewHashService(hashRepository ports.IHashRepository) ports.IHashService {
	return HashService{
		hashRepository,
	}
}
func (hs HashService) SaveHashDir(ctx context.Context, allHashData []models.HashData) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := hs.hashRepository.SaveHashDir(ctx, allHashData)
	if err != nil {
		return err
	}
	return nil
}

func (hs HashService) SaveHashData(ctx context.Context, hashData models.HashData) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := hs.hashRepository.SaveHashData(ctx, hashData)
	if err != nil {
		return err
	}
	return nil
}

func (hs HashService) GetHashSum(ctx context.Context, allHashData []models.HashData) ([]models.HashDataFromDB, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	var allHashDataFromDB []models.HashDataFromDB
	for _, hashData := range allHashData {
		hash, err := hs.hashRepository.GetHashSum(ctx, hashData.FullFilePath, hashData.Algorithm)
		if err != nil {
			log.Printf("hash service didn't get hash sum %s", err)
			return nil, err
		}
		allHashDataFromDB = append(allHashDataFromDB, hash)
	}
	return allHashDataFromDB, nil
}
