// Package main is the entry point to the advent2022 application.
package main

import (
	"github/danr57/advent-of-code-2022/advent"
	"github/danr57/advent-of-code-2022/day1"
	"github/danr57/advent-of-code-2022/day2"
)

func main() {
	advent.NewAdvent().RunCalendar([]advent.Day{
		day1.New(),
		day2.New(),
	})
}
