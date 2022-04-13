package ports

import (
	"context"
	"github.com/Kaborda-Irina/sha256sum/internal/core/models"
)

type IHashRepository interface {
	SaveHashSum(ctx context.Context, hashSum models.HashData) error
	GetHashSum(ctx context.Context, filePath string, algorithm string) (models.HashDataFromDB, error)
}
