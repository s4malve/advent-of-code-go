package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/s4malve/advent-of-code-go/utils"
)

const (
	year = "2022"
	day  = "01"
)

func main() {
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

	utils.PrintAdventResult(utils.AdventResult{
		Year:    year,
		Day:     day,
		Message: fmt.Sprintf("The highest number of calories is %d", utils.MaxNumber(totalElvesCalories...)),
	})
}