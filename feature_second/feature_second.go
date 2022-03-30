package feature_second

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func LookForPath() {
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Scan()
	commonPath := myScanner.Text()
	NavigateThroughFolders(commonPath)
}
func CreateSHA256(path string) {
	h := sha256.New()
	if _, err := io.Copy(h, strings.NewReader(path)); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name file: %s, hash sum: %x\n", path, h.Sum(nil))
}

func NavigateThroughFolders(commonPath string) {
	err := filepath.Walk(commonPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if !info.IsDir() {
			CreateSHA256(path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
