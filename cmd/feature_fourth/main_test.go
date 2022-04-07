package main

import (
	"context"
	"github.com/Kaborda-Irina/sha256sum/internal/utils"
	"os"
	"os/signal"
	"sync"
	"testing"
)

const testPath = "../../../h/h1"

//go test -bench=. main_test.go -benchmem
func BenchmarkDefault(b *testing.B) {

	b.SetBytes(1)
	b.ResetTimer()
	const countWorkers = 4
	jobs := make(chan string)
	results := make(chan string)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	defer func() {
		signal.Stop(c)
		cancel()
	}()

	go func() {
		var wg sync.WaitGroup
		for w := 1; w <= countWorkers; w++ {
			wg.Add(1)
			go utils.Worker(&wg, algorithm, jobs, results)
		}
		defer close(results)
		wg.Wait()
	}()

	go utils.SearchFilePath(testPath, jobs)

	utils.Result(ctx, results, c)

	b.StopTimer()
}
