package services

import (
	"context"
	"errors"
	"github.com/Kaborda-Irina/sha256sum/internal/core/models"
	mock_ports "github.com/Kaborda-Irina/sha256sum/internal/core/ports/mocks"
	"github.com/Kaborda-Irina/sha256sum/pkg/api"
	"github.com/Kaborda-Irina/sha256sum/pkg/hasher"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestNewHashService(t *testing.T) {
	alg := "SHA256"
	logger := logrus.New()
	expected := HashService{}

	c := gomock.NewController(t)
	defer c.Finish()

	repo := mock_ports.NewMockIHashRepository(c)
	h, err := hasher.NewHashSum(alg)
	if err != nil {
		require.Error(t, err)
	}
	hashService := HashService{
		hashRepository: repo,
		hasher:         h,
		alg:            alg,
		logger:         logger,
	}
	assert.NotEqual(t, expected, hashService, "they should not be equal")
}

func TestCreateHash(t *testing.T) {
	testTable := []struct {
		name         string
		alg          string
		path         string
		mockBehavior func(s *mock_ports.MockIHashService, path string)
		expected     api.HashData
	}{
		{
			name: "exist path",
			alg:  "SHA256",
			path: "../h/h1/test.txt",
			mockBehavior: func(s *mock_ports.MockIHashService, path string) {
				s.EXPECT().CreateHash(path).Return(api.HashData{
					Hash:         "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
					FileName:     "test.txt",
					FullFilePath: "../h/h1/test.txt",
					Algorithm:    "SHA256",
				})

			},
			expected: api.HashData{
				Hash:         "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
				FileName:     "test.txt",
				FullFilePath: "../h/h1/test.txt",
				Algorithm:    "SHA256",
			},
		},
		{
			name: "not exist path",
			alg:  "SHA256",
			path: "/test.txx",
			mockBehavior: func(s *mock_ports.MockIHashService, path string) {
				s.EXPECT().CreateHash(path).Return(api.HashData{})
			},
			expected: api.HashData{
				Hash:         "",
				FileName:     "",
				FullFilePath: "",
				Algorithm:    "",
			},
		},
		{
			name: "error in a hash sum",
			alg:  "SHA256",
			path: "/test.txx",
			mockBehavior: func(s *mock_ports.MockIHashService, path string) {
				s.EXPECT().CreateHash(path).Return(api.HashData{})
			},
			expected: api.HashData{},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_ports.NewMockIHashService(c)
			testCase.mockBehavior(service, testCase.path)

			file, err := os.Open(testCase.path)
			if err != nil {
				require.Error(t, err)
			}
			defer file.Close()

			result := service.CreateHash(testCase.path)

			assert.Equal(t, testCase.expected, result)
		})
	}
}

func TestSaveHashData(t *testing.T) {
	type mockBehavior func(r *mock_ports.MockIHashRepository, ctx context.Context, allHashData []api.HashData)
	testTable := []struct {
		name         string
		alg          string
		allHashData  []api.HashData
		mockBehavior mockBehavior
		expected     error
	}{
		{
			name: "exist path",
			alg:  "SHA256",
			allHashData: []api.HashData{{
				Hash:         "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
				FileName:     "new",
				FullFilePath: "../h/h1/new",
				Algorithm:    "SHA256",
			}},
			mockBehavior: func(r *mock_ports.MockIHashRepository, ctx context.Context, allHashData []api.HashData) {
				r.EXPECT().SaveHashData(ctx, allHashData).Return(nil)
			},
			expected: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			l := logrus.New()
			ctx := context.Background()
			repo := mock_ports.NewMockIHashRepository(c)
			h, err := hasher.NewHashSum(testCase.alg)
			if err != nil {
				assert.Error(t, err)
			}

			hashService := HashService{
				hashRepository: repo,
				hasher:         h,
				alg:            testCase.alg,
				logger:         l,
			}
			testCase.mockBehavior(repo, ctx, testCase.allHashData)

			err = hashService.hashRepository.SaveHashData(ctx, testCase.allHashData)
			if testCase.expected != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetHashSum(t *testing.T) {
	type mockBehavior func(r *mock_ports.MockIHashRepository, ctx context.Context, dirFiles, alg string)
	testTable := []struct {
		name         string
		alg          string
		dirFiles     string
		expected     []models.HashDataFromDB
		mockBehavior mockBehavior
		expectedErr  bool
	}{
		{
			name:     "exist path",
			alg:      "SHA256",
			dirFiles: "../h/h1/new",
			expected: []models.HashDataFromDB{{
				ID:           1,
				Hash:         "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
				FileName:     "new",
				FullFilePath: "../h/h1/new",
				Algorithm:    "SHA256",
			}},
			mockBehavior: func(r *mock_ports.MockIHashRepository, ctx context.Context, dirFiles, alg string) {
				r.EXPECT().GetHashSum(ctx, dirFiles, alg).Return([]models.HashDataFromDB{
					{
						ID:           1,
						Hash:         "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
						FileName:     "new",
						FullFilePath: "../h/h1/new",
						Algorithm:    "SHA256",
					},
				}, nil)
			},
			expectedErr: false,
		},
		{
			name:     "not exist path",
			alg:      "SHA256",
			dirFiles: "../h/h1/new",
			expected: []models.HashDataFromDB{{
				ID:           1,
				Hash:         "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
				FileName:     "new",
				FullFilePath: "../h/h1/new",
				Algorithm:    "SHA256",
			}},
			mockBehavior: func(r *mock_ports.MockIHashRepository, ctx context.Context, dirFiles, alg string) {
				r.EXPECT().GetHashSum(ctx, dirFiles, alg).Return([]models.HashDataFromDB{}, errors.New("hash service didn't get data"))
			},
			expectedErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			l := logrus.New()
			ctx := context.Background()
			repo := mock_ports.NewMockIHashRepository(c)
			h, err := hasher.NewHashSum(testCase.alg)
			if err != nil {
				assert.Error(t, err)
			}

			hashService := HashService{
				hashRepository: repo,
				hasher:         h,
				alg:            testCase.alg,
				logger:         l,
			}
			testCase.mockBehavior(repo, ctx, testCase.dirFiles, testCase.alg)

			data, err := hashService.hashRepository.GetHashSum(ctx, testCase.dirFiles, testCase.alg)

			if testCase.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expected, data)
			}
		})
	}
}
