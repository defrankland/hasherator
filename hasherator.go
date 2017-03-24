package Hasherator

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type AssetsDir struct {
	Map map[string]string
}

var noHashDirList *[]string

func (a *AssetsDir) Run(sourcePath, workingPath string, noHashDirs []string) error {

	noHashDirList = &noHashDirs

	a.Map = map[string]string{}

	err := os.RemoveAll(workingPath)
	if err != nil {
		return fmt.Errorf("failed to remove working directory prior to copy: " + err.Error())
	}

	err = a.recursiveHashAndCopy(sourcePath, workingPath)
	if err != nil {
		return err
	}

	return nil
}

func (a *AssetsDir) recursiveHashAndCopy(dirPath, runtimePath string) error {
	var err error

	assetDirs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("error reading directory: %s", err)
	}

	for _, fileEntry := range assetDirs {

		entryName := fileEntry.Name()

		if fileEntry.IsDir() {
			newPath := fmt.Sprintf("%s/%s", runtimePath, entryName)

			err := os.MkdirAll(newPath, 0777)
			if err != nil {
				panic("failed to make directory: " + err.Error())
			}

			err = a.recursiveHashAndCopy(fmt.Sprintf("%s%s/", dirPath, fileEntry.Name()), newPath)
			if err != nil {
				return err
			}

		} else {

			var fileExtension string
			var dot string
			if strings.Contains(entryName, ".") {
				fileExtension = entryName[strings.LastIndex(entryName, ".")+1:]
				entryName = entryName[:strings.LastIndex(entryName, ".")]
				dot = "."
			}

			file, err := ioutil.ReadFile(fmt.Sprintf("%s%s", dirPath, fileEntry.Name()))
			if err != nil {
				return fmt.Errorf("failed to read file: " + err.Error())
			}

			var hash string
			var noHash bool
			dir := strings.Split(runtimePath, "/")
			for _, noDir := range *noHashDirList {
				if noDir == dir[len(dir)-1] {
					noHash = true
				}
			}

			if !noHash {
				h := md5.Sum(file)
				hash = fmt.Sprintf("-%x", string(h[:16]))
			}

			err = copyFile(fmt.Sprintf("%s%s", dirPath, fileEntry.Name()),
				fmt.Sprintf("%s/%s%s%s%s", runtimePath, entryName, hash, dot, fileExtension))
			if err != nil {
				return fmt.Errorf("failed to return file: " + err.Error())
			}

			a.Map[fileEntry.Name()] = fmt.Sprintf("%s%s%s%s", entryName, hash, dot, fileExtension)
		}
	}
	return nil
}

func copyFile(src, dst string) error {

	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)

	return err
}
