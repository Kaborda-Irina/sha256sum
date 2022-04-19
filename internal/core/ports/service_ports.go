package ports

import (
	"context"
	"github.com/Kaborda-Irina/sha256sum/internal/core/models"
	"os"
)

type IAppService interface {
	StartGetHashData(ctx context.Context, flagName string, algorithm string, jobs chan string, results chan models.HashData, sig chan os.Signal)
	StartCheckHashData(ctx context.Context, flagName string, algorithm string, jobs chan string, results chan models.HashData, sig chan os.Signal)
}

type IHashService interface {
	SaveHashData(ctx context.Context, allHashData []models.HashData) error
	GetHashSum(ctx context.Context, dirFiles string, algorithm string) ([]models.HashDataFromDB, error)
	ChangedHashes(currentHashData []models.HashData, hashSumFromDB []models.HashDataFromDB) error
}
