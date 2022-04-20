package services

import (
	"context"
	"fmt"
	"github.com/Kaborda-Irina/sha256sum/internal/core/ports"
	"github.com/Kaborda-Irina/sha256sum/internal/repositories"
	"github.com/Kaborda-Irina/sha256sum/pkg/api"
	"os"
)

const countWorkers = 4

type AppService struct {
	ports.IHashService
}

//NewAppService creates a new struct AppService
func NewAppService(r *repositories.AppRepository) *AppService {
	return &AppService{
		IHashService: NewHashService(r.IHashRepository),
	}
}

//StartGetHashData getting the hash sum of all files, outputs to os.Stdout and saves to the database
func (as *AppService) StartGetHashData(ctx context.Context, flagName string, algorithm string, jobs chan string, results chan api.HashData, sig chan os.Signal) {
	go api.WorkerPool(ctx, countWorkers, algorithm, jobs, results)
	go api.SearchFilePath(ctx, flagName, jobs)
	allHashData := api.Result(ctx, results, sig)
	err := as.IHashService.SaveHashData(ctx, allHashData)
	if err != nil {
		fmt.Println("Error save hash data to database ", err)
		return
	}

}

//StartCheckHashData getting the hash sum of all files, matches them and outputs to os.Stdout changes
func (as *AppService) StartCheckHashData(ctx context.Context, flagName string, algorithm string, jobs chan string, results chan api.HashData, sig chan os.Signal) {
	go api.WorkerPool(ctx, countWorkers, algorithm, jobs, results)
	go api.SearchFilePath(ctx, flagName, jobs)
	allHashDataCurrent := api.ResultForCheck(ctx, results, sig)
	allHashDataFromDB, err := as.IHashService.GetHashSum(ctx, flagName, algorithm)
	if err != nil {
		fmt.Println("Error getting hash data from database ", err)
		return
	}
	err = as.IHashService.ChangedHashes(allHashDataCurrent, allHashDataFromDB)
	if err != nil {
		fmt.Println("Error match data currently and data from db ", err)
		return
	}
}
