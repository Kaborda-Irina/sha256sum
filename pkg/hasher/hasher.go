package hasher

import (
	"io"
)

type IHasher interface {
	Hash(file io.Reader) (string, error)
}

func NewHashSum(alg string) (h IHasher, err error) {
	switch alg {
	case "MD5":
		h = NewMD5()
	case "SHA1":
		h = NewSHA1()
	case "SHA224":
		h = NewSHA224()
	case "SHA384":
		h = NewSHA384()
	case "SHA512":
		h = NewSHA512()
	default:
		h = NewSHA256()
	}
	return h, nil
}
