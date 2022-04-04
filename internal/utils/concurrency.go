package utils

import (
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
)

func CreateSha256Sum(path string, alg string, c chan<- string) {
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

	c <- fmt.Sprintf("%x %s", sum, filepath.Base(path))
}

func LookForAllFilePath(commonPath string, d chan []string) {
	var allFilePath []string
	err := filepath.Walk(commonPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return internal.ErrorDirPath
		}
		if !info.IsDir() {
			allFilePath = append(allFilePath, path)
		}

		return nil
	})

	if err != nil {
		log.Println(internal.ErrorDirPath)
	}
	d <- allFilePath
}
