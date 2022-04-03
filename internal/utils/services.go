package internal

import (
	"crypto/sha256"
	"fmt"
	"github.com/Kaborda-Irina/sha256sum/internal"
	"io"
	"os"
	"path/filepath"
)

func CreateSha256Sum(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", internal.ErrorFilePath
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", internal.ErrorHash
	}

	return fmt.Sprintf("%x %s", h.Sum(nil), filepath.Base(path)), nil
}
