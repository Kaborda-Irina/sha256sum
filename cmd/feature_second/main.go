package main

import (
	"errors"
	"flag"
	"fmt"
	internal "github.com/Kaborda-Irina/sha256sum/internal/utils"
)

var filePath string
var dirPath string

func init() {
	//initializes the binding of the flag to a variable that must run before the main() function
	flag.StringVar(&filePath, "f", "", "file path")
	flag.StringVar(&dirPath, "d", "", "dir path")
}

func main() {
	flag.Parse()

	switch {
	case len(filePath) > 0:
		result, err := internal.CreateSha256Sum(filePath)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)

		if len(flag.Args()) > 0 {
			for _, nameArg := range flag.Args() {
				result, err := internal.CreateSha256Sum(nameArg)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(result)
			}
		}
	case len(dirPath) > 0:
		allFilePaths, err := internal.LookForAllFilePath(dirPath)
		if err != nil {
			fmt.Println(err)
		}
		for _, file := range allFilePaths {
			result, err := internal.CreateSha256Sum(file)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(result)
		}
		if len(flag.Args()) > 0 {
			for _, nameArg := range flag.Args() {
				allFilePaths, err := internal.LookForAllFilePath(nameArg)
				if err != nil {
					fmt.Println(err)
				}
				for _, file := range allFilePaths {
					result, err := internal.CreateSha256Sum(file)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(result)
				}
			}
		}
	default:
		errorCLS := errors.New("Use the -d flag on the command line to find hash sum files or directories")
		fmt.Println(errorCLS)
	}
}
