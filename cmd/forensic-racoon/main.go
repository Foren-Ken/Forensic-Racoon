package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	FileManagement "github.com/Foren-Ken/Forensic-Racoon/pkg/FileManagement"
)

func main() {
	// typeFlag := flag.String("t", "directory", "Specify either \"Directory\" or \"File\".")
	// Used to determine if the entire directory or just a singular file is being tested.

	// outputFlag := flag.String("o", ".", "Specify output location.")
	// Stores location of output file

	flag.Parse()
	// Provides the parsing.

	userFilePath := flag.Args()

	if len(userFilePath) > 1 {
		fmt.Println("Too many arguments!")
		fmt.Println("Please use only one directory as an argument.")
		fmt.Println("\"./forensic-racoon PATH/TO/FILE\" with any amount of optional flags. ")
		os.Exit(1)
	}

	filePointer, err := FileManagement.OpenFile(userFilePath)

	fmt.Println(err)
	fmt.Println(*filePointer)

	var scanner *bufio.Scanner = bufio.NewScanner(filePointer)

	var lineChecker int = 1

	for scanner.Scan() {
		fmt.Printf("Line %d: %s\n", lineChecker, scanner.Text())
		lineChecker++
	}

}
