package main

// FileSystem abstraction interface.
type FileSystem interface {
	Save(path string, content string) error
	Exists(path string)
	IsDir(path string)
	IsFile(path string)
}
