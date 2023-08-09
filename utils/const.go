package utils

import "fmt"

var (
	MainFileName    = "main.go"
	MainFileContent = func(year, day string) string {
		return fmt.Sprintf(`package main

const (
	year = "%s"
	day  = "%s"
)

func main() {

}
`, year, day)
	}
)

const (
	InputFileName    = "input.txt"
	InputFileContent = ""
)
