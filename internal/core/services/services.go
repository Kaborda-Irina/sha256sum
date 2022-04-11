package services

import (
	"context"
	"github.com/Kaborda-Irina/sha256sum/internal/core/models"
	"github.com/Kaborda-Irina/sha256sum/internal/core/ports"
	"log"
)

type HashService struct {
	hashRepository ports.IHashRepository
}

func NewHashService(hashRepository ports.IHashRepository) ports.IHashService {
	return HashService{
		hashRepository,
	}
}

func (hs HashService) Ping(_ context.Context) error {
	log.Println("hash service was initialized")
	return nil
}

func (hs HashService) SaveHashSum(hashSum models.HashSum, ctx context.Context) error {
	err := hs.hashRepository.SaveHashSum(hashSum, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (hs HashService) GetHashSum(filePath string, ctx context.Context) (models.HashSumFromDB, error) {
	hash, err := hs.hashRepository.GetHashSum(filePath, ctx)
	if err != nil {
		log.Printf("hash service didn't get hash sum %s", err)
		return models.HashSumFromDB{}, err
	}
	return hash, nil
}
