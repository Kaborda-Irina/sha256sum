package ports

import (
	"context"
	"github.com/Kaborda-Irina/sha256sum/internal/core/models"
)

type IAppRepository interface{}

type IHashRepository interface {
	SaveHashData(ctx context.Context, allHashData []models.HashData) error
	GetHashSum(ctx context.Context, dirFiles string, algorithm string) ([]models.HashDataFromDB, error)
	UpdateDeletedItems(deletedItems []models.DeletedHashes) error
}
