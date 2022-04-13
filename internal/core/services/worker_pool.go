package services

import (
	"context"
	"github.com/Kaborda-Irina/sha256sum/internal/core/models"
	"sync"
	"time"
)

// WorkerPool launches a certain number of workers for concurrent processing
func WorkerPool(ctx context.Context, countWorkers int, algorithm string, jobs chan string, results chan models.HashData) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	for w := 1; w <= countWorkers; w++ {
		wg.Add(1)
		go Worker(ctx, &wg, algorithm, jobs, results)
	}
	defer close(results)
	wg.Wait()
}

// Worker gets jobs from a pipe and writes the result to stdout and database
func Worker(ctx context.Context, wg *sync.WaitGroup, algorithm string, jobs <-chan string, results chan<- models.HashData) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	defer wg.Done()
	for j := range jobs {
		hashSum := CreateHash(j, algorithm)
		results <- hashSum
	}
}
