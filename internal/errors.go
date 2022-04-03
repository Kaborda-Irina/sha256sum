package internal

import "errors"

var (
	ErrorFilePath = errors.New("error: not exist file path")
	ErrorHash     = errors.New("error: hash getting from file")
)
