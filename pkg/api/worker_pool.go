package api

import (
	"context"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

// WorkerPool launches a certain number of workers for concurrent processing
func WorkerPool(ctx context.Context, countWorkers int, algorithm string, jobs chan string, results chan HashData, logger *logrus.Logger) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	for w := 1; w <= countWorkers; w++ {
		wg.Add(1)
		go Worker(ctx, &wg, algorithm, jobs, results, logger)
	}
	defer close(results)
	wg.Wait()
}

// Worker gets jobs from a pipe and writes the result to stdout and database
func Worker(ctx context.Context, wg *sync.WaitGroup, algorithm string, jobs <-chan string, results chan<- HashData, logger *logrus.Logger) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	defer wg.Done()
	for j := range jobs {
		hashSum := CreateHash(j, algorithm, logger)
		results <- hashSum
	}
}
