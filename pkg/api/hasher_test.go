package api

//
//import (
//	"github.com/Kaborda-Irina/sha256sum/internal/core/ports"
//	"github.com/Kaborda-Irina/sha256sum/internal/core/services"
//	"github.com/Kaborda-Irina/sha256sum/internal/repositories"
//	"github.com/Kaborda-Irina/sha256sum/pkg/hasher"
//	postrges "github.com/Kaborda-Irina/sha256sum/postgres"
//	"github.com/sirupsen/logrus"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//)
//
//func TestCreateHash(t *testing.T) {
//
//	testTable := []struct {
//		name     string
//		path     string
//		alg      string
//		expected HashData
//	}{
//		{"one", "../test.txt", "SHA256", HashData{
//			Hash:         "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
//			FileName:     "test.txt",
//			FullFilePath: "../test.txt",
//			Algorithm:    "SHA256",
//		}},
//	}
//	for _, testCase := range testTable {
//		l := logrus.New()
//		hashRep := repositories.NewHashRepository()
//		h, err := hasher.NewHashSum(testCase.alg)
//		if err != nil {
//			t.Error(err)
//			return
//		}
//
//		db := postrges.Initialize()
//		s := services.HashService{
//			ports.IHashRepository(),
//			h,
//			"SHA256",
//			l,
//		}
//		t.Run(testCase.name, func(t *testing.T) {
//
//			outputHashSum := s.CreateHash(testCase.path)
//			assert.Equal(t, testCase.expected, outputHashSum)
//		})
//	}
//}
