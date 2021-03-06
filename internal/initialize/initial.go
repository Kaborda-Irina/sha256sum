package initialize

import (
	"context"
	"flag"
	"fmt"
	"os"

	config "github.com/Kaborda-Irina/sha256sum/internal/configs"
	"github.com/Kaborda-Irina/sha256sum/internal/core/services"
	"github.com/Kaborda-Irina/sha256sum/internal/repositories"
	"github.com/Kaborda-Irina/sha256sum/pkg/api"

	"github.com/sirupsen/logrus"
)

func Initialize(ctx context.Context, cfg *config.Config, logger *logrus.Logger, sig chan os.Signal, doHelp bool, dirPath, algorithm, checkHashSumFile string) {
	// Initialize PostgreSQL
	logger.Info("Starting postgres connection")
	postgres, err := repositories.Initialize(cfg, logger)
	if err != nil {
		logger.Error("Failed to connection to Postgres", err)
	}
	logger.Info("Postgres connection is successful")

	// Initialize repository
	repository := repositories.NewAppRepository(postgres, logger)

	// Initialize service
	service, err := services.NewAppService(repository, algorithm, logger)
	if err != nil {
		logger.Fatalf("can't init service: %s", err)
	}

	jobs := make(chan string)
	results := make(chan api.HashData)

	switch {
	// Initialize custom -h flag
	case doHelp:
		customHelpFlag()
		return
	// Initialize custom -d flag
	case len(dirPath) > 0:
		err := service.StartGetHashData(ctx, dirPath, jobs, results, sig)
		if err != nil {
			logger.Error("Error when starting to get hash data ", err)
			return
		}
		return
	// Initialize custom -c flag
	case len(checkHashSumFile) > 0:
		err := service.StartCheckHashData(ctx, checkHashSumFile, jobs, results, sig)
		if err != nil {
			logger.Error("Error when starting to check hash data ", err)
			return
		}
		return
	// If the user has not entered a flag
	default:
		logger.Println("use the -h flag on the command line to see all the flags in this app")
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
