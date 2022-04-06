package main

import (
	"flag"
	"fmt"
	"github.com/Kaborda-Irina/sha256sum/internal/utils"
	"os"
	"sync"
	"time"
)

var dirPath string
var doHelp bool
var algorithm string

func init() {
	flag.StringVar(&dirPath, "d", "", "directory path")
	flag.BoolVar(&doHelp, "h", false, "help")
	flag.StringVar(&algorithm, "a", "", "algorithm md5, 1, 224, 256, 384, 512")
}

func main() {
	flag.Parse()
	switch {
	case doHelp:
		flag.Usage = func() {
			fmt.Fprintf(os.Stderr, "Custom help %s:\nYou can use the following flag:\n", os.Args[0])

			flag.VisitAll(func(f *flag.Flag) {
				fmt.Fprintf(os.Stderr, "  flag -%v \n       %v\n", f.Name, f.Usage)
			})
		}
		flag.Usage()
	case len(dirPath) > 0:
		start := time.Now()

		const countWorkers = 4
		jobs := make(chan string)
		results := make(chan string)

		go func() {
			var wg sync.WaitGroup
			for w := 1; w <= countWorkers; w++ {
				wg.Add(1)
				go utils.Worker(&wg, algorithm, jobs, results)
			}
			defer close(results)
			wg.Wait()
		}()

		go utils.SearchFilePath(dirPath, jobs)
		utils.Result(results)

		elapsed := time.Since(start)
		fmt.Printf("Took ===============> %s\n", elapsed)

	default:
		fmt.Println("use the -h flag on the command line to see all the flags in this app")
	}
}
