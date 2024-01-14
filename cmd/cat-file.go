package cmd

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	d "wyag/internal/domain"
	e "wyag/internal/error"
)

func PrintCatFileUsage() {
	fmt.Println("NAME")
	fmt.Println("    wyag cat-file - Provide content or type and size information for repository objects")
	fmt.Println("SYNOPSIS")
	fmt.Println("    wyag cat-file [-t|-s|-p] <object>")
}

func CatFileCmd(objectType string, object string) {
	repo, err := d.FindRepository(".")
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	obj, err := ReadObject(repo, FindObject(repo, object, objectType, true))
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	fmt.Printf("%s\n", obj.Serialize())
}

func FindObject(repo *d.Repository, name string, fmt string, follow bool) string {
	return name
}

// Read object sha from Git repository.
// Return a GirObject whose exact type depends on the object.
func ReadObject(repo *d.Repository, sha string) (d.GitObject, error) {
	path := d.RepositoryFilePath(repo, "objects", sha[0:2], sha[2:])

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%w : %s", e.ErrNoSuchFile, path)
	}

	raw, err := zlib.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("%w : %s", e.ErrFailToReadFile, path)
	}
	defer raw.Close()

	// Read the object type and size
	data, err := io.ReadAll(raw)
	if err != nil {
		return nil, fmt.Errorf("%w : %s", e.ErrFailToReadFile, path)
	}

	index := bytes.IndexAny(data, "\x00")
	if index == -1 {
		return nil, fmt.Errorf("%w : %s", e.ErrMalformedObject, sha)
	}

	metaStr := string(data[:index])
	content := data[index:]

	meta := strings.Split(metaStr, " ")
	objectType := meta[0]
	size, err := strconv.Atoi(meta[1])
	if err != nil {
		return nil, fmt.Errorf("%w %s: bad length", e.ErrMalformedObject, sha)
	}

	// Validate the length
	if size != len(content)-1 {
		return nil, fmt.Errorf("%w %s: bad length", e.ErrMalformedObject, sha)
	}

	// Pick the right constructor
	switch objectType {
	case "commit":
		return &d.Commit{Data: content}, nil
	case "tree":
		return &d.Tree{Data: content}, nil
	case "tag":
		return &d.Tag{Data: content}, nil
	case "blob":
		return &d.Blob{Data: content}, nil
	default:
		return nil, fmt.Errorf("%w %s: unknown type %s", e.ErrMalformedObject, sha, objectType)
	}
}
