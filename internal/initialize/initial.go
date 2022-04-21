package initialize

import (
	"context"
	"flag"
	"fmt"
	config "github.com/Kaborda-Irina/sha256sum/internal/configs"
	"github.com/Kaborda-Irina/sha256sum/internal/core/services"
	"github.com/Kaborda-Irina/sha256sum/internal/repositories"
	"github.com/Kaborda-Irina/sha256sum/pkg/api"
	postrges "github.com/Kaborda-Irina/sha256sum/postgres"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func Initialize(ctx context.Context, cfg config.Config, logger *logrus.Logger, sig chan os.Signal, doHelp bool, dirPath string, algorithm string, checkHashSumFile string) {

	// Initialize PostgreSQL
	logger.Info("Starting postgres connection")
	postgres, err := postrges.Initialize(cfg, logger)
	if err != nil {
		logger.Error("Failed to connection to Postgres", err)
	}
	logger.Info("Postgres connection is successful")

	// Initialize repository
	repository := repositories.NewAppRepository(postgres, logger)

	// Initialize service
	service := services.NewAppService(repository, logger)

	jobs := make(chan string)
	results := make(chan api.HashData)
	algorithm = strings.ToUpper(algorithm)

	switch {
	//Initialize custom -h flag
	case doHelp:
		customHelpFlag()

	//Initialize custom -d flag
	case len(dirPath) > 0:
		service.StartGetHashData(ctx, dirPath, algorithm, jobs, results, sig)

	//Initialize custom -c flag
	case len(checkHashSumFile) > 0:
		service.StartCheckHashData(ctx, checkHashSumFile, algorithm, jobs, results, sig)

	//If the user has not entered a flag
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
