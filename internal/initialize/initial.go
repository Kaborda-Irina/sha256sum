package initialize

import (
	"context"
	"flag"
	"fmt"
	config "github.com/Kaborda-Irina/sha256sum/internal/configs"
	"github.com/Kaborda-Irina/sha256sum/internal/core/models"
	"github.com/Kaborda-Irina/sha256sum/internal/core/services"
	"github.com/Kaborda-Irina/sha256sum/internal/repositories"
	postrges "github.com/Kaborda-Irina/sha256sum/postgres"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
)

const countWorkers = 4

func Initialize(ctx context.Context, cfg config.Config, doHelp bool, dirPath string, algorithm string, checkHashSumFile string) {

	// Initialize PostgreSQL
	log.Println("Starting postgres connection")
	postgres, err := postrges.Initialize(cfg)
	if err != nil {
		log.Println("Failed to connection to Postgres", err)
	}
	log.Println("Postgres connection is successful")

	// Initialize repository
	repository := repositories.NewHashRepository(postgres)

	// Initialize service
	service := services.NewHashService(repository)

	jobs := make(chan string)
	results := make(chan models.HashData)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	algorithm = strings.ToUpper(algorithm)

	switch {
	//Initialize custom -h flag
	case doHelp:
		customHelpFlag()

	//Initialize custom -d flag
	case len(dirPath) > 0:
		go services.WorkerPool(ctx, countWorkers, algorithm, jobs, results)
		go services.SearchFilePath(ctx, dirPath, jobs)
		allHashData := services.Result(ctx, results, c)
		var wg sync.WaitGroup
		wg.Add(1)
		go func(ctx context.Context, allHashData []models.HashData) {
			defer wg.Done()
			err := service.SaveHashData(ctx, allHashData)
			if err != nil {
				return
			}
		}(ctx, allHashData)
		wg.Wait()

	//Initialize custom -c flag
	case len(checkHashSumFile) > 0:
		go services.WorkerPool(ctx, countWorkers, algorithm, jobs, results)
		go services.SearchFilePath(ctx, checkHashSumFile, jobs)
		allHashDataCurrent := services.ResultForCheck(ctx, results, c)
		allHashDataFromDB, err := service.GetHashSum(ctx, checkHashSumFile, algorithm)
		if err != nil {
			fmt.Println("Error getting hash data from database ", err)
		}
		resultChanged, resultDeleted, addedResult, err := service.ChangedHashes(allHashDataCurrent, allHashDataFromDB)
		if err != nil {
			fmt.Println("Error match data currently and data from db ", err)
		}

		printCheckResult(resultChanged, resultDeleted, addedResult)

	//If the user has not entered a flag
	default:
		fmt.Println("use the -h flag on the command line to see all the flags in this app")
	}
}

func customHelpFlag() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Custom help %s:\nYou can use the following flag:\n", os.Args[0])

		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stderr, "  flag -%v \n       %v\n", f.Name, f.Usage)
		})
	}
	flag.Usage()
}

func printCheckResult(resultChanged []models.ChangedHashes, resultDeleted []models.DeletedHashes, addedResult []models.AddedHashes) {
	if len(resultChanged) > 0 {
		fmt.Println("files hash sum changed:")
		for _, hash := range resultChanged {
			fmt.Printf("Changes were made to the file - %s located along the path %s, old hash sum %s, new hash sum %s\n",
				hash.FileName, hash.FilePath, hash.OldChecksum, hash.NewChecksum)
		}
	}
	if len(resultDeleted) > 0 {
		fmt.Println("files deleted:")
		for _, del := range resultDeleted {
			fmt.Printf("The file - %s was deleted at the specified path %s hash sum %s\n",
				del.FileName, del.FilePath, del.OldChecksum)
		}
	}
	if len(addedResult) > 0 {
		fmt.Println("files added:")
		for _, add := range addedResult {
			fmt.Printf("The file - %s has been added to the specified path %s hash sum %s\n",
				add.FileName, add.FilePath, add.NewChecksum)
		}
	}
	if len(resultChanged) == 0 && len(resultDeleted) == 0 && len(addedResult) == 0 {
		fmt.Println("Files didn't changed, added or deleted")
	}
}
