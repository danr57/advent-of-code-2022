// Package main is the entry point to the advent2022 application.
package main

import (
	"github/danr57/advent-of-code-2022/advent"
	"github/danr57/advent-of-code-2022/day1"
	"github/danr57/advent-of-code-2022/day2"
	"github/danr57/advent-of-code-2022/day3"
	"github/danr57/advent-of-code-2022/day4"
	"github/danr57/advent-of-code-2022/day5"
)

func main() {
	advent.NewAdvent().RunCalendar([]advent.Day{
		day1.New(),
		day2.New(),
		day3.New(),
		day4.New(),
		day5.New(),
	})
}
