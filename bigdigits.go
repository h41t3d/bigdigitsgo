package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var thedigits = [][]string{
	{
		"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  ",
	},
	{
		" 1 ",
		"11 ",
		" 1 ",
		" 1 ",
		" 1 ",
		" 1 ",
		"111",
	},
	{
		" 222 ",
		"2   2",
		"   2 ",
		"  2  ",
		" 2   ",
		"2    ",
		"22222",
	},
	{
		" 333 ",
		"3   3",
		"    3",
		"  33 ",
		"    3",
		"3   3",
		" 333 ",
	},
	// same for all digits xD too lazy
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	digits := os.Args[1]
	for row := range thedigits[0] {
		line := ""
		for column := range digits {
			hmmdigit := digits[column] - '0'
			if 0 <= hmmdigit && hmmdigit <= 9 {
				line += thedigits[hmmdigit][row] + "  "
			} else {
				log.Fatal("nooooooooooo invalid")
			}
		}
		fmt.Println(line)
	}
}
