package ports

import (
	"context"
	"github.com/Kaborda-Irina/sha256sum/internal/core/models"
)

type IHashService interface {
	SaveHashData(ctx context.Context, allHashData []models.HashData) error
	GetHashSum(ctx context.Context, dirFiles string, algorithm string) ([]models.HashDataFromDB, error)
	ChangedHashes(currentHashData []models.HashData, hashSumFromDB []models.HashDataFromDB) ([]models.ChangedHashes, []models.DeletedHashes, []models.AddedHashes, error)
}
