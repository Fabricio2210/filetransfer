package parsenamefile

import (
	"time"
)

func Parsenamefile() string {
	today := time.Now()
	fileNameDate := today.AddDate(0, 0, -1)
	fileNameFull := fileNameDate.Format("2006-01-02")

	return fileNameFull
}
