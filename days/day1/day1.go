/*

		 .
	  __/ \__
	  \     /
	  /.'o'.\
	   .o.'.       Michael Peters
	  .'.'o'.      Advent of Code 2021 - Go Edition
	 o'.o.'.o.     Day1!
	.'.o.'.'.o.
   .o.'.o.'.o.'.
	  [_____]
	   \___/    ldb

*/

package day1

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Depths struct {
	Measurements []int64
}

func Day1(filename string) {
	fmt.Printf("AoC Day1 2021\n\n")

	var depths Depths

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		depth, err := strconv.ParseInt(string(line), 10, 64)
		if err != nil {
			log.Fatal("failed to parse integer")
		}
		depths.Measurements = append(depths.Measurements, depth)
	}

	// PART ONE

	var partOneTotal int64
	depthsLen := len(depths.Measurements)
	for k, depth := range depths.Measurements {
		if k < depthsLen-1 {
			if depth < depths.Measurements[k+1] {
				partOneTotal += 1
			}
		}
	}

	fmt.Printf("Total Depth Increases (Part 1): %d\n", partOneTotal)

	// PART TWO

	var partTwoTotal int64
	for k := range depths.Measurements {
		if k < depthsLen-3 {
			firstSum := depths.Measurements[k] + depths.Measurements[k+1] + depths.Measurements[k+2]
			secondSum := depths.Measurements[k+1] + depths.Measurements[k+2] + depths.Measurements[k+3]
			if firstSum < secondSum {
				partTwoTotal += 1
			}
		}
	}

	fmt.Printf("Total Depth Increases (Part 2): %d\n", partTwoTotal)
}
