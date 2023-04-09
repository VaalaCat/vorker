package utils

import "os"

func WriteFile(path string, content string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
