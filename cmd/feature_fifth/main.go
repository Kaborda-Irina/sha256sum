package main

import (
	"context"
	"flag"
	"fmt"
	config "github.com/Kaborda-Irina/sha256sum/internal/configs"
	"github.com/Kaborda-Irina/sha256sum/internal/initialize"
	"log"
	"time"
)

var dirPath string
var doHelp bool
var algorithm string
var checkHashSumFile string

//initializes the binding of the flag to a variable that must run before the main() function
func init() {
	flag.StringVar(&dirPath, "d", "", "directory path")
	flag.BoolVar(&doHelp, "h", false, "help")
	flag.StringVar(&algorithm, "a", "SHA256", "algorithm MD5, SHA1, SHA224, SHA256, SHA384, SHA512, default: SHA256")
	flag.StringVar(&checkHashSumFile, "c", "", "check hash sum files in directory")
}

func main() {
	flag.Parse()
	start := time.Now()
	//Initialize config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error during loading from config file", err)
	}
	ctx := context.Background()

	initialize.Initialize(ctx, cfg, doHelp, dirPath, algorithm, checkHashSumFile)
	fmt.Println(time.Since(start).Seconds())
}
