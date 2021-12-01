/*

		 .
	  __/ \__
	  \     /
	  /.'o'.\
	   .o.'.       Michael Peters
	  .'.'o'.      Advent of Code 2021 - Go Edition
	 o'.o.'.o.     Day23!
	.'.o.'.'.o.
   .o.'.o.'.o.'.
	  [_____]
	   \___/    ldb

*/

package day23

import (
	"fmt"
	"os"
)

func Day23(filename string){
	fmt.Println("AoC Day23 2021\n")

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}