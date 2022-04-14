package ports

import (
	"context"
	"github.com/Kaborda-Irina/sha256sum/internal/core/models"
)

type IHashService interface {
	SaveHashDir(ctx context.Context, allHashData []models.HashData) error
	SaveHashData(ctx context.Context, hashData models.HashData) error
	GetHashSum(ctx context.Context, allHashData []models.HashData) ([]models.HashDataFromDB, error)
}
