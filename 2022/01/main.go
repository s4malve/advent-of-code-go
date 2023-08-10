package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/s4malve/advent-of-code-go/utils"
)

const (
	year = "2022"
	day  = "01"
)

func main() {
	part := utils.ParsePartFlag()

	if part == 1 {
		partOne()
	} else {
		partTwo()
	}
}

func GetElvesCalories() []uint {
	var totalElvesCalories []uint
	f, err := os.Open(utils.GetFullInputPath(year, day))
	utils.Fatal(err)
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	totalElveCalories := 0
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		isElveCalory := txt != ""

		if isElveCalory {
			elveCalory, err := strconv.Atoi(txt)
			utils.Fatal(err)
			totalElveCalories += elveCalory
		} else {
			totalElvesCalories = append(totalElvesCalories, uint(totalElveCalories))
			totalElveCalories = 0
		}
	}

	return totalElvesCalories
}

func partOne() {
	totalElvesCalories := GetElvesCalories()

	utils.PrintAdventResult(utils.AdventResult{
		Year:    year,
		Day:     day,
		Part:    1,
		Message: fmt.Sprintf("The highest number of calories is %d", utils.MaxNumber(totalElvesCalories...)),
	})
}

func partTwo() {
	totalElvesCalories := GetElvesCalories()

	sort.Slice(
		totalElvesCalories,
		func(i, j int) bool { return totalElvesCalories[i] < totalElvesCalories[j] },
	)

	highestThreeCalories := totalElvesCalories[len(totalElvesCalories)-3:]

	utils.PrintAdventResult(utils.AdventResult{
		Year:    year,
		Day:     day,
		Part:    2,
		Message: fmt.Sprintf("The sum of the three highest calories is %d", utils.MaxNumber(highestThreeCalories...)),
	})
}
