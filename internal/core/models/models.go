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

type DeletedHashes struct {
	FileName    string
	OldChecksum string
	FilePath    string
	Algorithm   string
}
