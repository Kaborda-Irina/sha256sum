package models

type HashData struct {
	Hash         []byte
	FileName     string
	FullFilePath string
	Algorithm    string
}

type HashDataFromDB struct {
	Id           int
	Hash         string
	FileName     string
	FullFilePath string
	Algorithm    string
}
