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

func (hs HashService) SaveHashSum(ctx context.Context, hashSum models.HashData) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := hs.hashRepository.SaveHashSum(ctx, hashSum)
	if err != nil {
		return err
	}
	return nil
}

func (hs HashService) GetHashSum(ctx context.Context, filePath string, algorithm string) (models.HashDataFromDB, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	hash, err := hs.hashRepository.GetHashSum(ctx, filePath, algorithm)
	if err != nil {
		log.Printf("hash service didn't get hash sum %s", err)
		return models.HashDataFromDB{}, err
	}
	return hash, nil
}
