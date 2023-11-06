package transferFiles

import (
	"fmt"
	"io"
	"os"
)

func TransferFiles(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(destination)
	if err != nil {
		fmt.Println("Error creating the file", err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		fmt.Println("Error copying the file", err)
	}

	fmt.Printf("%s was copied to %s.\n", source, destination)
	return nil
}
