package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Setup sets up any required data for the days puzzle.
func (d *Day) Setup() error {
	readFile, err := os.Open(d.InputFile)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	defer func() {
		if err := readFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		cleaners := strings.Split(line, ",")
		d.Teams = append(d.Teams, d.makeTeam(cleaners))
	}
	return nil
}

func (d *Day) makeTeam(cleaners []string) *Team {
	team := &Team{}
	for _, cleaner := range cleaners {
		work := strings.Split(cleaner, "-")
		team.Cleaners = append(team.Cleaners, d.giveCleanerWork(work[0], work[1]))
	}
	return team
}

func (d *Day) giveCleanerWork(start string, end string) *Cleaner {
	s, err := strconv.Atoi(start)
	if err != nil {
		log.Fatalf("Error converting string: %s \n err:%v", start, err)
	}
	e, err := strconv.Atoi(end)
	if err != nil {
		log.Fatalf("Error converting string: %s \n err:%v", end, err)
	}

	return &Cleaner{
		min: s,
		max: e,
	}
}
