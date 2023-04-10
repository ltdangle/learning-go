package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// To use the program, compile it and run it with the -dir and -unit flags, like this:
// $ go build dirsize.go
// $ ./dirsize -dir /path/to/directory -unit MB
func getDirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size, err
}

func formatSize(size int64, unit string) string {
	switch unit {
	case "GB":
		return fmt.Sprintf("%.2f GB", float64(size)/float64(1<<30))
	case "MB":
		return fmt.Sprintf("%.2f MB", float64(size)/float64(1<<20))
	case "KB":
		return fmt.Sprintf("%.2f KB", float64(size)/float64(1<<10))
	default:
		return fmt.Sprintf("%d bytes", size)
	}
}

func main() {
	dirPath := flag.String("dir", ".", "The directory to calculate the size of")
	unit := flag.String("unit", "bytes", "The unit to display the size in (bytes, KB, MB, GB)")
	flag.Parse()

	size, err := getDirSize(*dirPath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	formattedSize := formatSize(size, *unit)
	fmt.Printf("Size of directory %s: %s\n", *dirPath, formattedSize)
}
