package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	createSha256Sum()
}

func createSha256Sum() {
	h := sha256.New()
	s, err := ioutil.ReadFile(os.Args[1])
	h.Write(s)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.WriteString(hex.EncodeToString(h.Sum(nil)))
}
