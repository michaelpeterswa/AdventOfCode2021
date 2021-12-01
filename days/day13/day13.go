/*

         .
      __/ \__
      \     /
      /.'o'.\
       .o.'.       Michael Peters
      .'.'o'.      Advent of Code 2021 - Go Edition
     o'.o.'.o.     Day13!
    .'.o.'.'.o.
   .o.'.o.'.o.'.
      [_____]
       \___/    ldb

*/

package day13

import (
    "fmt"
    "os"
)

func Day13(filename string){
    fmt.Println("AoC Day13 2021\n")

    data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}