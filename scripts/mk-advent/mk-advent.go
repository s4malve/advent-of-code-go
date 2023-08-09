package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/s4malve/advent-of-code-go/utils"
)

type Flag struct {
	ShortName string
	LongName  string
	Usage     string
}

func main() {
	var (
		day     uint
		flagDay = Flag{
			ShortName: "d",
			LongName:  "day",
			Usage:     "day of the advent",
		}
	)
	var (
		year     uint
		flagYear = Flag{
			ShortName: "y",
			LongName:  "year",
			Usage:     "year of the advent",
		}
	)

	flag.UintVar(&day, flagDay.LongName, 0, flagDay.Usage)
	flag.UintVar(&day, flagDay.ShortName, 0, flagDay.Usage)
	flag.UintVar(&year, flagYear.LongName, 0, flagYear.Usage)
	flag.UintVar(&year, flagYear.ShortName, 0, flagYear.Usage)
	flag.Parse()

	RequireFlag(flagYear.LongName, year)
	RequireFlag(flagDay.LongName, day)

	cwd, _ := os.Getwd()
	yearPath := path.Join(cwd, fmt.Sprint(year))
	dayPath := path.Join(yearPath, fmt.Sprintf("%02d", day))

	if yearDirExists := PathExists(yearPath); !yearDirExists {
		CreateFolder("Year", yearPath)
	}
	if dayDirExists := PathExists(dayPath); !dayDirExists {
		CreateFolder("Day", dayPath)
	}

	CreateFile(utils.MainFileName, dayPath, utils.MainFileContent)
	CreateFile(utils.InputFileName, dayPath, utils.InputFileContent)
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

func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsExist(err) {
		return true
	}

	return false
}

func RequireFlag(name string, value uint) {
	if value == 0 {
		log.Fatalf("flag -%s or -%c is required\n", name, name[0])
	}
}
