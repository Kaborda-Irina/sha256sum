package api

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"time"
)

// SearchFilePath searches for all files in the given directory
func SearchFilePath(ctx context.Context, commonPath string, jobs chan<- string, logger *logrus.Logger) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err := filepath.Walk(commonPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logger.Error(err)
			return err
		}
		if !info.IsDir() {
			jobs <- path
		}
		return nil
	})
	close(jobs)

	if err != nil {
		logger.Error("not exist directory path", err)
	}
}

// CreateHash creates a hash sum of file depending on the algorithm
func CreateHash(path string, alg string, logger *logrus.Logger) HashData {
	f, err := os.Open(path)
	if err != nil {
		logger.Error("not exist file path", err)
	}
	defer f.Close()

	outputHashSum := HashData{}

	switch alg {
	case "MD5":
		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			logger.Error("error while copping file to hash", err)
		}
		outputHashSum.Hash = hex.EncodeToString(h.Sum(nil))

	case "SHA1":
		h := sha1.New()
		if _, err := io.Copy(h, f); err != nil {
			logger.Error("error while copping file to hash", err)
		}
		outputHashSum.Hash = hex.EncodeToString(h.Sum(nil))

	case "SHA224":
		h := sha256.New224()
		if _, err := io.Copy(h, f); err != nil {
			logger.Error("error while copping file to hash", err)
		}
		outputHashSum.Hash = hex.EncodeToString(h.Sum(nil))

	case "SHA384":
		h := sha512.New384()
		if _, err := io.Copy(h, f); err != nil {
			logger.Error("error while copping file to hash", err)
		}
		outputHashSum.Hash = hex.EncodeToString(h.Sum(nil))

	case "SHA512":
		h := sha512.New()
		if _, err := io.Copy(h, f); err != nil {
			logger.Error("error while copping file to hash", err)
		}
		outputHashSum.Hash = hex.EncodeToString(h.Sum(nil))

	default:
		h := sha256.New()
		if _, err := io.Copy(h, f); err != nil {
			logger.Error("error while copping file to hash", err)
		}
		outputHashSum.Hash = hex.EncodeToString(h.Sum(nil))
		alg = "SHA256"
	}

	outputHashSum.FileName = filepath.Base(path)
	outputHashSum.FullFilePath = path
	outputHashSum.Algorithm = alg
	return outputHashSum
}

// Result launching an infinite loop of receiving and outputting to Stdout the result and signal control
func Result(ctx context.Context, results chan HashData, c chan os.Signal) []HashData {
	var allHashData []HashData
	for {
		select {
		case hashData, ok := <-results:
			if !ok {
				return allHashData
			}
			fmt.Printf("%s %s\n", hashData.Hash, hashData.FileName)
			allHashData = append(allHashData, hashData)
		case <-c:
			fmt.Println("exit program")
			return []HashData{}
		case <-ctx.Done():
		}
	}
}

// ResultForCheck launching an infinite loop of receiving and return all hash sum
func ResultForCheck(ctx context.Context, results chan HashData, c chan os.Signal) []HashData {
	var allHashData []HashData
	for {
		select {
		case hashData, ok := <-results:
			if !ok {
				return allHashData
			}
			allHashData = append(allHashData, hashData)
		case <-c:
			fmt.Println("exit program")
			return nil
		case <-ctx.Done():
		}
	}
}
