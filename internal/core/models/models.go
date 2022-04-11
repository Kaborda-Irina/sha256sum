package models

type HashSum struct {
	Hash         []byte
	FileName     string
	FullFilePath string
}

type HashSumFromDB struct {
	Id           int
	Hash         string
	FileName     string
	FullFilePath string
}
