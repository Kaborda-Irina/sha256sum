package internal

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateSha256Sum(path string) string {
	h := sha256.New()
	if _, err := io.Copy(h, strings.NewReader(path)); err != nil {
		ErrorProcessing(err)
	}
	return fmt.Sprintf("name file: %s, hash sum: %x", path, h.Sum(nil))
}

func LookForAllFilePath(commonPath string /* func(p string, info os.FileInfo) */) []string {
	var allFilePath []string
	err := filepath.Walk(commonPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			allFilePath = append(allFilePath, info.Name())
		}

		return nil
	})
	if err != nil {
		ErrorProcessing(err)
	}
	return allFilePath
}

func ErrorProcessing(err error) {
	log.Printf("Error execution: %s", err)
}
