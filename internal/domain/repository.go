package domain

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	e "wyag/internal/error"

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

	if !force {
		info, err := os.Stat(repo.GitDir)
		if err != nil {
			return nil, fmt.Errorf("%w : %s", e.ErrCannotGetFileStat, repo.GitDir)
		}

		if !info.IsDir() {
			return nil, fmt.Errorf("%w : %s", e.ErrNotGitRepository, path)
		}
	}

	// Read configuration file in .git/config
	cf := EnsureRepositoryFilePath(&repo, false, "config")

	if cf != "" {
		if _, err := os.Stat(cf); err == nil {
			if f, err := ini.Load(cf); err != nil {
				return nil, fmt.Errorf("%w : %s", e.ErrFailToReadFile, cf)
			} else {
				repo.Config = f
			}
		}
	} else if !force {
		return nil, fmt.Errorf("%w : %s", e.ErrNoSuchFile, cf)
	}

	if !force {
		versKey := repo.Config.Section("core").Key("repositoryformatversion")
		if vers, err := versKey.Int(); err != nil {
			return nil, fmt.Errorf("%w : %s", e.ErrUnsupportedRepositoryFormatVersion, "Cannot read core.repositoryformatversion")
		} else {
			if vers != 0 {
				return nil, fmt.Errorf("%w : %d", e.ErrUnsupportedRepositoryFormatVersion, vers)
			}
		}
	}
	return &repo, nil
}

// RepositoryFilePath : Compute path under repository's gitdir.
func RepositoryFilePath(repo *Repository, path ...string) string {
	args := append([]string{repo.GitDir}, path...)
	return filepath.Join(args...)
}

// EnsureRepositoryFilePath : Same as repositoryFilePath, but create dirname(*path) if absent.
// For example, repo_file(r, "refs", "remotes", "origin", "HEAD")
// will create .git/refs/remotes/origin.
// repo: Repository
// mkdir: If true, make parent directory (like `p` option of mkdir command)
// path: string array of path
func EnsureRepositoryFilePath(repo *Repository, mkdir bool, path ...string) string {
	if dir, err := EnsureRepositoryDirPath(repo, mkdir, path[:len(path)-1]...); err == nil {
		return RepositoryFilePath(repo, path...)
	} else {
		return dir
	}
}

// EnsureRepositoryDirPath : Same as repositoryFilePath, but mkdir *path if absent if mkdir == true.
func EnsureRepositoryDirPath(repo *Repository, mkdir bool, path ...string) (string, error) {
	p := RepositoryFilePath(repo, path...)

	if info, err := os.Stat(p); err == nil {
		if info.IsDir() {
			return p, nil
		} else {
			return "", fmt.Errorf("%w : %s", e.ErrNotDirectory, p)
		}
	}

	if mkdir {
		if err := os.MkdirAll(p, fs.ModePerm); err != nil {
			return "", fmt.Errorf("%w : %s", err, p)
		}
		return p, nil
	} else {
		return "", nil
	}
}

func FindRepository(path string) (*Repository, error) {
	p, err := filepath.EvalSymlinks(path)
	if err != nil {
		return nil, err
	}

	info, err := os.Stat(filepath.Join(p, ".git"))
	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		return NewRepository(p, false)
	}

	parent := filepath.Join(p, "..")
	if p == parent {
		return nil, errors.New("No git repository")
	}

	return FindRepository(parent)
}
