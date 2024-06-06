// Section 1: Language Fundamentals

// Interfaces
package section1

import (
	"fmt"
	"os"
)

type Logger interface {
	Log(message string)
}

type FileLogger struct {
	File *os.File
}

func (f *FileLogger) Log(message string) {
	f.File.WriteString(message + "\n")
}

type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	fmt.Println(message)
}
