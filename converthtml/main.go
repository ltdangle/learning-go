package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <InputDirectory> <OutputDirectory>")
		os.Exit(1)
	}

	inputDir := os.Args[1]
	outputDir := os.Args[2]

	err := convertHTMLFilesToText(inputDir, outputDir)
	if err != nil {
		fmt.Println("Error converting HTML files:", err)
	}
}

func convertHTMLFilesToText(inputDir string, outputDir string) error {
	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		return err
	}

	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".html") {
			inputPath := filepath.Join(inputDir, file.Name())
			outputPath := filepath.Join(outputDir, strings.TrimSuffix(file.Name(), ".html")+".txt")

			cmd := exec.Command("w3m", "-dump", "-cols", "80", "-T", "text/html", inputPath)
			output, err := cmd.Output()
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(outputPath, output, 0644)
			if err != nil {
				return err
			}

			fmt.Printf("Converted %s to %s\n", inputPath, outputPath)
		}
	}

	return nil
}
