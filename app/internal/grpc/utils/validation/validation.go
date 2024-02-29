package validation_grpc

import (
	"runtime"
	"strconv"
)

func getFileInfo(fileName string) string {
	_, _, line, _ := runtime.Caller(1)
	return "internal/grpc/utils/validation/" + fileName + " line: " + strconv.Itoa(line)
}
