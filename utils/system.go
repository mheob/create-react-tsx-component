package utils

import "os"

func GetWd() string {
	path, err := os.Getwd()
	if err != nil {
		return "nothing"
	}
	return path
}
