package main

import (
	"flag"
	"fmt"
)

func main() {
	var typeFlag = flag.String("t", "directory", "Specify either \"Directory\" or \"File\".")
	//Used to determine if the entire directory or just a singular file is being tested.

	var outputFlag = flag.String("o", ".", "Specify output location.")
	// Stores location of output file

	flag.Parse()

	fmt.Println(*typeFlag)
	fmt.Println(*outputFlag)

}
