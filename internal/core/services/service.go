package services

import (
	"context"
	"github.com/Kaborda-Irina/sha256sum/internal/core/ports"
	"github.com/Kaborda-Irina/sha256sum/internal/repositories"
	"github.com/Kaborda-Irina/sha256sum/pkg/api"
	"github.com/sirupsen/logrus"
	"os"
)

const countWorkers = 4

type AppService struct {
	ports.IHashService
	logger *logrus.Logger
}

//NewAppService creates a new struct AppService
func NewAppService(r *repositories.AppRepository, logger *logrus.Logger) *AppService {
	return &AppService{
		IHashService: NewHashService(r.IHashRepository, logger),
		logger:       logger,
	}
}

//StartGetHashData getting the hash sum of all files, outputs to os.Stdout and saves to the database
func (as *AppService) StartGetHashData(ctx context.Context, flagName string, algorithm string, jobs chan string, results chan api.HashData, sig chan os.Signal) {
	go api.WorkerPool(ctx, countWorkers, algorithm, jobs, results, as.logger)
	go api.SearchFilePath(ctx, flagName, jobs, as.logger)
	allHashData := api.Result(ctx, results, sig)
	err := as.IHashService.SaveHashData(ctx, allHashData)
	if err != nil {
		as.logger.Error("Error save hash data to database ", err)
		return
	}

}

//StartCheckHashData getting the hash sum of all files, matches them and outputs to os.Stdout changes
func (as *AppService) StartCheckHashData(ctx context.Context, flagName string, algorithm string, jobs chan string, results chan api.HashData, sig chan os.Signal) {
	go api.WorkerPool(ctx, countWorkers, algorithm, jobs, results, as.logger)
	go api.SearchFilePath(ctx, flagName, jobs, as.logger)
	allHashDataCurrent := api.ResultForCheck(ctx, results, sig)
	allHashDataFromDB, err := as.IHashService.GetHashSum(ctx, flagName, algorithm)
	if err != nil {
		as.logger.Error("Error getting hash data from database ", err)
		return
	}
	err = as.IHashService.ChangedHashes(allHashDataCurrent, allHashDataFromDB)
	if err != nil {
		as.logger.Error("Error match data currently and data from db ", err)
		return
	}
}
