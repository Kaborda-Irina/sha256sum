package main

import (
	"flag"
	"fmt"
	"github.com/Kaborda-Irina/sha256sum/internal/utils"
	"testing"
)

const testPath = "../testPath"

//go test -bench=. main_test.go -benchmem
func BenchmarkDefault(b *testing.B) {
	flag.Parse()
	c := make(chan string)
	d := make(chan []string)
	algorithm := "256"

	b.SetBytes(1)
	b.ResetTimer()
	for i := 0; i < 1000; i++ {
		go utils.LookForAllFilePath(testPath, d)
		allFilePaths := <-d
		for _, file := range allFilePaths {
			go utils.CreateSha256Sum(file, algorithm, c)
			fmt.Println(<-c)
		}
	}
	b.StopTimer()
}
