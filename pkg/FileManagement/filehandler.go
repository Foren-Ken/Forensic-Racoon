package FileManagement

import (
	"fmt"
	"os"
)

func OpenFile(userFP []string) (*os.File, error) {
	if len(userFP) < 1 {
		fmt.Println("Error: Filepath is required")
		fmt.Println("The following syntax is expected: program.name \\PATH\\TO\\FILE")
		os.Exit(1)
	}

	filePointer, err := os.Open(userFP[0])

	return filePointer, err
}
