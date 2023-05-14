// Package main is the entry point to the advent2022 application.
package main

import (
	"github/danr57/advent-of-code-2022/advent"
	"github/danr57/advent-of-code-2022/d1"
)

func main() {
	advent.NewAdvent().RunCalendar([]advent.Day{
		d1.New(),
	})
}
