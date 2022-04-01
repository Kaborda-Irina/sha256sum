package internal

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"strings"
)

func CreateSha256Sum(path string) string {
	h := sha256.New()
	if _, err := io.Copy(h, strings.NewReader(path)); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("name file: %s, hash sum: %x", path, h.Sum(nil))
}
