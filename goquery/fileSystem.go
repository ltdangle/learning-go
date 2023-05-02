package main

import (
	"io"
	"os"
	"path/filepath"
)

// FileSystemInterface abstraction interface.
type FileSystemInterface interface {
	Save(path string, content string) error
	IsFile(path string) bool
	// Delete works both files and directories.
	Delete(path string) error
	Move(sourcePath string, destinationPath string) error
	Copy(sourcePath string, destinationPath string) error
}

type FileSystem struct {
}

func NewFileSystem() *FileSystem {
	return &FileSystem{}
}

func (fs *FileSystem) Save(filePath string, content string) error {
	dir := filepath.Dir(filePath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	// Write the html to the file.
	err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (fs *FileSystem) IsFile(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func (fs *FileSystem) Delete(path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		err = os.RemoveAll(path)
		if err != nil {
			return err
		}
	} else {
		err = os.Remove(path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (fs *FileSystem) Move(sourcePath string, destinationPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	err = os.Remove(sourcePath)
	if err != nil {
		return err
	}

	return nil
}

func (fs *FileSystem) Copy(sourcePath string, destinationPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}
