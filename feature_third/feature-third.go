package feature_third

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func LookForPath(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("enter path or file")
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Scan()
	commonPath := myScanner.Text()
	commonPaths := strings.Split(commonPath, " ")
	for _, path := range commonPaths {
		fmt.Printf("path : %s\n", path)
		wg.Add(1)
		go NavigateThroughFolders(path, wg)
	}
}
func CreateSHA256(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	h := sha256.New()
	if _, err := io.Copy(h, strings.NewReader(path)); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name file: %s, hash sum: %x\n", path, h.Sum(nil))
}

func NavigateThroughFolders(commonPath string, wg *sync.WaitGroup) {
	defer wg.Done()
	err := filepath.Walk(commonPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if !info.IsDir() {
			wg.Add(1)
			go CreateSHA256(path, wg /*c*/)
			time.Sleep(1 * time.Second)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
