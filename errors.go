package main

import "errors"

var (
	ErrCannotGetFileStat     = errors.New("Cannot get file stat")
	ErrNotGitRepository      = errors.New("Not a Get repository")
	ErrNotDirectory          = errors.New("Not a directory")
	ErrNotFile               = errors.New("Not a file")
	ErrNoSuchDirectory       = errors.New("No such directory")
	ErrNoSuchFile            = errors.New("No such file")
	ErrNotEmptyDirectory     = errors.New("Not an empty directory")
	ErrFailToReadDirectory   = errors.New("Failed to read file")
	ErrFailToReadFile        = errors.New("Failed to read file")
	ErrFailToCreateDirectory = errors.New("Failed to create directory")
	ErrFailToCreateFile      = errors.New("Failed to create file")
	ErrMalformedObject       = errors.New("Malformed object")

	ErrUnsupportedRepositoryFormatVersion = errors.New("Unsupported repositoryformatversion")
)
