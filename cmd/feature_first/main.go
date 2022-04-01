package main

import (
	"flag"
	"fmt"
	internal "github.com/Kaborda-Irina/sha256sum/internal/utils"
	"log"
	"strings"
)

var filePath = flag.String("long path", "", "file path")

func init() {
	flag.StringVar(filePath, "f", "", "file path")
}

func main() {
	flag.Parse()

	switch {
	case len(*filePath) > 0:
		if strings.Contains(*filePath, ",") {
			filePaths := strings.Split(*filePath, ",")
			for _, fPath := range filePaths {
				result := internal.CreateSha256Sum(fPath)
				fmt.Println(result)
			}
		} else {
			result := internal.CreateSha256Sum(*filePath)
			fmt.Println(result)
		}
	default:
		log.Println("error in command line, need write file path")

	}

}
