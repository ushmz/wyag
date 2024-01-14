package main

import (
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func WriteObject(object GitObject, repo *Repository) *string {
	// Serialize the object data
	data := object.Serialize()

	// Add header
	result := fmt.Sprintf("%s %d\x00%s", object.Fmt(), len(data), data)

	// Compute the hash
	sha1 := sha1.New()
	io.WriteString(sha1, result)
	sha := hex.EncodeToString(sha1.Sum(nil))

	if repo != nil {
		// Compute the path
		path := ensureRepositoryFilePath(repo, true, "objects", sha[0:2], sha[2:])

		// Write to the object store
		file, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		writer := zlib.NewWriter(file)
		defer writer.Close()

		writer.Write(data)
	}

	return &sha
}

func pringHashObjectUsage() {
	fmt.Println("usage: hash-object [-t <type>] [-w] <file>")
}
func pringHashObjectHelp() {
	fmt.Println("Compute object ID and optionally creates a blob from a file")
}

func GetObjectHash(fd string, objectType string, repo *Repository) (*string, error) {
	// Read the file content
	file, err := os.Open(fd)
	if err != nil {
		return nil, fmt.Errorf("%w : %s", ErrNoSuchFile, fd)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("%w : %s", ErrFailToReadFile, fd)
	}

	switch objectType {
	case "commit":
		return WriteObject(&Commit{Data: content}, repo), nil
	case "tree":
		return WriteObject(&Tree{Data: content}, repo), nil
	case "tag":
		return WriteObject(&Tag{Data: content}, repo), nil
	case "blob":
		return WriteObject(&Blob{Data: content}, repo), nil
	default:
		return nil, fmt.Errorf("%w : unknown type %s", ErrMalformedObject, objectType)
	}
}

func HashObjectCmd(objectTypeFlag *string, writeFlag *bool, path string) {
	if *writeFlag {
		repo, err := FindRepository(".")
		if err != nil {
			log.Fatalf("%+v\n", err)
		}
		sha, err := GetObjectHash(path, *objectTypeFlag, repo)
		if err != nil {
			log.Fatalf("%+v\n", err)
		}
		fmt.Println(*sha)
	} else {
		sha, err := GetObjectHash(path, *objectTypeFlag, nil)
		if err != nil {
			log.Fatalf("%+v\n", err)
		}
		fmt.Println(*sha)
	}

}
