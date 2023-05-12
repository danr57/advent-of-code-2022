package main

import (
	"github/danr57/advent-of-code-2022/d1"
	"log"
)

func main() {
	// d1output := d1.Part1(69)

	// log.Print(d1output)
	// d1.ReadAndSplit("d1/input")
	log.Printf("Number of elves: %v", d1.ReadAndSplit("d1/input"))
	log.Printf("Most calories an elf has: %v", d1.Part1())
}
