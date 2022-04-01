package main

import (
	"errors"
	"flag"
	"fmt"
	internal "github.com/Kaborda-Irina/sha256sum/internal/utils"
	"strings"
)

var filePath = flag.String("long path", "", "file path")
var dirPath = flag.String("long dir", "", "dir path")

func init() {
	flag.StringVar(filePath, "f", "", "file path")
	flag.StringVar(dirPath, "d", "", "dir path")
}

func main() {
	flag.Parse()

	switch {
	case len(*filePath) > 0:
		if strings.Contains(*filePath, ",") {
			filePaths := strings.Split(*filePath, ",")
			for _, fPath := range filePaths {
				fmt.Println(internal.CreateSha256Sum(fPath))
			}
		} else {
			fmt.Println(internal.CreateSha256Sum(*filePath))
		}
	case len(*dirPath) > 0:
		if strings.Contains(*dirPath, ",") {
			disPaths := strings.Split(*dirPath, ",")
			for _, dPath := range disPaths {
				allFilePaths := internal.LookForAllFilePath(dPath)
				for _, file := range allFilePaths {
					fmt.Println(internal.CreateSha256Sum(file))
				}
			}
		} else {
			allFilePaths := internal.LookForAllFilePath(*dirPath)
			for _, file := range allFilePaths {
				fmt.Println(internal.CreateSha256Sum(file))
			}
		}
	default:
		err := errors.New("error in command line, need to write file path or dir")
		internal.ErrorProcessing(err)
	}
}
