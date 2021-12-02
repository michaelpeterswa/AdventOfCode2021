/*

		 .
	  __/ \__
	  \     /
	  /.'o'.\
	   .o.'.       Michael Peters
	  .'.'o'.      Advent of Code 2021 - Go Edition
	 o'.o.'.o.     Day2!
	.'.o.'.'.o.
   .o.'.o.'.o.'.
	  [_____]
	   \___/    ldb

*/

package day2

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Two struct {
	movements []Movement
}

type Movement struct {
	direction string
	distance  int64
}

type Position struct {
	horiz int64
	vert  int64
	aim   int64
}

func Day2(filename string) {
	fmt.Printf("AoC Day2 2021\n\n")

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var two Two

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		mvStr := strings.Split(string(line), " ")
		dst, err := strconv.ParseInt(mvStr[1], 10, 64)
		if err != nil {
			log.Println("failed to parse distance")
		}

		mvmt := Movement{
			direction: mvStr[0],
			distance:  dst,
		}

		two.movements = append(two.movements, mvmt)
	}

	// PART 1
	p1 := &Position{
		horiz: 0,
		vert:  0,
		aim:   0,
	}

	for _, movement := range two.movements {
		switch movement.direction {
		case "forward":
			p1.Forward(movement.distance)
		case "up":
			p1.Up(movement.distance)
		case "down":
			p1.Down(movement.distance)
		default:
			log.Println("error switching on direction")
		}
	}

	fmt.Println(p1.Mult())

	// PART 2
	p2 := &Position{
		horiz: 0,
		vert:  0,
		aim:   0,
	}

	for _, movement := range two.movements {
		switch movement.direction {
		case "forward":
			p2.Forward2(movement.distance)
		case "up":
			p2.Up2(movement.distance)
		case "down":
			p2.Down2(movement.distance)
		default:
			log.Println("error switching on direction")
		}
	}

	fmt.Println(p2.Mult())

}

func (p *Position) Down(i int64) {
	p.vert += i
}

func (p *Position) Up(i int64) {
	p.vert -= i
}

func (p *Position) Forward(i int64) {
	p.horiz += i
}

func (p *Position) Mult() int64 {
	return p.horiz * p.vert
}

func (p *Position) Down2(i int64) {
	p.aim += i
}

func (p *Position) Up2(i int64) {
	p.aim -= i
}

func (p *Position) Forward2(i int64) {
	p.horiz += i
	p.vert += p.aim * i
}
