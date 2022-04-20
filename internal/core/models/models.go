package models

type HashDataFromDB struct {
	Id           int
	Hash         string
	FileName     string
	FullFilePath string
	Algorithm    string
}

type DeletedHashes struct {
	FileName    string
	OldChecksum string
	FilePath    string
	Algorithm   string
}
