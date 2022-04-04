package main

import (
	"flag"
	"fmt"
	"github.com/Kaborda-Irina/sha256sum/internal/utils"
)

var dirPath string
var doHelp bool
var algorithm string

func init() {
	flag.StringVar(&dirPath, "d", "", "dir path")
	flag.BoolVar(&doHelp, "h", false, "help")
	flag.StringVar(&algorithm, "a", "", "algorithm md5, 1, 224, 256, 384, 512")
}

func main() {
	flag.Parse()

	c := make(chan string)
	d := make(chan []string)

	switch {
	case doHelp:
		flag.Usage()
	case len(dirPath) > 0:
		go utils.LookForAllFilePath(dirPath, d)
		allFilePaths := <-d
		for _, file := range allFilePaths {
			go utils.CreateSha256Sum(file, algorithm, c)
			fmt.Println(<-c)
		}
		if len(flag.Args()) > 0 {
			for _, nameArg := range flag.Args() {
				go utils.LookForAllFilePath(nameArg, d)
				allFilePaths := <-d
				for _, file := range allFilePaths {
					go utils.CreateSha256Sum(file, algorithm, c)
					fmt.Println(<-c)
				}
			}
		}
	default:
		fmt.Println("use the -h flag on the command line to see all the flags in this app")
	}
}
