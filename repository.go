package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

type Repository struct {
	WorkTree string
	GitDir   string
	Config   *ini.File
}

func NewRepository(path string, force bool) (*Repository, error) {
	repo := Repository{
		WorkTree: path,
		GitDir:   filepath.Join(path, ".git"),
	}

	info, err := os.Stat(repo.GitDir)
	if err != nil {
		return nil, fmt.Errorf("%w : %s", ErrCannotGetFileStat, repo.GitDir)
	}

	if !(force || !info.IsDir()) {
		return nil, fmt.Errorf("%w : %s", ErrNotGitRepository, path)
	}

	// Read configuration file in .git/config
	cf := mustRepositoryFilePath(repo, false, "config")

	if cf == "" {
		return nil, fmt.Errorf("%w : %s", ErrNoSuchDirectory, repo.GitDir)
	}

	if info, err := os.Stat(cf); err != nil {
		return nil, fmt.Errorf("%w : %s", ErrCannotGetFileStat, cf)
	} else {
		if info.IsDir() {
			if f, err := ini.Load(cf); err != nil {
				return nil, fmt.Errorf("%w : %s", ErrFailToReadFile, cf)
			} else {
				repo.Config = f
			}
		} else if !force {
			return nil, fmt.Errorf("%w : %s", ErrNoSuchFile, cf)
		}

		if !force {
			versKey := repo.Config.Section("core").Key("repositoryformatversion")
			if vers, err := versKey.Int(); err != nil {
				if vers != 0 {
					return nil, fmt.Errorf("%w : %d", ErrUnsupportedRepositoryFormatVersion, vers)
				}
			} else {
				return nil, fmt.Errorf("%w : %d", ErrUnsupportedRepositoryFormatVersion, vers)
			}
		}
	}

	return &repo, nil
}

// repositoryFilePath : Compute path under repository's gitdir.
func repositoryFilePath(repo Repository, path ...string) string {
	args := append([]string{repo.GitDir}, path...)
	return filepath.Join(args...)
}

// mustRepositoryFilePath : Same as repositoryFilePath, but create dirname(*path) if absent.
// For example, repo_file(r, "refs", "remotes", "origin", "HEAD")
// will create .git/refs/remotes/origin.
// repo: Repository
// mkdir: If true, make parent directory (like `p` option of mkdir command)
// path: string array of path
func mustRepositoryFilePath(repo Repository, mkdir bool, path ...string) string {
	if dir, err := repositoryDirPath(repo, mkdir, path[:len(path)-1]...); err == nil {
		return repositoryFilePath(repo, path...)
	} else {
		return dir
	}
}

// repositoryDirPath : Same as repositoryFilePath, but mkdir *path if absent if mkdir == true.
func repositoryDirPath(repo Repository, mkdir bool, path ...string) (string, error) {
	p := repositoryFilePath(repo, path...)

	if info, err := os.Stat(p); err == nil {
		if info.IsDir() {
			return p, nil
		} else {
			return "", fmt.Errorf("%w : %s", ErrNotDirectory, p)
		}
	}

	if mkdir {
		if err := os.Mkdir(p, fs.ModePerm); err != nil {
			return "", fmt.Errorf("%w : %s", err, p)
		}
		return p, nil
	} else {
		return "", fmt.Errorf("%w : %s", ErrNoSuchDirectory, p)
	}
}
