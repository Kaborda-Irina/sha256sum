package main

import (
	"flag"
	"fmt"
	internal "github.com/Kaborda-Irina/sha256sum/internal/utils"
)

var filePath string

func init() {
	//initializes the binding of the flag to a variable that must run before the main() function
	flag.StringVar(&filePath, "f", "", "file path")
}

func main() {
	flag.Parse()
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

}
