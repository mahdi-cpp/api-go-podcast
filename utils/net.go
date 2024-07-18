package utils

import (
	"io"
	"os"
	"strings"
)

func ReadHTMLFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the contents of the file
	var content strings.Builder
	_, err = io.Copy(&content, file)
	if err != nil {
		return "", err
	}

	return content.String(), nil
}
