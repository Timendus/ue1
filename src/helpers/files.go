package helpers

import (
	"fmt"
	"os"
)

func LoadTextFile(filename string) (string, error) {
	if _, err := os.Stat(filename); err != nil {
		return "", fmt.Errorf("requested file '%s' not found", filename)
	}
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error reading file '%s': %s", filename, err.Error())
	}
	return string(file), nil
}

func LoadBinaryFile(filename string) ([]byte, error) {
	if _, err := os.Stat(filename); err != nil {
		return nil, fmt.Errorf("requested file '%s' not found", filename)
	}
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file '%s': %s", filename, err.Error())
	}
	return file, nil
}
