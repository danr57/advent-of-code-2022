// Package advent provides helpers for running the advent calendar.
package advent

import (
	"fmt"
	"log"
	"sync"
)

type (
	// Day is the interface that wraps the excpected methods for a day.
	Day interface {
		Setup() error
		Number() int
		Part1() (string, error)
		Part2() (string, error)
	}

	// Advent is the struct that runs the advent calendar.
	Advent struct {
		wg      *sync.WaitGroup
		mut     *sync.RWMutex
		results [25][2]string
	}
)

func (a *Advent) runDay(day Day) {
	defer a.wg.Done()

	if err := day.Setup(); err != nil {
		log.Println(err)

		return
	}

	firstResult, err := day.Part1()
	if err != nil {
		log.Println(err)

		return
	}

	a.mut.Lock()
	a.results[day.Number()-1][0] = firstResult
	a.mut.Unlock()

	secondResult, err := day.Part2()
	if err != nil {
		log.Println(err)

		return
	}

	a.mut.Lock()
	a.results[day.Number()-1][1] = secondResult
	a.mut.Unlock()
}

func (a *Advent) printCalendar() {
	for dayNum, day := range &a.results {
		for partNum, part := range day {
			if part != "" {
				//nolint:forbidigo // This is the format used by the advent package to print the results.
				fmt.Printf("Day %d | Part %d = %s\n",
					dayNum+1,
					partNum+1,
					part,
				)
			}
		}
	}
}

// RunCalendar runs the advent calendar.
func (a *Advent) RunCalendar(days []Day) {
	for _, day := range days {
		a.wg.Add(1)
		go a.runDay(day)
	}

	a.wg.Wait()

	a.printCalendar()
}

// NewAdvent returns a new instance of Advent.
func NewAdvent() *Advent {
	return &Advent{
		wg:      &sync.WaitGroup{},
		mut:     &sync.RWMutex{},
		results: [25][2]string{},
	}
}
