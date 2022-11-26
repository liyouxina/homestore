package file_utils

import "os"

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateFile(path string) error {
	_, err := os.OpenFile(path, os.O_CREATE, 0644)
	return err
}
