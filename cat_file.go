package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func printCatFileUsage() {
	fmt.Println("usage: cat-file [-t|-s|-p] <object>")
}

func FindObject(repo *Repository, name string, fmt string, follow bool) string {
	return name
}

// Read object sha from Git repository.
// Return a GirObject whose exact type depends on the object.
func ReadObject(repo *Repository, sha string) (GitObject, error) {
	path := repositoryFilePath(repo, "objects", sha[0:2], sha[2:])

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%w : %s", ErrNoSuchFile, path)
	}

	raw, err := zlib.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("%w : %s", ErrFailToReadFile, path)
	}
	defer raw.Close()

	// Read the object type and size
	data, err := io.ReadAll(raw)
	if err != nil {
		return nil, fmt.Errorf("%w : %s", ErrFailToReadFile, path)
	}

	index := bytes.IndexAny(data, "\x00")
	if index == -1 {
		return nil, fmt.Errorf("%w : %s", ErrMalformedObject, sha)
	}

	metaStr := string(data[:index])
	content := data[index:]

	meta := strings.Split(metaStr, " ")
	objectType := meta[0]
	size, err := strconv.Atoi(meta[1])
	if err != nil {
		return nil, fmt.Errorf("%w %s: bad length", ErrMalformedObject, sha)
	}

	// Validate the length
	if size != len(content)-1 {
		return nil, fmt.Errorf("%w %s: bad length", ErrMalformedObject, sha)
	}

	// Pick the right constructor
	switch objectType {
	case "commit":
		return &Commit{Data: content}, nil
	case "tree":
		return &Tree{Data: content}, nil
	case "tag":
		return &Tag{Data: content}, nil
	case "blob":
		return &Blob{Data: content}, nil
	default:
		return nil, fmt.Errorf("%w %s: unknown type %s", ErrMalformedObject, sha, objectType)
	}
}

func CatFileCmd(objectType string, object string) {
	repo, err := FindRepository(".")
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	obj, err := ReadObject(repo, FindObject(repo, object, objectType, true))
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	fmt.Printf("%s\n", obj.Serialize())
}
