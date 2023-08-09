package utils

import (
	"log"
	"os"
	"path"
)

func ReadInputFile(year, day string) string {
	cwd, _ := os.Getwd()
	filepath := path.Join(cwd, year, day, InputFileName)
	bytes, err := os.ReadFile(filepath)

	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}
