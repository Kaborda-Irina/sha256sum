package main

import (
	"fmt"
	"github.com/Kaborda-Irina/sha256sum/feature_third"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go feature_third.LookForPath(&wg)
	wg.Wait()
	fmt.Println("Goroutines finished executing")
}
