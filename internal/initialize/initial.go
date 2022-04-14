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

	switch {
	//Initialize custom -h flag
	case doHelp:
		customHelpFlag()
	//Initialize custom -d flag
	case len(dirPath) > 0:
		go services.WorkerPool(ctx, countWorkers, algorithm, jobs, results)
		go services.SearchFilePath(ctx, dirPath, jobs)
		allHashData := services.Result(ctx, results, c, service)
		err := service.SaveHashDir(ctx, allHashData)
		if err != nil {
			return
		}

	//Initialize custom -c flag
	case len(checkHashSumFile) > 0:
		go services.WorkerPool(ctx, countWorkers, algorithm, jobs, results)
		go services.SearchFilePath(ctx, checkHashSumFile, jobs)
		//	allHashDataCurrent := services.Result(ctx, results, c, service)
		allHashDataCurrent := services.ResultForCheck(ctx, results, c, service)
		allHashDataFromDB, err := service.GetHashSum(ctx, allHashDataCurrent)
		if err != nil {
			fmt.Println("Error getting hash data from database ", err)
		}
		result := services.MatchHashSum(allHashDataCurrent, allHashDataFromDB)
		fmt.Println(result)
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
