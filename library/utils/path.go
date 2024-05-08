package utils

import "os"

func IsFile(filepath string) bool {
	statInfo, err := os.Stat(filepath)
	if err != nil {
		return false
	}
	if statInfo.IsDir() {
		return false
	}
	return true
}

func IsDir(filepath string) bool {
	statInfo, err := os.Stat(filepath)
	if err != nil {
		return false
	}
	if statInfo.IsDir() {
		return true
	}
	return false
}
