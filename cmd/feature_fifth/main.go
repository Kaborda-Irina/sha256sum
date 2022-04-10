package main

import (
	"context"
	"flag"
	config "github.com/Kaborda-Irina/sha256sum/internal/configs"
	"github.com/Kaborda-Irina/sha256sum/internal/initialize"
	"log"
)

var dirPath string
var doHelp bool
var algorithm string

//initializes the binding of the flag to a variable that must run before the main() function
func init() {
	flag.StringVar(&dirPath, "d", "", "directory path")
	flag.BoolVar(&doHelp, "h", false, "help")
	flag.StringVar(&algorithm, "a", "", "algorithm md5, 1, 224, 256, 384, 512")
}

func main() {
	flag.Parse()

	//Initialize config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error during loading from config file")
	}
	ctx := context.Background()
	initialize.Initialize(ctx, cfg, doHelp, dirPath, algorithm)

}
