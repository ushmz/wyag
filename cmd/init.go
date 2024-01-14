package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	d "wyag/internal/domain"
	e "wyag/internal/error"

	"gopkg.in/ini.v1"
)

func InitCmd(path string) {
	if _, err := createRepository(path); err != nil {
		log.Fatalf("%+v\n", err)
	}
}

// createRepository creates a new repositry at `path`.
func createRepository(path string) (*d.Repository, error) {
	repo, err := d.NewRepository(path, true)
	if err != nil {
		return nil, err
	}

	if info, err := os.Stat(repo.WorkTree); err != nil {
		os.Mkdir(repo.WorkTree, os.ModePerm)
	} else {
		if !info.IsDir() {
			return nil, fmt.Errorf("%w : %s", e.ErrNotDirectory, path)
		}

		files, err := ioutil.ReadDir(repo.WorkTree)
		if err != nil {
			return nil, fmt.Errorf("%w : %s", e.ErrFailToReadDirectory, path)
		}

		if len(files) != 0 {
			return nil, fmt.Errorf("%w : %s", e.ErrNotEmptyDirectory, path)
		}
	}

	if _, err := d.EnsureRepositoryDirPath(repo, true, "branches"); err != nil {
		return nil, fmt.Errorf("%w : %s", e.ErrFailToCreateDirectory, repo.GitDir+"/branches")
	}

	if _, err := d.EnsureRepositoryDirPath(repo, true, "objects"); err != nil {
		return nil, fmt.Errorf("%w : %s", e.ErrFailToCreateDirectory, repo.GitDir+"/objects")
	}
	if _, err := d.EnsureRepositoryDirPath(repo, true, "refs", "tags"); err != nil {
		return nil, fmt.Errorf("%w : %s", e.ErrFailToCreateDirectory, repo.GitDir+"/refs/tags")
	}
	if _, err := d.EnsureRepositoryDirPath(repo, true, "refs", "heads"); err != nil {
		return nil, fmt.Errorf("%w : %s", e.ErrFailToCreateDirectory, repo.GitDir+"/refs/heads")
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
		return nil, fmt.Errorf("%w : %s", e.ErrFailToCreateFile, "Cannot get default configs")
	}
	cfg := d.EnsureRepositoryFilePath(repo, false, "config")
	dc.SaveTo(cfg)

	return repo, nil
}

func ensureFileAndContent(repo *d.Repository, mkdir bool, content string, path ...string) error {
	fp := d.EnsureRepositoryFilePath(repo, mkdir, path...)
	f, err := os.Open(fp)
	if err != nil {
		f, err = os.Create(fp)
		if err != nil {
			return fmt.Errorf("%w : %s", e.ErrFailToReadFile, fp)
		}
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	if _, err := f.Write([]byte(content)); err != nil {
		return fmt.Errorf("%w : %s", e.ErrFailToCreateFile, fp)
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
