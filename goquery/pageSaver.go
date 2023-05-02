package main

import (
	"io/fs"
	"log"
	"net/url"
	"path/filepath"
)

type PageSaver struct {
	fileSystem FileSystemInterface
	mountPath  string
}

func NewPageSaver(fileSystem FileSystemInterface, mountPath string) *PageSaver {
	// Strip dangling slash if present
	if mountPath[len(mountPath)-1] == '/' {
		mountPath = mountPath[0 : len(mountPath)-1]
	}
	return &PageSaver{fileSystem: fileSystem, mountPath: mountPath}
}

func (s *PageSaver) SavePage(urlStr string, html string) (string, error) {
	urlObj, _ := url.Parse(urlStr)
	path := urlObj.Path
	host := urlObj.Host

	if path == "" || path == "/" {
		path = "/index"
	}
	if path[len(path)-1] == '/' {
		path += "index"
	}

	fullPath := s.mountPath + string(filepath.Separator) + host + path

	err := s.fileSystem.Save(fullPath, html)
	//
	if err != nil {
		// Check if the error is of type fs.PathError
		if pathErr, ok := err.(*fs.PathError); ok {
			// Cannot create directory, meaning that file with the same path exist.
			if pathErr.Op == "mkdir" {
				dir, file := filepath.Split(fullPath)
				offendingFile := dir[0 : len(dir)-1]

				// Copy offending file to new location
				tmpFile := s.mountPath + string(filepath.Separator) + "copy" + string(filepath.Separator) + file
				s.fileSystem.Copy(offendingFile, tmpFile)

				// Delete offending file.
				err := s.fileSystem.Delete(offendingFile)
				if err != nil {
					log.Fatal(err)
				}

				// Write our fullPath.
				err = s.fileSystem.Save(fullPath, html)
				if err != nil {
					log.Fatal(err)
				}

				// Copy file as index to current directory.
				s.fileSystem.Copy(tmpFile, offendingFile+"/index")
			}
		}

	}

	return fullPath, nil
}
