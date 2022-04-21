package api

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateHash(t *testing.T) {

	testTable := []struct {
		name     string
		path     string
		alg      string
		expected HashData
	}{
		{"one", "../test.txt", "SHA256", HashData{
			Hash:         "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			FileName:     "test.txt",
			FullFilePath: "../test.txt",
			Algorithm:    "SHA256",
		}},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			f, err := os.Open(testCase.path)
			if err != nil {
				assert.Error(t, err)
			}
			defer f.Close()

			outputHashSum := HashData{}

			switch testCase.alg {
			case "MD5":
				h := md5.New()
				if _, err := io.Copy(h, f); err != nil {
					assert.Error(t, err)
				}
				outputHashSum.Hash = hex.EncodeToString(h.Sum(nil))

			case "SHA1":
				h := sha1.New()
				if _, err := io.Copy(h, f); err != nil {
					assert.Error(t, err)
				}
				outputHashSum.Hash = hex.EncodeToString(h.Sum(nil))

			case "SHA224":
				h := sha256.New224()
				if _, err := io.Copy(h, f); err != nil {
					assert.Error(t, err)
				}
				outputHashSum.Hash = hex.EncodeToString(h.Sum(nil))

			case "SHA384":
				h := sha512.New384()
				if _, err := io.Copy(h, f); err != nil {
					assert.Error(t, err)
				}
				outputHashSum.Hash = hex.EncodeToString(h.Sum(nil))

			case "SHA512":
				h := sha512.New()
				if _, err := io.Copy(h, f); err != nil {
					assert.Error(t, err)
				}
				outputHashSum.Hash = hex.EncodeToString(h.Sum(nil))

			default:
				h := sha256.New()
				if _, err := io.Copy(h, f); err != nil {
					assert.Error(t, err)
				}
				outputHashSum.Hash = hex.EncodeToString(h.Sum(nil))
				testCase.alg = "SHA256"
			}

			outputHashSum.FileName = filepath.Base(testCase.path)
			outputHashSum.FullFilePath = testCase.path
			outputHashSum.Algorithm = testCase.alg

			assert.Equal(t, testCase.expected, outputHashSum)
		})
	}
}
