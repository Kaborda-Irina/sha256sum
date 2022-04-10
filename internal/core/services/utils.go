package services

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"github.com/Kaborda-Irina/sha256sum/internal"
	"github.com/Kaborda-Irina/sha256sum/internal/core/models"
	"github.com/Kaborda-Irina/sha256sum/internal/core/ports"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// WorkerPool launches a certain number of workers for concurrent processing
func WorkerPool(ctx context.Context, countWorkers int, algorithm string, jobs chan string, results chan models.HashSum, service ports.IHashService) {
	var wg sync.WaitGroup
	for w := 1; w <= countWorkers; w++ {
		wg.Add(1)
		go Worker(ctx, &wg, algorithm, jobs, results, service)
	}
	defer close(results)
	wg.Wait()
}

// SearchFilePath searches for all files in the given directory
func SearchFilePath(commonPath string, jobs chan<- string) {
	err := filepath.Walk(commonPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return internal.ErrorDirPath
		}
		if !info.IsDir() {
			jobs <- path
		}
		return nil
	})
	close(jobs)

	if err != nil {
		log.Println(internal.ErrorDirPath)
	}
}

// Worker gets jobs from a pipe and writes the result to stdout and database
func Worker(ctx context.Context, wg *sync.WaitGroup, algorithm string, jobs <-chan string, results chan<- models.HashSum, service ports.IHashService) {
	defer wg.Done()
	for j := range jobs {
		hashSum := CreateHash(j, algorithm)
		results <- hashSum
		err := service.SaveHashSum(hashSum, ctx)
		if err != nil {
			log.Println("error while save hash in db")
		}
	}

}

// CreateHash creates a hash sum of file depending on the algorithm
func CreateHash(path string, alg string) models.HashSum {
	f, err := os.Open(path)
	if err != nil {
		log.Println(internal.ErrorFilePath)
	}
	defer f.Close()

	outputHashSum := models.HashSum{}

	switch strings.ToLower(alg) {
	case "md5":
		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Println(internal.ErrorHash)
		}
		outputHashSum.Hash = h.Sum(nil)
	case "1":
		h := sha1.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Println(internal.ErrorHash)
		}
		outputHashSum.Hash = h.Sum(nil)
	case "224":
		h := sha256.New224()
		if _, err := io.Copy(h, f); err != nil {
			log.Println(internal.ErrorHash)
		}
		outputHashSum.Hash = h.Sum(nil)
	case "384":
		h := sha512.New384()
		if _, err := io.Copy(h, f); err != nil {
			log.Println(internal.ErrorHash)
		}
		outputHashSum.Hash = h.Sum(nil)
	case "512":
		h := sha512.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Println(internal.ErrorHash)
		}
		outputHashSum.Hash = h.Sum(nil)
	default:
		h := sha256.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Println(internal.ErrorHash)
		}
		outputHashSum.Hash = h.Sum(nil)
	}

	outputHashSum.FileName = filepath.Base(path)
	return outputHashSum
}

// Result launching an infinite loop of receiving and outputting to Stdout the result and signal control
func Result(ctx context.Context, results chan models.HashSum, c chan os.Signal) {

	for {
		select {
		case outputHashSum, ok := <-results:
			if !ok {
				return
			}
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("%x %s\n", outputHashSum.Hash, outputHashSum.FileName)
		case <-c:
			fmt.Println("exit program")
			return
		case <-ctx.Done():
		}
	}

}
