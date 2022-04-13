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
	"time"
)

// SearchFilePath searches for all files in the given directory
func SearchFilePath(ctx context.Context, commonPath string, jobs chan<- string) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err := filepath.Walk(commonPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			jobs <- path
		}
		return nil
	})
	close(jobs)

	if err != nil {
		log.Println(internal.ErrorDirPath, err)
	}
}

// CreateHash creates a hash sum of file depending on the algorithm
func CreateHash(path string, alg string) models.HashData {
	f, err := os.Open(path)
	if err != nil {
		log.Println(internal.ErrorFilePath)
	}
	defer f.Close()

	outputHashSum := models.HashData{}
	alg = strings.ToUpper(alg)
	switch alg {
	case "MD5":
		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Println(internal.ErrorHash)
		}
		outputHashSum.Hash = h.Sum(nil)
		outputHashSum.Algorithm = alg
	case "SHA1":
		h := sha1.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Println(internal.ErrorHash)
		}
		outputHashSum.Hash = h.Sum(nil)
		outputHashSum.Algorithm = alg
	case "SHA224":
		h := sha256.New224()
		if _, err := io.Copy(h, f); err != nil {
			log.Println(internal.ErrorHash)
		}
		outputHashSum.Hash = h.Sum(nil)
		outputHashSum.Algorithm = alg
	case "SHA384":
		h := sha512.New384()
		if _, err := io.Copy(h, f); err != nil {
			log.Println(internal.ErrorHash)
		}
		outputHashSum.Hash = h.Sum(nil)
		outputHashSum.Algorithm = alg
	case "SHA512":
		h := sha512.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Println(internal.ErrorHash)
		}
		outputHashSum.Hash = h.Sum(nil)
		outputHashSum.Algorithm = alg
	default:
		h := sha256.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Println(internal.ErrorHash)
		}
		outputHashSum.Hash = h.Sum(nil)
		outputHashSum.Algorithm = "SHA256"
	}

	outputHashSum.FileName = filepath.Base(path)
	outputHashSum.FullFilePath = path

	return outputHashSum
}

// Result launching an infinite loop of receiving and outputting to Stdout the result and signal control
func Result(ctx context.Context, results chan models.HashData, c chan os.Signal, service ports.IHashService) {

	for {
		select {
		case hashData, ok := <-results:
			if !ok {
				return
			}
			fmt.Printf("%x %s\n", hashData.Hash, hashData.FileName)
			err := service.SaveHashSum(ctx, hashData)
			if err != nil {
				log.Println("error while save hash in db", err)
			}
		case <-c:
			fmt.Println("exit program")
			return
		case <-ctx.Done():
		}
	}

}

func ResultForCheck(ctx context.Context, results chan models.HashData, c chan os.Signal, service ports.IHashService) []models.HashData {

	for {
		select {
		case currentHashData, ok := <-results:
			if !ok {
				return []models.HashData{}
			}

			hashDataFromDB, err := service.GetHashSum(ctx, currentHashData.FullFilePath, currentHashData.Algorithm)
			if err != nil {
				fmt.Println("Error getting hash data from database ", err)
			}
			matchHashSum(currentHashData, hashDataFromDB)
		case <-c:
			fmt.Println("exit program")
			return []models.HashData{}
		case <-ctx.Done():
		}
	}
}
func matchHashSum(currentHashData models.HashData, hashSumFromDB models.HashDataFromDB) /*[]models.HashData*/ {
	hashSumCurrent := fmt.Sprintf("%x", currentHashData.Hash)

	if hashSumCurrent != hashSumFromDB.Hash {
		fmt.Printf("Changes were made to the file - %v located along the path %v\n", currentHashData.FileName, currentHashData.FullFilePath)
	}
	if currentHashData.FullFilePath != hashSumFromDB.FullFilePath {
		fmt.Printf("File deleted  - %v located along the path %v\n", currentHashData.FileName, currentHashData.FullFilePath)
	}
}
