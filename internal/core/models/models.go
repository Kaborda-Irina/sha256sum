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

type ChangedHashes struct {
	FileName    string
	OldChecksum string
	NewChecksum string
	FilePath    string
}
type DeletedHashes struct {
	FileName    string
	OldChecksum string
	FilePath    string
	Algorithm   string
}
type AddedHashes struct {
	FileName    string
	NewChecksum string
	FilePath    string
	Algorithm   string
}
