package main

import "errors"

var (
	ErrCannotGetFileStat                  = errors.New("Cannot get file stat")
	ErrNotGitRepository                   = errors.New("Not a Get repository")
	ErrNotDirectory                       = errors.New("Not a directory")
	ErrNotFile                            = errors.New("Not a file")
	ErrNoSuchDirectory                    = errors.New("No such directory")
	ErrNoSuchFile                         = errors.New("No such file")
	ErrFailToReadFile                     = errors.New("Failed to read file")
	ErrUnsupportedRepositoryFormatVersion = errors.New("Unsupported repositoryformatversion")
)
