package main

import (
	"testing"
)

func TestNewPageSaver(t *testing.T) {
	// Define a table of test cases.
	testCases := []struct {
		name string
		url  string
		file string
	}{
		{"Homepage", "https://domain.com", "/tmp/domain.com/index"},
		{"Homepage with slash", "https://domain.com/", "/tmp/domain.com/index"},
		{"1st level page", "https://domain.com/page1", "/tmp/domain.com/page1"},
		{"1st level page with slash", "https://domain.com/page1/", "/tmp/domain.com/page1/index"},
		{"2nd level page", "https://domain.com/page1/page2", "/tmp/domain.com/page1/page2"},
		{"2nd level page with slash", "https://domain.com/page1/page2/", "/tmp/domain.com/page1/page2/index"},
	}

	saver := NewPageSaver(&FileSystemStub{}, "/tmp")
	// Iterate through the test cases.
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := saver.SavePage(tt.url, "dummy content")
			if result != tt.file {
				t.Error(tt.name + "\n" + tt.url + " => " + tt.file + ", got " + result)
			}
		})
	}
}

type FileSystemStub struct{}

func (FileSystemStub) Save(path string, content string) error {
	return nil
}

func (FileSystemStub) Exists(path string) {
	return
}

func (FileSystemStub) IsDir(path string) {
	return
}

func (FileSystemStub) IsFile(path string) {
	return
}
