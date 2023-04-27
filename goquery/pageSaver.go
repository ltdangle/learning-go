package main

import (
	"net/url"
	"path/filepath"
)

type PageSaver struct {
	fileSystem FileSystem
	mountPath  string
}

func NewPageSaver(fileSystem FileSystem, mountPath string) *PageSaver {
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

	filePath := s.mountPath + string(filepath.Separator) + host + path

	return filePath, nil
}
