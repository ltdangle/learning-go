package logger

import (
	"fmt"
	"log"
	"os"
)

type ILogger interface {
	Log(msg interface{})
}

type logger struct {
	filename string
}

func NewLogger(filename string) ILogger {
	l := &logger{filename: filename}
	l.emptyLog()

	return l
}

func (self logger) Log(msg interface{}) {
	file, err := os.OpenFile(self.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := fmt.Fprintln(file, msg); err != nil {
		log.Fatal(err)
	}
}

func (self logger) emptyLog() error {
	// Open the file for writing, truncating its contents to 0.
	file, err := os.OpenFile(self.filename, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write an empty byte slice to the file to clear its contents.
	if _, err := file.Write([]byte{}); err != nil {
		return err
	}

	return nil
}
