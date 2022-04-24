package ports

import (
	"context"
	"os"
	"sync"

	"github.com/Kaborda-Irina/sha256sum/internal/core/models"
	"github.com/Kaborda-Irina/sha256sum/pkg/api"

	"github.com/sirupsen/logrus"
)

type IAppService interface {
	StartGetHashData(ctx context.Context, flagName string, algorithm string, jobs chan string, results chan api.HashData, sig chan os.Signal)
	StartCheckHashData(ctx context.Context, flagName string, algorithm string, jobs chan string, results chan api.HashData, sig chan os.Signal)
}

type IHashService interface {
	SaveHashData(ctx context.Context, allHashData []api.HashData) error
	GetHashSum(ctx context.Context, dirFiles string) ([]models.HashDataFromDB, error)
	ChangedHashes(currentHashData []api.HashData, hashSumFromDB []models.HashDataFromDB) error
	CreateHash(path string) api.HashData
	WorkerPool(ctx context.Context, countWorkers int, jobs chan string, results chan api.HashData, logger *logrus.Logger)
	Worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan string, results chan<- api.HashData, logger *logrus.Logger)
}
