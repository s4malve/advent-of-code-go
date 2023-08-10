package utils

import (
	"bufio"
	"flag"
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
	Part    uint
}

func PrintAdventResult(adResult AdventResult) {
	fmt.Printf("Advent of Code %s\n", adResult.Year)
	fmt.Printf("Day %s Solution Part %d\n", adResult.Day, adResult.Part)
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

func CreateFile(name, filePath, content string) {
	mainFilePath := path.Join(filePath, name)
	fmt.Printf("opening or creating %s file...\n", name)
	f, err := os.OpenFile(mainFilePath, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("opening or creating %s file: %v\n", name, err)
	}
	fmt.Printf("file %s created\n", name)

	defer f.Close()
	fmt.Printf("writing %s file...\n", name)
	_, err = f.Write([]byte(content))

	if err != nil {
		log.Fatalf("writing in %s file: %v\n", name, err)
	}
	fmt.Printf("%s file created\n", name)
}

func CreateFolder(name, path string) {
	fmt.Printf("Creating %s folder...\n", name)
	err := os.Mkdir(path, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s folder created\n", name)
}

func GetInputFileScanner(year, day string) (file *os.File, fileScanner *bufio.Scanner) {
	f, err := os.Open(GetFullInputPath(year, day))
	Fatal(err)

	fileScanner = bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	return f, fileScanner
}

func PathExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

func RequireFlag(name string, value uint) {
	if value == 0 {
		log.Fatalf("flag -%s or -%c is required\n", name, name[0])
	}
}

func ParsePartFlag() (part uint) {
	const (
		name  = "part"
		usage = "part of the advent, can be 1 or 2"
	)

	flag.UintVar(&part, name, 1, usage)
	flag.Parse()
	RequireFlag(name, part)

	if part != 2 && part != 1 {
		log.Fatal("The part must be 1 or 2")
	}

	return part
}
