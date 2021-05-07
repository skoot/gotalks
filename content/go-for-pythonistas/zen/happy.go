package example

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("file open error: %w", err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("file read error: %w", err)
	}

	return string(bytes), nil
}
