package validation

import (
	"runtime"
	"strconv"
)

func getFileInfo(fileName string) string {
	_, _, line, _ := runtime.Caller(1)
	return "internal/lib/validation/" + fileName + " line: " + strconv.Itoa(line)
}
