package main

import (
	"flag"
	"fmt"
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

	utils.RequireFlag(flagYear.LongName, year)
	utils.RequireFlag(flagDay.LongName, day)

	cwd, _ := os.Getwd()
	yearPath := path.Join(cwd, fmt.Sprint(year))
	dayPath := path.Join(yearPath, fmt.Sprintf("%02d", day))

	if yearDirExists := utils.PathExists(yearPath); !yearDirExists {
		utils.CreateFolder("Year", yearPath)
	}
	if dayDirExists := utils.PathExists(dayPath); !dayDirExists {
		utils.CreateFolder("Day", dayPath)
	}

	utils.CreateFile(
		utils.MainFileName,
		dayPath,
		utils.MainFileContent(fmt.Sprint(year), fmt.Sprint(day)),
	)
	utils.CreateFile(utils.InputFileName, dayPath, utils.InputFileContent)
}
