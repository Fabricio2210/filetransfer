package wrapper

import (
	"fmt"
	"transferfiles.com/memStatus"
	"transferfiles.com/parseNameFile"
	"transferfiles.com/transferFiles"
)

func Wrapper(inputFolderPath, outputFolderPath, nameExtension, fileExtension string) {

	inputFolder := fmt.Sprintf("%s/%s%s.%s", inputFolderPath, parsenamefile.Parsenamefile(), nameExtension, fileExtension)
	outputFolder := fmt.Sprintf("%s/%s%s.%s", outputFolderPath, parsenamefile.Parsenamefile(), nameExtension, fileExtension)

	transferFiles.TransferFiles(inputFolder, outputFolder)
	fmt.Printf("Saving transfer file %s \n", parsenamefile.Parsenamefile())
	memStatus.Memstatus()
}
