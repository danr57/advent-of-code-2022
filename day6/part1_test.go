package day6_test

import (
	"strings"
	"testing"

	"github/danr57/advent-of-code-2022/day6"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	var (
		exampleSignals = []string{
			"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			"bvwbjplbgvbhsrlpgdmjqwftvncz",
			"nppdvjthqldpwncqszvftbrmjlhg",
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
		}
		exampleAnswers = []string{
			"7",
			"5",
			"6",
			"10",
			"11",
		}
	)

	for index := range exampleSignals {
		index := index
		t.Run(exampleSignals[index], func(t *testing.T) {
			t.Parallel()

			day := day6.Day{
				DataStream: strings.Split(exampleSignals[index], ""),
			}

			ans, err := day.Part1()
			if err != nil {
				t.Fatal(err)
			}

			if ans != exampleAnswers[index] {
				t.Fatalf("wrong answer, expected %v got %v", exampleAnswers[index], ans)
			}
		})
	}
}
