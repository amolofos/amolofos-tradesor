package tests

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"reflect"
)

func isContentTheSame(dirFirst, dirSecond string, failFast bool) (result bool, err error) {
	dataFirst := map[string]string{}
	dataSecond := map[string]string{}

	// 1. Get all the files and their hashes.
	errFirst := filepath.Walk(dirFirst, func(file string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			hashFilename, hashContent, errHash := hashFileContent(file)
			if errHash != nil {
				return errHash
			}

			dataFirst[hashFilename] = hashContent
		}
		return err
	})
	if errFirst != nil {
		return false, errFirst
	}

	errSecond := filepath.Walk(dirSecond, func(file string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			hashFilename, hashContent, errHash := hashFileContent(file)
			if errHash != nil {
				return errHash
			}

			dataSecond[hashFilename] = hashContent
		}
		return err
	})
	if errSecond != nil {
		return false, errSecond
	}

	// 2. Compare the number of files.
	resultCount := true
	if len(dataFirst) != len(dataSecond) {
		resultCount = false
		slog.Error(fmt.Sprintf("Different files detected. First directory #%d while second directory #%d.", len(dataFirst), len(dataSecond)))
	}
	if !resultCount && failFast {
		return resultCount, nil
	}

	// 3. Compare the filenames and contents.
	resultContent := reflect.DeepEqual(dataFirst, dataSecond)
	if !resultContent {
		slog.Error("Files are not the same.")
	}

	return resultCount && resultContent, nil
}

func hashFileContent(file string) (hashFilename, hashContent string, err error) {
	hasher := sha256.New()
	hasher.Reset()

	fileName := path.Base(file)
	_, err = hasher.Write([]byte(fileName))
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to hash filename %s.", fileName))
		return
	}

	hashFilename = string(hasher.Sum(nil))
	hasher.Reset()

	f, errOpen := os.Open(file)
	if errOpen != nil {
		slog.Error(fmt.Sprintf("Failed to open file %s.", file))
		err = errOpen
		return
	}
	defer f.Close()

	_, err = io.Copy(hasher, f)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to hash file %s.", file))
		return
	}

	hashContent = string(hasher.Sum(nil))
	return
}
