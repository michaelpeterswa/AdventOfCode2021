/*

		 .
	  __/ \__
	  \     /
	  /.'o'.\
	   .o.'.       Michael Peters
	  .'.'o'.      Advent of Code 2021 - Go Edition
	 o'.o.'.o.     Day25!
	.'.o.'.'.o.
   .o.'.o.'.o.'.
	  [_____]
	   \___/    ldb

*/

package day25

import (
	"fmt"
	"os"
)

func Day25(filename string){
	fmt.Printf("AoC Day25 2021\n\n")

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}