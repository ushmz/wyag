package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

func InitCmd(path string) {
	if _, err := createRepository(path); err != nil {
		log.Fatalf("%+v\n", err)
	}
}

// createRepository creates a new repositry at `path`.
func createRepository(path string) (*Repository, error) {
	repo, err := NewRepository(path, true)
	if err != nil {
		return nil, err
	}

	// First, we make sure the path either doesn't exist
	// or is an empty directory.

	if info, err := os.Stat(repo.WorkTree); err != nil {
		os.Mkdir(repo.WorkTree, os.ModePerm)
	} else {
		if !info.IsDir() {
			return nil, fmt.Errorf("%w : %s", ErrNotDirectory, path)
		}

		files, err := ioutil.ReadDir(repo.WorkTree)
		if err != nil {
			return nil, fmt.Errorf("%w : %s", ErrFailToReadDirectory, path)
		}

		if len(files) != 0 {
			return nil, fmt.Errorf("%w : %s", ErrNotEmptyDirectory, path)
		}
	}

	if _, err := ensureRepositoryDirPath(repo, true, "branches"); err != nil {
		return nil, fmt.Errorf("%w : %s", ErrFailToCreateDirectory, repo.GitDir+"/branches")
	}

	if _, err := ensureRepositoryDirPath(repo, true, "objects"); err != nil {
		return nil, fmt.Errorf("%w : %s", ErrFailToCreateDirectory, repo.GitDir+"/objects")
	}
	if _, err := ensureRepositoryDirPath(repo, true, "refs", "tags"); err != nil {
		return nil, fmt.Errorf("%w : %s", ErrFailToCreateDirectory, repo.GitDir+"/refs/tags")
	}
	if _, err := ensureRepositoryDirPath(repo, true, "refs", "heads"); err != nil {
		return nil, fmt.Errorf("%w : %s", ErrFailToCreateDirectory, repo.GitDir+"/refs/heads")
	}

	// .git/description
	desc := "Unnamed repository; edit this file 'description' to name the repository.\n"
	if err := ensureFileAndContent(repo, false, desc, "description"); err != nil {
		return nil, err
	}

	// .git/HEAD
	head := "ref: refs/heads/master\n"
	if err := ensureFileAndContent(repo, false, head, "HEAD"); err != nil {
		return nil, err
	}

	dc, err := defaultConfig()
	if err != nil {
		return nil, fmt.Errorf("%w : %s", ErrFailToCreateFile, "Cannot get default configs")
	}
	cfg := ensureRepositoryFilePath(repo, false, "config")
	dc.SaveTo(cfg)

	return repo, nil
}

func ensureFileAndContent(repo *Repository, mkdir bool, content string, path ...string) error {
	fp := ensureRepositoryFilePath(repo, mkdir, path...)
	f, err := os.Open(fp)
	if err != nil {
		f, err = os.Create(fp)
		if err != nil {
			return fmt.Errorf("%w : %s", ErrFailToReadFile, fp)
		}
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	if _, err := f.Write([]byte(content)); err != nil {
		return fmt.Errorf("%w : %s", ErrFailToCreateFile, fp)
	}
	return nil
}

func defaultConfig() (*ini.File, error) {
	ret := ini.Empty()

	sec, err := ret.NewSection("core")
	if err != nil {
		return nil, err
	}

	sec.NewKey("repositoryformatversion", "0")
	sec.NewKey("filemode", "false")
	sec.NewKey("bare", "false")

	return ret, nil
}
