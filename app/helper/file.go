package helper

import (
	"os"
	"strings"
)

func Mkdir(PathStorage string) error {
	arrPath := strings.Split(PathStorage, "/")

	var folderPath = arrPath[0]

	for key, _ := range arrPath {
		path := folderPath

		if key < len(arrPath)-1 {
			folderPath = folderPath + "/" + arrPath[key+1]
		}

		if _, err := os.Stat(path); !os.IsNotExist(err) {
			continue
		}

		errMkd := os.Mkdir(path, 0755)

		if errMkd != nil {
			return errMkd
		}
	}

	return nil
}
