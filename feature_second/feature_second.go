package feature_second

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func CreateSHA256() {
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Scan()
	path := myScanner.Text()

	f, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x", h.Sum(nil))
}
