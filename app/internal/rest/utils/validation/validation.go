package validation_rest

import (
	"runtime"
	"strconv"
)

func getFileInfo(fileName string) string {
	_, _, line, _ := runtime.Caller(1)
	return "internal/rest/utils/validation/" + fileName + " line: " + strconv.Itoa(line)
}
