package utils

import (
	"fmt"
	"log"
	"math"
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

func MaxNumber(numbers ...uint) uint {
	maxNumber := numbers[0]

	for _, num := range numbers {
		maxNumber = uint(math.Max(float64(maxNumber), float64(num)))
	}

	return maxNumber
}

type AdventResult struct {
	Year    string
	Day     string
	Message string
}

func PrintAdventResult(adResult AdventResult) {
	fmt.Printf("Advent of Code %s\n", adResult.Year)
	fmt.Printf("Day %s Solution\n", adResult.Day)
	fmt.Println("Answer:", adResult.Message)
}

func GetFullInputPath(year, day string) string {
	cwd, _ := os.Getwd()
	return path.Join(cwd, year, day, InputFileName)
}

func Fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}

}
