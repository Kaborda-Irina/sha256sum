package utils

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"github.com/Kaborda-Irina/sha256sum/internal"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func CreateHash(path string, alg string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Println(internal.ErrorFilePath)
	}
	defer f.Close()

	var sum interface{}
	switch strings.ToLower(alg) {
	case "md5":
		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Println(internal.ErrorHash)
		}
		sum = h.Sum(nil)
	case "1":
		h := sha1.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Println(internal.ErrorHash)
		}
		sum = h.Sum(nil)
	case "224":
		h := sha256.New224()
		if _, err := io.Copy(h, f); err != nil {
			log.Println(internal.ErrorHash)
		}
		sum = h.Sum(nil)
	case "384":
		h := sha512.New384()
		if _, err := io.Copy(h, f); err != nil {
			log.Println(internal.ErrorHash)
		}
		sum = h.Sum(nil)
	case "512":
		h := sha512.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Println(internal.ErrorHash)
		}
		sum = h.Sum(nil)
	default:
		h := sha256.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Println(internal.ErrorHash)
		}
		sum = h.Sum(nil)
	}

	return fmt.Sprintf("%x %s", sum, filepath.Base(path))
}

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

func Worker(wg *sync.WaitGroup, algorithm string, jobs <-chan string, results chan<- string) {
	defer wg.Done()
	for j := range jobs {
		results <- CreateHash(j, algorithm)
	}

}

func Result(ctx context.Context, results chan string, c chan os.Signal) {

	for {
		select {
		case hash, ok := <-results:
			if !ok {
				return
			}
			time.Sleep(500 * time.Millisecond)
			fmt.Println(hash)
		case <-c:
			fmt.Println(" exit program")
			return
		case <-ctx.Done():
		}
	}

}
