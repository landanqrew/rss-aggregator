package utils

import (
	"fmt"
	"os"
	"strings"
)

func FileExists(path string) bool {
	objInfo, err := os.Stat(path)
	if err != nil {
		return !os.IsNotExist(err)
	}
	if (objInfo.IsDir()) {
		return false
	}
	return true
}

func DirectoryExists(path string) bool {
	objInfo, err := os.Stat(path)
	if err != nil {
		return !os.IsNotExist(err)
	}
	if (objInfo.IsDir()) {
		return true
	}
	return false
}


func CreateFile(path string, data []byte) error {
	directoryParts := strings.Split(path, "/")
	directoryParts = directoryParts[:len(directoryParts)-1]
	err := os.MkdirAll(strings.Join(directoryParts, "/"), 0755)
	if err != nil {
		return fmt.Errorf("error creating directory %s: %w", strings.Join(directoryParts, "/"), err)
	}
	err = os.WriteFile(path, data, 0664)
	if err != nil {
		return fmt.Errorf("error writing file %s: %w", path, err)
	}
	return nil
}

func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", path, err)
	}
	return data, nil
}

func DeleteFile(path string) error {
	return os.Remove(path)
}