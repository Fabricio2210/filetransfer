package clearfolders

import (
	"fmt"
	"os"
	"path/filepath"
)

func ClearFolders(folderPath string) error {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(folderPath, file.Name())
		err := os.Remove(filePath)
		if err != nil {
			fmt.Printf("Error deleting file %s: %v\n", filePath, err)
		} else {
			fmt.Printf("Deleted file: %s\n", filePath)
		}
	}

	fmt.Println("Files deleted from folder", folderPath)
	return nil
}
