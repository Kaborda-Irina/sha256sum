package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Kaborda-Irina/sha256sum/internal/utils"
)

var filePath string
var dirPath string

func init() {
	flag.StringVar(&filePath, "f", "", "file path")
	flag.StringVar(&dirPath, "d", "", "dir path")

}

func main() {
	flag.Parse()

	c := make(chan string)
	d := make(chan []string)

	switch {
	case len(filePath) > 0:
		go utils.CreateSha256Sum(filePath, c)
		fmt.Println(<-c)

		if len(flag.Args()) > 0 {
			for _, nameArg := range flag.Args() {
				go utils.CreateSha256Sum(nameArg, c)
				fmt.Println(<-c)
			}
		}
	case len(dirPath) > 0:
		go utils.LookForAllFilePath(dirPath, d)
		allFilePaths := <-d
		for _, file := range allFilePaths {
			go utils.CreateSha256Sum(file, c)
			fmt.Println(<-c)
		}
		if len(flag.Args()) > 0 {
			for _, nameArg := range flag.Args() {
				go utils.LookForAllFilePath(nameArg, d)
				allFilePaths := <-d
				for _, file := range allFilePaths {
					go utils.CreateSha256Sum(file, c)
					fmt.Println(<-c)
				}
			}
		}
	default:
		errorCLS := errors.New("use the -d flag on the command line to find hash sum files or directories")
		fmt.Println(errorCLS)
	}
}
