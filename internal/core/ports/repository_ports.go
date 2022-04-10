package ports

import (
	"context"
	"github.com/Kaborda-Irina/sha256sum/internal/core/models"
)

type IHashRepository interface {
	Ping(_ context.Context) error
	SaveHashSum(hashSum models.HashSum, ctx context.Context) error
}
