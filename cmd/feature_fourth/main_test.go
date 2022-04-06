package main

import (
	"github.com/Kaborda-Irina/sha256sum/internal/utils"
	"sync"
	"testing"
)

const testPath = "../../../h"

//go test -bench=. main_test.go -benchmem
func BenchmarkDefault(b *testing.B) {

	b.SetBytes(1)
	b.ResetTimer()
	const countWorkers = 5
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

	go utils.SearchFilePath(testPath, jobs)
	utils.Result(results)
	b.StopTimer()
}
