package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	createSHA256()
}

func createSHA256() {

	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	path := myscanner.Text()

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