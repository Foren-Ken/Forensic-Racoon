package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	typeFlag := flag.String("t", "directory", "Specify either \"Directory\" or \"File\".")
	//Used to determine if the entire directory or just a singular file is being tested.

	outputFlag := flag.String("o", ".", "Specify output location.")
	// Stores location of output file

	flag.Parse()
	// Provides the parsing.

	userFilePath := flag.Args()

	if len(userFilePath) < 1 {
		fmt.Println("Error: Filepath is required")
		fmt.Println("The following syntax is expected: .\\ PATH\\TO\\FILE")
		os.Exit(1)
	}
	doesFileExist, err := os.Stat(userFilePath[0])

	fmt.Println(doesFileExist)
	fmt.Println(err)
	fmt.Println(*outputFlag)
	fmt.Println(*typeFlag)

}
