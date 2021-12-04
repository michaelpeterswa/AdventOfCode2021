/*

		 .
	  __/ \__
	  \     /
	  /.'o'.\
	   .o.'.       Michael Peters
	  .'.'o'.      Advent of Code 2021 - Go Edition
	 o'.o.'.o.     Day3!
	.'.o.'.'.o.
   .o.'.o.'.o.'.
	  [_____]
	   \___/    ldb

*/

package day3

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Count struct {
	zero int64
	one  int64
}

type Counts struct {
	count []Count
}

type BitSlice struct {
	values []string
}

type ErrorLog struct {
	rows []BitSlice
}

type BitCriteria struct {
	oxy []int64
	co2 []int64
}

func Day3(filename string) {
	fmt.Printf("AoC Day3 2021\n\n")

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var three ErrorLog

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		var row BitSlice
		row.values = strings.Split(string(line), "")
		three.rows = append(three.rows, row)
	}

	// PART ONE

	var counts Counts
	bitLen := len(three.rows[0].values)
	countSlice := make([]Count, bitLen)
	counts = Counts{
		count: countSlice,
	}

	for countIdx := range counts.count {
		for _, bitObj := range three.rows {
			counts.count[countIdx].increment(bitObj.values[countIdx])
		}
	}
	gamma := counts.calculateGamma()
	epsilon := counts.calculateEpsilon()

	prod := gamma * epsilon

	fmt.Printf("Product of Epsilon and Gamma: %d\n", prod)

	// PART TWO
	// Take it row by row
	prod2 := binaryStringToDecimal(returnRating(three, true)) * binaryStringToDecimal(returnRating(three, false))

	fmt.Printf("Life Support Rating: %d\n", prod2)
}

func (cnt Count) computeValue() (int64, error) {
	if cnt.zero == cnt.one {
		err := errors.New("equality detected")
		return -1, err
	} else if cnt.zero > cnt.one {
		return 0, nil
	} else {
		return 1, nil
	}
}

func (cnt Count) computeInverseValue() (int64, error) {
	if cnt.zero == cnt.one {
		err := errors.New("equality detected")
		return -1, err
	} else if cnt.zero > cnt.one {
		return 1, nil
	} else {
		return 0, nil
	}
}

func binaryStringToDecimal(i string) int64 {
	decimal, err := strconv.ParseInt(i, 2, 64)
	if err != nil {
		log.Println(err.Error())
		return -1
	}
	return decimal
}

func (cnt *Count) increment(val string) {
	switch val {
	case "0":
		cnt.incrementZero()
	case "1":
		cnt.incrementOne()
	default:
		log.Println("incorrect number encountered")
	}
}

func (c *Count) init() {
	c = &Count{
		zero: 0,
		one:  0,
	}
}

func (c *Count) incrementZero() {
	c.zero++
}

func (c *Count) incrementOne() {
	c.one++
}

func (c *Counts) calculateGamma() int64 {
	var bin string
	for k := range c.count {
		i, err := c.count[k].computeValue()
		if err != nil {
			log.Println(err.Error())
		}
		bin += strconv.FormatInt(i, 10)
	}

	return binaryStringToDecimal(bin)
}

func (c *Counts) calculateEpsilon() int64 {
	// should do bit flipping but this is 1.5 bytes and i don't wanna go down that path

	var bin string
	for k := range c.count {
		i, err := c.count[k].computeInverseValue()
		if err != nil {
			log.Println(err.Error())
		}
		bin += strconv.FormatInt(i, 10)
	}

	return binaryStringToDecimal(bin)
}

func returnRating(elog ErrorLog, isOxy bool) string {
	var criteria []int64
	var newBS []BitSlice

	newBS = elog.rows
	for i := 0; i < 12; i++ {
		sliceLen := len(newBS)
		if sliceLen != 1 {
			sum := 0
			lenLog := len(newBS)
			for row := range newBS {
				intVal, err := strconv.ParseInt(newBS[row].values[i], 10, 64)
				if err != nil {
					log.Println(err)
				}
				sum += int(intVal)
			}
			floatVal := float64(sum) / float64(lenLog)

			if isOxy {
				if floatVal >= .5 {
					criteria = append(criteria, 1)
				} else {
					criteria = append(criteria, 0)
				}
			} else {
				if floatVal >= .5 {
					criteria = append(criteria, 0)
				} else {
					criteria = append(criteria, 1)
				}
			}

			newBSCopy := newBS
			newBS = nil
			for _, row := range newBSCopy {
				sly := criteria[0 : i+1]
				filter := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(sly)), ""), "[]")
				regex := fmt.Sprintf(`^%s`, filter)

				conc := strings.Join(row.values, "")
				matched, err := regexp.MatchString(regex, conc)
				if err != nil {
					log.Println(err)
				}
				if matched {
					newBS = append(newBS, row)
				}
			}
		} else {
			return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(newBS[0].values)), ""), "[]")
		}
	}
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(newBS[0].values)), ""), "[]")
}
